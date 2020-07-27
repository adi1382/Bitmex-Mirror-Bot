package main

import (
	"fmt"
	"github.com/adi1382/Bitmex-Mirror-Bot/Mirror"
	"github.com/adi1382/Bitmex-Mirror-Bot/hostClient"
	"github.com/adi1382/Bitmex-Mirror-Bot/websocket"
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
		errDir := os.MkdirAll("logs", 0755)
		if errDir != nil {
			ErrorLogger.Fatal(err)
		}
	}

	file, err := os.OpenFile("logs/logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
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
	var mirror Mirror.Mirror

	fmt.Println("started")

	baseUrl := "testnet.bitmex.com"

	// Connect to WS
	conn, err := websocket.Connect(baseUrl)

	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup

	chReadFromWS := make(chan []byte, 100)
	go websocket.ReadFromWSToChannel(conn, chReadFromWS, RestartCounter)

	// Listen write WS
	chWriteToWS := make(chan interface{}, 100)
	wg.Add(1)
	go websocket.WriteFromChannelToWS(conn, chWriteToWS, RestartCounter, &wg)

	wg.Add(1)
	go mirror.SocketResponseDistributor(chReadFromWS, RestartCounter, &wg)
	//go config.SocketResponseDistributor(chReadFromWS, RestartCounter, &wg)

	websocket.PingPong(conn, RestartCounter)

	RestartCounter.Store(0)

	//hostClient.Initialize(cfg.HostAccount.ApiKey, cfg.HostAccount.ApiSecret,
	//	"hostAccount", chWriteToWS, false,
	//	0, cfg.Settings.Testnet, cfg.Settings.RatioUpdateRate, &dummyClient, cfg.Settings.CalibrationRate,
	//	cfg.Settings.LimitFilledTimeout, &RestartCounter)
	//hostClient.subscribeTopics("order", "position", "margin")

	hostApi := "hPawIhWrPeMAEpdmjDBZXZqw"
	hostSecret := "_IQpR1WpEX2Ls4J8QhJHUX82W9xbjZHRsyUOoWlko2tfB0AK"
	host := hostClient.NewHostClient(hostApi, hostSecret, true, chWriteToWS, 10, RestartCounter)
	mirror.Host = host
	fmt.Println("reached")
	host.Initialize()
	host.SubscribeTopics("order", "position", "margin")
	//host.WaitForPartial()
	time.Sleep(5 * time.Second)
	fmt.Println("Printing margin balance")
	fmt.Println(host.GetMarginBalance())
	host.WaitForPartial()
	fmt.Println("Partials Received")

	//fmt.Println(host.ActiveOrders())
	n := 0
	for {
		if len(host.ActiveOrders()) != n {
			n = len(host.ActiveOrders())
			fmt.Println(host.ActiveOrders())
		}
		time.Sleep(time.Nanosecond)
	}

	//select {}
}
