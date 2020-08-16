package main

import (
	"fmt"
	"github.com/adi1382/Bitmex-Mirror-Bot/Mirror"
	"github.com/adi1382/Bitmex-Mirror-Bot/hostClient"
	"github.com/adi1382/Bitmex-Mirror-Bot/subClient"
	"github.com/adi1382/Bitmex-Mirror-Bot/websocket"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/atomic"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

var (
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
)

func init() {

	_, err := os.Stat("logs")

	if os.IsNotExist(err) {
		errDir := os.MkdirAll("logs", 0750)
		if errDir != nil {
			ErrorLogger.Fatal(err)
		}
	}

	file, err := os.OpenFile("logs/logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		ErrorLogger.Fatal(err)
	}

	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	InfoLogger.Println("\n\n\n\n" + strings.Repeat("#", 50) + "\tNew Session\t" + strings.Repeat("#", 50) + "\n\n\n\n")
}

func main() {

	var RestartC uint32
	RestartCounter := atomic.NewUint32(RestartC)
	var wg sync.WaitGroup
	var mirror Mirror.Mirror
	mirror.RestartCounter = RestartCounter

	//cfg := config.LoadConfig("config.json")
	viper.SetConfigName("config") // name of config file (without extension)
	//viper.SetConfigType("json") // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")    // optionally look for config in the working directory
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	//fmt.Println(viper.AllKeys())
	fmt.Println(viper.AllKeys())
	fmt.Println(viper.Sub("Settings").GetDuration("CalibrationRate") * time.Second)

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println(in)
		fmt.Println("Config File Changed!")
	})

	fmt.Println(viper.Sub("SubAccounts").AllSettings())
	fmt.Println(len(viper.Sub("SubAccounts").AllSettings()))
	//fmt.Println(viper.AllSettings()["subaccounts"])

	//for {
	//
	//}

	//os.Exit(0)

	fmt.Println("started")

	baseUrl := "testnet.bitmex.com"

	// Connect to WS
	conn, err := websocket.Connect(baseUrl)

	if err != nil {
		panic(err)
	}

	chReadFromWS := make(chan []byte, 100)
	go websocket.ReadFromWSToChannel(conn, chReadFromWS, RestartCounter, &wg)

	// Listen write WS
	chWriteToWS := make(chan interface{}, 100)
	go websocket.WriteFromChannelToWS(conn, chWriteToWS, RestartCounter, &wg)
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
