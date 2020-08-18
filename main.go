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
	"strings"
	"sync"
	"time"
)

var (
	sessionID            zap.Field
	logger               *zap.Logger
	socketIncomingLogger *zap.Logger
	socketOutgoingLogger *zap.Logger
	botStatus            *atomic.Bool
)

func ReadConfig(RestartCounter *atomic.Uint32, restart bool) {
	err := viper.ReadInConfig() // Find and read the config file

	if restart {
		RestartCounter.Add(1)
	}

	if err != nil {
		fmt.Println("The recent changes that were made to the config file have made it inaccessible.")
		fmt.Println("Kindly reconfigure the configuration and restart the program.")
		logger.Error("Unable to Read config file", zap.Error(err))
		botStatus.Store(false)
		tools.EnterToExit()
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
		OutputPaths:      []string{"./logs/" + fileName + ".log"},
		ErrorOutputPaths: []string{"./logs/" + fileName + ".log"},
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

func init() {
	botStatus.Store(true)
	sessionID = zap.String("sessionID", guuid.New().String())
	logger, _ = NewLogger("Mirror", "debug")
	socketIncomingLogger, _ = NewLogger("Mirror", "debug")
	socketOutgoingLogger, _ = NewLogger("Mirror", "debug")
}

func main() {
	var RestartC uint32
	RestartCounter := atomic.NewUint32(RestartC)
	var wg sync.WaitGroup
	var mirror Mirror.Mirror
	mirror.RestartCounter = RestartCounter

	//cfg := config.LoadConfig("config.json")
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("json")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	ReadConfig(RestartCounter, false)
	//err := viper.ReadInConfig()   // Find and read the config file
	//
	//if err != nil {
	//	fmt.Println("The recent changes that were made to the config file have made it inaccessible.")
	//	fmt.Println("Kindly reconfigure the configuration and restart the program.")
	//	logger.Error("Unable to Read config file", zap.Error(err))
	//	tools.EnterToExit()
	//}

	//fmt.Println(viper.AllKeys())
	//fmt.Println(viper.AllKeys())
	//fmt.Println(viper.Sub("Settings").GetDuration("CalibrationRate") * time.Second)

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("Config File Changed!")
		ReadConfig(RestartCounter, true)
	})

	//fmt.Println(viper.Sub("SubAccounts").AllSettings())
	//fmt.Println(len(viper.Sub("SubAccounts").AllSettings()))
	//fmt.Println(viper.AllSettings()["subaccounts"])

	//for {
	//
	//}

	//os.Exit(0)

	fmt.Println("started")

	var baseUrl string
	if viper.Sub("Settings").GetBool("Testnet") {
		baseUrl = "testnet.bitmex.com"
	} else {
		baseUrl = "www.bitmex.com"
	}

	// Connect to WS

	conn := websocket.GetNewConnection()

	{
		var err error

		for i := 0; ; i++ {
			conn, err = websocket.Connect(baseUrl)
			if err != nil {
				logger.Error(
					"Unable to establish connection with websockets", zap.Error(err),
					zap.Int("attempt", i))

				if i > 5 {
					botStatus.Store(false)
				}

				continue
			}
			break
		}
	}

	chReadFromWS := make(chan []byte, 100)
	go websocket.ReadFromWSToChannel(conn, chReadFromWS, RestartCounter, &wg, socketOutgoingLogger)

	// Listen write WS
	chWriteToWS := make(chan interface{}, 100)
	go websocket.WriteFromChannelToWS(conn, chWriteToWS, RestartCounter, &wg, socketIncomingLogger)
	//go config.SocketResponseDistributor(chReadFromWS, RestartCounter, &wg)

	websocket.PingPong(conn, RestartCounter, &wg)

	RestartCounter.Store(0)

	//hostClient.Initialize(cfg.HostAccount.ApiKey, cfg.HostAccount.ApiSecret,
	//	"hostAccount", chWriteToWS, false,
	//	0, cfg.Settings.Testnet, cfg.Settings.RatioUpdateRate, &dummyClient, cfg.Settings.CalibrationRate,
	//	cfg.Settings.LimitFilledTimeout, &RestartCounter)
	//hostClient.subscribeTopics("order", "position", "margin")

	subApi := "L-uYMpCeBqQswmFL971PcZIs"
	subSecret := "zlhP-sCl0rXEgm5Y1o7I2qmOtYZPQXR6_0wigHk05kYZ2Dej"

	host := hostClient.NewHostClient(
		viper.Sub("HostAccount").GetString("ApiKey"),
		viper.Sub("HostAccount").GetString("ApiSecret"),
		viper.Sub("Settings").GetBool("Testnet"),
		chWriteToWS,
		viper.Sub("Settings").GetInt64("RatioUpdateRate"),
		RestartCounter)

	for _, value := range viper.Sub("SubAccounts").AllSettings() {
		mirror.Subs = append(
			mirror.Subs,
			subClient.NewSubClient(
				value.(map[string]interface{})["ApiKey"].(string),
				value.(map[string]interface{})["Secret"].(string),
				viper.Sub("Settings").GetBool("Testnet"),
				value.(map[string]interface{})["BalanceProportion"].(bool),
				1,
				chWriteToWS,
				60,
				10,
				10,
				RestartCounter,
				host))
	}

	sub := subClient.NewSubClient(subApi, subSecret, true, false, 1, chWriteToWS,
		60, 10, 10, RestartCounter, host)
	mirror.Host = host
	mirror.Subs = append(mirror.Subs, sub)

	go mirror.SocketResponseDistributor(chReadFromWS, &wg)

	sub.Initialize()
	host.Initialize()
	host.SubscribeTopics("order", "position", "margin")

	sub.SubscribeTopics("order", "position", "margin")

	sub.WaitForPartial()
	fmt.Println("Sub Partial Received")

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
