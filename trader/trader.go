package trader

import (
	"fmt"
	"github.com/adi1382/Bitmex-Mirror-Bot/Mirror"
	"github.com/adi1382/Bitmex-Mirror-Bot/configuration"
	"github.com/adi1382/Bitmex-Mirror-Bot/hostClient"
	"github.com/adi1382/Bitmex-Mirror-Bot/subClient"
	"github.com/adi1382/Bitmex-Mirror-Bot/tools"
	"github.com/adi1382/Bitmex-Mirror-Bot/websocket"
	"go.uber.org/atomic"
	"go.uber.org/zap"
	"sync"
	"time"
)

func BotHandler(logger, socketIncomingLogger, socketOutgoingLogger *zap.Logger, botStatus *tools.RunningStatus, restartRequired *atomic.Bool) {
	for {
		if botStatus.IsRunning.Load() {
			restartRequired.Store(false)
			fmt.Println("Trader Initiated..")
			trader(logger, socketIncomingLogger, socketOutgoingLogger, botStatus, restartRequired)
			fmt.Printf("\n\n\n")
		}
		time.Sleep(time.Millisecond * 100)
	}
}

func trader(logger, socketIncomingLogger, socketOutgoingLogger *zap.Logger, botStatus *tools.RunningStatus, restartRequired *atomic.Bool) {

	defer func() {
		fmt.Println("Trader closed..")
	}()

	config := configuration.ReadConfig(logger, botStatus, restartRequired)
	isOK, message := configuration.ValidateConfig(&config, logger)

	if !isOK {
		fmt.Println("Invalid Configuration File")
		fmt.Println(message)
		botStatus.Message.Store(message)
		botStatus.IsRunning.Store(false)
	}

	if !botStatus.IsRunning.Load() {
		return
	}

	var wg sync.WaitGroup

	mirror := Mirror.NewMirror(restartRequired, logger, &wg)

	// Connect to WS
	conn, err := websocket.Connect(config.Settings.Testnet, logger)

	if err != nil {
		botStatus.IsRunning.Store(false)
		botStatus.Message.Store("Unable to establish websocket connection")
		fmt.Println(err)
		tools.EnterToExit("Unable to establish websocket connection")
	}

	// Listen write WS
	chReadFromWS := make(chan []byte, 100)
	go websocket.ReadFromWSToChannel(conn, chReadFromWS, socketIncomingLogger, logger, restartRequired, &wg)

	// Write to WS
	chWriteToWS := make(chan interface{}, 100)
	go websocket.WriteFromChannelToWS(conn, chWriteToWS, socketOutgoingLogger, logger, restartRequired, &wg)
	//go viperConfig.SocketResponseDistributor(chReadFromWS, RestartCounter, &wg)

	websocket.PingPong(conn, restartRequired, logger, &wg)

	host := hostClient.NewHostClient(
		config.HostAccount.ApiKey,
		config.HostAccount.Secret,
		config.Settings.Testnet,
		chWriteToWS,
		config.Settings.RatioUpdateRate,
		restartRequired,
		logger,
		&wg,
		botStatus)

	mirror.SetHost(host)

	subKeys := make([]string, 0, 5)
	for key := range config.SubAccounts {
		subKeys = append(subKeys, key)
	}

	//subAccounts := viperConfig.Sub("SubAccounts")
	for i := range subKeys {

		if config.SubAccounts[subKeys[i]].Enabled {
			sub := subClient.NewSubClient(
				config.SubAccounts[subKeys[i]].ApiKey,
				config.SubAccounts[subKeys[i]].Secret,
				config.Settings.Testnet,
				config.SubAccounts[subKeys[i]].BalanceProportion,
				config.SubAccounts[subKeys[i]].FixedProportion,
				config.Settings.RatioUpdateRate,
				config.Settings.CalibrationRate,
				config.Settings.LimitFilledTimeout,
				chWriteToWS,
				restartRequired,
				host,
				logger,
				&wg)
			mirror.AddSub(sub)
		}
	}

	go mirror.SocketResponseDistributor(chReadFromWS)

	mirror.InitializeAll()

	mirror.StartMirroring()

	fmt.Println("Running...")

	go func() {
		wg.Add(1)
		defer wg.Done()
		for {
			if restartRequired.Load() {
				_ = conn.Close()
				chWriteToWS <- "quit"
				break
			}
			<-time.After(time.Millisecond * 100)
		}
	}()

	wg.Wait()
	logger.Info("All wait groups completed")
}
