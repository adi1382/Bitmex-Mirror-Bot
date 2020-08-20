package main

import (
	"fmt"
	"github.com/adi1382/Bitmex-Mirror-Bot/Mirror"
	"github.com/adi1382/Bitmex-Mirror-Bot/hostClient"
	"github.com/adi1382/Bitmex-Mirror-Bot/subClient"
	"github.com/adi1382/Bitmex-Mirror-Bot/tools"
	"github.com/adi1382/Bitmex-Mirror-Bot/websocket"
	"github.com/fsnotify/fsnotify"
	guuid "github.com/google/uuid"
	"github.com/spf13/viper"
	"go.uber.org/atomic"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
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
			//tools.EnterToExit()
		}
	} else if err != nil {
		fmt.Println("Unable to perform logging operations due to the following error(s)")
		fmt.Println(err)
	}

	botStatus = atomic.NewString("running")
	restartRequired = atomic.NewBool(false)
	sessionID = zap.String("sessionID", guuid.New().String())
	config = viper.New()
}

func main() {

	logger, _ := NewLogger("Mirror", "debug")
	socketIncomingLogger, _ := NewLogger("SocketIncoming", "debug")
	socketOutgoingLogger, _ := NewLogger("SocketOutgoing", "debug")

	var RestartC uint32
	RestartCounter := atomic.NewUint32(RestartC)

	var wg sync.WaitGroup

	mirror := Mirror.NewMirror(restartRequired, logger, &wg)

	//cfg := config.LoadConfig("config.json")
	config.SetConfigName("config") // name of config file (without extension)
	config.SetConfigType("json")   // REQUIRED if the config file does not have the extension in the name
	config.AddConfigPath(".")      // optionally look for config in the working directory
	ReadConfig(false, logger)

	config.WatchConfig()
	config.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("Config File Changed!")
		ReadConfig(true, logger)
	})

	//fmt.Println(viper.Sub("SubAccounts").AllSettings())
	//fmt.Println(len(viper.Sub("SubAccounts").AllSettings()))
	//fmt.Println(viper.AllSettings()["subaccounts"])

	//for {
	//
	//}

	//os.Exit(0)

	fmt.Println("started")

	logger.Info("hellp")
	//os.Exit(1)

	// Connect to WS
	conn, err := websocket.Connect(config.Sub("Settings").GetBool("Testnet"), logger)

	if err != nil {
		botStatus.Store("Stop")
		fmt.Println(err)
		tools.EnterToExit()
	}

	// Listen write WS
	chReadFromWS := make(chan []byte, 100)
	go websocket.ReadFromWSToChannel(conn, chReadFromWS, socketIncomingLogger, logger, restartRequired)

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
		RestartCounter)

	mirror.SetHost(host)

	subKeys := make([]string, 0, 5)
	for key := range config.Sub("SubAccounts").AllSettings() {
		subKeys = append(subKeys, key)
	}

	subAccounts := config.Sub("SubAccounts")
	for i := range subKeys {

		sub := subClient.NewSubClient(
			subAccounts.Sub(subKeys[i]).GetString("ApiKey"),
			subAccounts.Sub(subKeys[i]).GetString("Secret"),
			config.Sub("Settings").GetBool("Testnet"),
			subAccounts.Sub(subKeys[i]).GetBool("BalanceProportion"),
			subAccounts.Sub(subKeys[i]).GetFloat64("FixedRatio"),
			config.Sub("Settings").GetFloat64("RatioUpdateRate"),
			config.Sub("Settings").GetFloat64("CalibrationRate"),
			config.Sub("Settings").GetFloat64("LimitFilledTimeout"),
			chWriteToWS,
			RestartCounter,
			host,
			logger)

		mirror.AddSub(sub)
	}

	go mirror.SocketResponseDistributor(chReadFromWS)

	mirror.InitializeAll()
	mirror.StartMirroring()

	fmt.Println("reached")

	//host.WaitForPartial()
	//time.Sleep(5 * time.Second)
	fmt.Println("Printing margin balance")
	fmt.Println(host.GetMarginBalance())
	host.WaitForPartial()
	fmt.Println("Partials Received")

	go func() {
		wg.Add(1)
		defer wg.Done()
		for {
			if RestartCounter.Load() > 0 {
				_ = conn.Close()
				chWriteToWS <- "quit"
				break
			}
		}
	}()

	//go func() {
	//	time.Sleep(time.Second*20)
	//	RestartCounter.Add(1)
	//	fmt.Println("Added to counter", time.Now())
	//}()

	wg.Wait()
	fmt.Println("All wait groups completed", time.Now())

	//fmt.Println(host.ActiveOrders())
	//n := 0
	//for {
	//	if len(host.ActiveOrders()) != n {
	//		n = len(host.ActiveOrders())
	//		fmt.Println(host.ActiveOrders())
	//	}
	//	time.Sleep(time.Nanosecond)
	//}

	select {}
}

func ReadConfig(restart bool, logger *zap.Logger) {

	defer logger.Sync()

	err := config.ReadInConfig() // Find and read the config file

	if err != nil {
		fmt.Println("The recent changes that were made to the config file have made it inaccessible.")
		fmt.Println("Kindly reconfigure the configuration and restart the program.")
		logger.Error("Unable to Read config file", zap.Error(err))
		botStatus.Store("stop")
		//tools.EnterToExit()
	} else {
		fmt.Println(time.Now())
		isConfigValid, str := tools.CheckConfig(config)

		if !isConfigValid {
			fmt.Println("Invalid Configuration")
			fmt.Println(str)
			tools.EnterToExit()
		}

		if restart {
			restartRequired.Store(true)
			botStatus.Store("running")
		}
	}
}

func NewLogger(fileName, level string) (*zap.Logger, error) {

	logLevel := zap.DebugLevel

	if strings.EqualFold(level, "INFO") {
		logLevel = zap.InfoLevel
	} else if strings.EqualFold(level, "WARN") {
		logLevel = zap.WarnLevel
	} else if strings.EqualFold(level, "ERROR") {
		logLevel = zap.ErrorLevel
	}

	config := zap.Config{
		Level:       zap.NewAtomicLevelAt(logLevel),
		Development: true,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         "json",
		EncoderConfig:    zap.NewProductionEncoderConfig(),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
		//OutputPaths:      []string{"./logs/" + fileName + ".log"},
		//ErrorOutputPaths: []string{"./logs/" + fileName + ".log"},
	}

	config.EncoderConfig.TimeKey = "TimeUTC"

	config.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.UTC().Format("2006-01-02T15:04:05Z0700"))
		// 2019-08-13T04:39:11Z
	}

	logger, err := config.Build()

	if err != nil {
		panic(err)
	}

	logger = logger.With(sessionID)
	return logger, err
}
