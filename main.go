package main

import (
	"fmt"
	"github.com/adi1382/Bitmex-Mirror-Bot/Mirror"
	"github.com/adi1382/Bitmex-Mirror-Bot/hostClient"
	"github.com/adi1382/Bitmex-Mirror-Bot/subClient"
	"github.com/adi1382/Bitmex-Mirror-Bot/tools"
	"github.com/adi1382/Bitmex-Mirror-Bot/websocket"
	"github.com/fsnotify/fsnotify"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"go.uber.org/atomic"
	"go.uber.org/zap"
	"os"
	"sync"
	"time"
)

var (
	sessionID       zap.Field
	botStatus       *atomic.String
	config          *viper.Viper
	restartRequired *atomic.Bool
)

func init() {
	_, err := os.Stat("logs")

	if os.IsNotExist(err) {
		err = os.MkdirAll("logs", 0750)
		if err != nil {
			fmt.Println("Unable to create log folder.")
			tools.EnterToExit()
		}
	} else if err != nil {
		fmt.Println("Unable to perform logging operations due to the following error(s)")
		fmt.Println(err)
		tools.EnterToExit()
	}

	botStatus = atomic.NewString("running")
	restartRequired = atomic.NewBool(false)
	sessionID = zap.String("sessionID", uuid.New().String())
	config = viper.New()
}

func main() {

	logger, _ := tools.NewLogger("Mirror", "debug", sessionID)
	socketIncomingLogger, _ := tools.NewLogger("SocketIncoming", "debug", sessionID)
	socketOutgoingLogger, _ := tools.NewLogger("SocketOutgoing", "debug", sessionID)
	resourceLogger, _ := tools.NewLogger("ResourceLogger", "debug", sessionID)

	go tools.NewMonitor(1, resourceLogger)

	config.SetConfigName("config") // name of config file (without extension)
	config.SetConfigType("json")   // REQUIRED if the config file does not have the extension in the name
	config.AddConfigPath(".")      // optionally look for config in the working directory
	tools.ReadConfig(false, logger, config, botStatus, restartRequired)

	config.WatchConfig()
	config.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("Config File Changed!")
		tools.ReadConfig(true, logger, config, botStatus, restartRequired)
	})

	fmt.Println("started")

	for {

		if botStatus.Load() == "running" {
			restartRequired.Store(false)
			trader(logger, socketIncomingLogger, socketOutgoingLogger)
		}

	}
}

func trader(logger, socketIncomingLogger, socketOutgoingLogger *zap.Logger) {
	var wg sync.WaitGroup

	mirror := Mirror.NewMirror(restartRequired, logger, &wg)

	logger.Info("logging started")

	// Connect to WS
	conn, err := websocket.Connect(config.Sub("Settings").GetBool("Testnet"), logger)

	if err != nil {
		botStatus.Store("Stop")
		fmt.Println(err)
		tools.EnterToExit()
	}

	// Listen write WS
	chReadFromWS := make(chan []byte, 100)
	go websocket.ReadFromWSToChannel(conn, chReadFromWS, socketIncomingLogger, logger, restartRequired, &wg)

	// Write to WS
	chWriteToWS := make(chan interface{}, 100)
	go websocket.WriteFromChannelToWS(conn, chWriteToWS, socketOutgoingLogger, logger, restartRequired, &wg)
	//go config.SocketResponseDistributor(chReadFromWS, RestartCounter, &wg)

	websocket.PingPong(conn, restartRequired, logger, &wg)

	//RestartCounter.Store(0)

	//hostClient.Initialize(cfg.HostAccount.ApiKey, cfg.HostAccount.ApiSecret,
	//	"hostAccount", chWriteToWS, false,
	//	0, cfg.Settings.Testnet, cfg.Settings.RatioUpdateRate, &dummyClient, cfg.Settings.CalibrationRate,
	//	cfg.Settings.LimitFilledTimeout, &RestartCounter)
	//hostClient.subscribeTopics("order", "position", "margin")

	host := hostClient.NewHostClient(
		config.Sub("HostAccount").GetString("ApiKey"),
		config.Sub("HostAccount").GetString("Secret"),
		config.Sub("Settings").GetBool("Testnet"),
		chWriteToWS,
		config.Sub("Settings").GetInt64("RatioUpdateRate"),
		restartRequired,
		logger,
		&wg)

	mirror.SetHost(host)

	subKeys := make([]string, 0, 5)
	for key := range config.Sub("SubAccounts").AllSettings() {
		subKeys = append(subKeys, key)
	}

	subAccounts := config.Sub("SubAccounts")
	for i := range subKeys {

		if subAccounts.Sub(subKeys[i]).GetBool("Enabled") {
			sub := subClient.NewSubClient(
				subAccounts.Sub(subKeys[i]).GetString("ApiKey"),
				subAccounts.Sub(subKeys[i]).GetString("Secret"),
				config.Sub("Settings").GetBool("Testnet"),
				subAccounts.Sub(subKeys[i]).GetBool("BalanceProportion"),
				subAccounts.Sub(subKeys[i]).GetFloat64("FixedProportion"),
				config.Sub("Settings").GetFloat64("RatioUpdateRate"),
				config.Sub("Settings").GetFloat64("CalibrationRate"),
				config.Sub("Settings").GetFloat64("LimitFilledTimeout"),
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

	//fmt.Println("reached")
	//
	////host.WaitForPartial()
	////time.Sleep(5 * time.Second)
	//fmt.Println("Printing margin balance")
	//fmt.Println(host.GetMarginBalance())
	host.WaitForPartial()
	fmt.Println("Running...")
	//fmt.Println("Partials Received")

	go func() {
		wg.Add(1)
		defer wg.Done()
		for {
			if restartRequired.Load() {
				_ = conn.Close()
				chWriteToWS <- "quit"
				break
			}
			<-time.After(time.Nanosecond)
		}
	}()

	wg.Wait()
	logger.Info("All wait groups completed")

	//fmt.Println(host.ActiveOrders())
	//n := 0
	//for {
	//	if len(host.ActiveOrders()) != n {
	//		n = len(host.ActiveOrders())
	//		fmt.Println(host.ActiveOrders())
	//	}
	//	time.Sleep(time.Nanosecond)
	//}
}
