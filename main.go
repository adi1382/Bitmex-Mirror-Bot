package main

//
//import (
//	"fmt"
//	"github.com/adi1382/Bitmex-Mirror-Bot/subClient"
//	"github.com/adi1382/Bitmex-Mirror-Bot/config"
//	"github.com/adi1382/Bitmex-Mirror-Bot/tools"
//	"github.com/adi1382/Bitmex-Mirror-Bot/websocket"
//	"log"
//	"os"
//	"strings"
//	"sync"
//	"sync/atomic"
//	"time"
//)
//
//var (
//	InfoLogger     *log.Logger
//	ErrorLogger    *log.Logger
//	RestartCounter int32
//)
//
//func init() {
//
//	_, err := os.Stat("logs")
//
//	if os.IsNotExist(err) {
//		errDir := os.MkdirAll("logs", 0755)
//		if errDir != nil {
//			ErrorLogger.Fatal(err)
//		}
//	}
//
//	file, err := os.OpenFile("logs/logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
//	if err != nil {
//		ErrorLogger.Fatal(err)
//	}
//
//	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
//	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
//
//	InfoLogger.Println("\n\n\n\n" + strings.Repeat("#", 50) + "\tNew Session\t" + strings.Repeat("#", 50) + "\n\n\n\n")
//}
//
//// Usage example
//func main() {
//	//c := wmic.GetHashedKey()
//	//if c != keys.HashedKey {
//	//	fmt.Println("Unauthorized access to mirror trader detected.")
//	//	fmt.Println("Closing in 5 seconds...")
//	//	time.Sleep(time.Second*5)
//	//	os.Exit(-1)
//	//}
//
//
//
//	const expireTime = 1596689979
//	//fmt.Println(time.Now().Add(15*24*time.Hour).Unix())
//	//fmt.Println(time.Now().Add(time.Minute*1).Unix())
//	//fmt.Println(time.Now().Add(time.Hour*24).Unix())
//	//fmt.Println((expireTime-time.Now().Unix())/3600)
//
//	if time.Now().Unix() > expireTime {
//		fmt.Println("Trial Expired!")
//		time.Sleep(time.Second*10)
//		os.Exit(-1)
//	}
//
//	go func() {
//		for {
//			if time.Now().Unix() > expireTime {
//				fmt.Println("Trial Expired!")
//				time.Sleep(time.Second*10)
//				os.Exit(0)
//			}
//			time.Sleep(5 * time.Second)
//		}
//	}()
//
//	go tools.NewMonitor(1)
//
//	for {
//		mirrorTrader()
//		fmt.Printf("\n\nRestarting!!\n\n")
//		RestartCounter = 0
//	}
//
//}
//
//func mirrorTrader() {
//
//	//AllClients := make([]subClient.Client, 0, 5)
//	//subClient.AllClients = []*subClient.Client
//
//	//panic("aaa")
//
//	defer func() {
//		subClient.AllClientsLock.Lock()
//		for _, v := range subClient.AllClients {
//			v.DropConnection()
//		}
//		subClient.AllClients = make([]*subClient.Client, 0, 5)
//		subClient.AllClientsLock.Unlock()
//		subClient.InfoLogger.Println("Distributor closed")
//	}()
//	var wg sync.WaitGroup
//
//	//go func() {
//	//	time.Sleep(10*time.Second)
//	//	fmt.Println(runtime.NumGoroutine())
//	//	atomic.AddInt32(&RestartCounter, 1)
//	//}()
//	//quit := make(chan int, 100)
//
//	// Close Protocol
//	// 1. Read WS to Channel, Write to ws (stop sending)
//	// 2. Socket Distributor, Ping Pong
//	// 3. Order Handler(3), Data Handler, Margin Update
//	// 4. Close Socket
//	// 5. Write to ws
//
//	fmt.Println("\nInitiating.....")
//	cfg := config.LoadConfig("config.json")
//
//	go func() {
//		file, err := os.Stat("config.json")
//
//		if err != nil {
//			fmt.Println(err)
//		}
//
//		modifiedTime := file.ModTime().Unix()
//		for {
//			//fmt.Println("running")
//			file, err := os.Stat("config.json")
//
//			if err != nil {
//				fmt.Println(err)
//			}
//
//			if file.ModTime().Unix() != modifiedTime {
//				atomic.AddInt32(&RestartCounter, 1)
//				fmt.Println("Config Change Detected")
//				break
//			}
//
//			//cf2 := config.LoadConfig("config.json")
//			//
//			//fmt.Println("Last modified time : ", modifiedTime)
//			//
//			//if !reflect.DeepEqual(cfg, cf2) {
//			//	atomic.AddInt32(&RestartCounter, 1)
//			//	fmt.Println("Config Change Detected")
//			//	break
//			//}
//			time.Sleep(time.Nanosecond)
//		}
//	}()
//
//	var baseUrl string
//	if cfg.Settings.Testnet {
//		baseUrl = "testnet.bitmex.com"
//	} else {
//		baseUrl = "www.bitmex.com"
//	}
//
//	//tools.CheckConnection(&baseUrl)
//
//	var hostClient subClient.Client
//
//	// Connect to WS
//	conn, err := websocket.Connect(baseUrl)
//
//	if err != nil {
//		InfoLogger.Println("Connection Failed!")
//	} else {
//		defer func() {
//			//fmt.Println("Socket closed")
//			_ = conn.Close()
//		}()
//
//		// Listen read WS
//		chReadFromWS := make(chan []byte, 100)
//		go websocket.ReadFromWSToChannel(conn, chReadFromWS, &RestartCounter)
//
//		// Listen write WS
//		chWriteToWS := make(chan interface{}, 100)
//		wg.Add(1)
//		go websocket.WriteFromChannelToWS(conn, chWriteToWS, &RestartCounter, &wg)
//
//		wg.Add(1)
//		go config.SocketResponseDistributor(chReadFromWS, &RestartCounter, &wg)
//
//		var dummyClient subClient.Client
//
//		hostClient.Initialize(cfg.HostAccount.ApiKey, cfg.HostAccount.ApiSecret,
//			"hostAccount", chWriteToWS, false,
//			0, cfg.Settings.Testnet, cfg.Settings.RatioUpdateRate, &dummyClient, cfg.Settings.CalibrationRate,
//			cfg.Settings.LimitFilledTimeout, &RestartCounter)
//		hostClient.subscribeTopics("order", "position", "margin")
//
//		var numberOfSubs int
//		for _, v := range cfg.SubAccounts {
//			if v.Enabled {
//				numberOfSubs++
//			}
//		}
//
//		subClients := make([]subClient.Client, numberOfSubs)
//
//		{
//			k := 0
//			for _, v := range cfg.SubAccounts {
//				if v.Enabled {
//					subClients[k].Initialize(v.ApiKey, v.Secret, "", chWriteToWS,
//						v.BalanceProportion, v.FixedProportion, cfg.Settings.Testnet,
//						cfg.Settings.RatioUpdateRate,
//						&hostClient, cfg.Settings.CalibrationRate, cfg.Settings.LimitFilledTimeout, &RestartCounter)
//					k++
//				}
//			}
//		}
//		for i := range subClients {
//			subClients[i].subscribeTopics("order", "position", "margin")
//		}
//
//		//fmt.Println("Stop 1")
//
//		for i := range subClients {
//			subClients[i].WaitForPartial()
//		}
//
//		hostClient.WaitForPartial()
//
//		//fmt.Println("Stop 2")
//
//		websocket.PingPong(conn, &RestartCounter)
//
//		for i := range subClients {
//			go subClients[i].OrderHandler()
//		}
//
//		fmt.Println("\nNow Running..")
//
//		for {
//			if atomic.LoadInt32(&RestartCounter) > 0 {
//				wg.Wait()
//				break
//			}
//			time.Sleep(time.Nanosecond)
//		}
//	}
//}
