package websocket

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/adi1382/Bitmex-Mirror-Bot/tools"
	"github.com/gorilla/websocket"
	"go.uber.org/atomic"
	"log"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"
)

var (
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
	receiveLogger *log.Logger
	sendLogger    *log.Logger
)

func init() {

	_, err := os.Stat("logs")

	if os.IsNotExist(err) {
		errDir := os.MkdirAll("logs", 0750)
		if errDir != nil {
			ErrorLogger.Fatal(err)
		}

	}

	file, err := os.OpenFile("logs/logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		ErrorLogger.Fatal(err)
	}

	received, err := os.OpenFile("logs/socketReceived.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		ErrorLogger.Fatal(err)
	}

	sent, err := os.OpenFile("logs/socketSent.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		ErrorLogger.Fatal(err)
	}

	receiveLogger = log.New(received, "RECEIVED: ", log.Ldate|log.Ltime|log.Lshortfile)
	sendLogger = log.New(sent, "SENT: ", log.Ldate|log.Ltime|log.Lshortfile)

	receiveLogger.Println("\n\n\n\n" + strings.Repeat("#", 50) + "\tNew Session\t" + strings.Repeat("#", 50) + "\n\n\n\n")
	sendLogger.Println("\n\n\n\n" + strings.Repeat("#", 50) + "\tNew Session\t" + strings.Repeat("#", 50) + "\n\n\n\n")

	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

type Message struct {
	Op   string        `json:"op,omitempty"`
	Args []interface{} `json:"args,omitempty"`
}

func (m *Message) AddArgument(argument string) {
	InfoLogger.Println("Add args to websocket msg", *m, "adding: ", argument)
	m.Args = append(m.Args, argument)
}

func Connect(host string) (*websocket.Conn, error) {
	u := url.URL{Scheme: "wss", Host: host, Path: "/realtimemd"}
	InfoLogger.Println("\n\n\nConnecting to: ", u.String())
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	//if err != nil{
	//	panic(err)
	//}

	return conn, err
}

//func ReadFromWSToChannel(c *websocket.Conn, chRead chan<- []byte, RestartCounter *int32) {
//
//	for {
//		_, message, err := c.ReadMessage()
//		if err != nil {
//			panic(err)
//		}
//
//		fmt.Println(string(message))
//
//		if atomic.LoadInt32(RestartCounter) > 0 {
//			fmt.Println("Read Socket Closed")
//			break
//		}
//
//		chRead <- message
//	}
//}

//func ReadFromWSToChannel(c *websocket.Conn, chRead chan<- []byte, RestartCounter *atomic.Uint32, wg *sync.WaitGroup) {
//	wg.Add(1)
//	defer wg.Done()
//	var message []byte
//	var err error
//	var readWSLocal atomic.Int32
//L:
//	for {
//		readWSLocal.Store(0)
//		go func() {
//			readWSLocal.Store(0)
//			_, message, err = c.ReadMessage()
//			fmt.Println(string(message))
//			readWSLocal.Store(1)
//		}()
//
//		for {
//			time.Sleep(time.Nanosecond)
//			if readWSLocal.Load() == 1 {
//				break
//			} else {
//				if RestartCounter.Load() > 0 {
//					receiveLogger.Println("ReadFromWSToChannel Closed")
//					InfoLogger.Printf("ReadFromWSToChannel Closed")
//					break L
//				}
//			}
//		}
//
//		if RestartCounter.Load() > 0 {
//			receiveLogger.Println("ReadFromWSToChannel Closed")
//			InfoLogger.Printf("ReadFromWSToChannel Closed")
//			break L
//		}
//
//		tools.WebsocketErr(err, RestartCounter)
//		receiveLogger.Println("Length of channel:", len(chRead), "Message:", string(message))
//
//		chRead <- message
//	}
//}

func ReadFromWSToChannel(c *websocket.Conn, chRead chan<- []byte, RestartCounter *atomic.Uint32, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	defer func() {
		fmt.Println("Channel Reader Closed")
	}()

	for {
		messageType, message, err := c.ReadMessage()

		if RestartCounter.Load() != 0 {
			chRead <- []byte("quit")
			break
		}

		if err != nil {
			RestartCounter.Add(1)
			ErrorLogger.Println(err)
			chRead <- []byte("quit")
			break
		}

		if messageType == 1 {
			receiveLogger.Println("Length of channel:", len(chRead), "Message:", string(message))
			chRead <- message
		} else {
			continue
		}
	}
}

//func WriteFromChannelToWS(c *websocket.Conn, chWrite <-chan interface{}, RestartCounter *atomic.Uint32, wg *sync.WaitGroup) {
//	wg.Add(1)
//	defer wg.Done()
//L:
//	for {
//		//a := chWrite.(string)
//		time.Sleep(time.Nanosecond)
//
//		select {
//		case message := <-chWrite:
//
//			if RestartCounter.Load() > 0 {
//
//				if RestartCounter.Load() >= 3 {
//					for len(chWrite) > 0 {
//						<-chWrite
//					}
//					sendLogger.Printf("\n\nWriteFromChannelToWS Closed\n\n")
//					_ = c.Close()
//					InfoLogger.Println(runtime.NumGoroutine())
//					InfoLogger.Printf("\n\nWriteFromChannelToWS Closed\n\n")
//					break L
//				}
//
//				for len(chWrite) > 0 {
//					<-chWrite
//				}
//				continue L
//			}
//
//			message, err := json.Marshal(message)
//			tools.WebsocketErr(err, RestartCounter)
//
//			sendLogger.Println("Length of channel:", len(chWrite), "Message:", string(message.([]byte)))
//			err = c.WriteMessage(websocket.TextMessage, message.([]byte))
//
//			tools.WebsocketErr(err, RestartCounter)
//
//		default:
//			//fmt.Println(atomic.LoadInt32(RestartCounter))
//			if RestartCounter.Load() >= 3 {
//				for len(chWrite) > 0 {
//					<-chWrite
//				}
//				sendLogger.Printf("\n\nWriteFromChannelToWS Closed\n\n")
//				_ = c.Close()
//				InfoLogger.Println(runtime.NumGoroutine())
//				InfoLogger.Printf("\n\nWriteFromChannelToWS Closed\n\n")
//				break L
//			}
//		}
//	}
//}

func WriteFromChannelToWS(c *websocket.Conn, chWrite <-chan interface{}, RestartCounter *atomic.Uint32, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	defer func() {
		fmt.Println("Socket writer Closed")
	}()

	for {
		message := <-chWrite
		s, ok := message.(string)
		if ok {
			if s == "quit" {
				break
			}
		}

		message, err := json.Marshal(message)
		if err != nil {
			continue
		}

		sendLogger.Println("Length of channel:", len(chWrite), "Message:", string(message.([]byte)))
		err = c.WriteMessage(websocket.TextMessage, message.([]byte))

		tools.WebsocketErr(err, RestartCounter)
	}
}

func GetAuthMessage(key, secret string) Message {
	timestamp := time.Now().Add(time.Second * 412).Unix()
	sig := hmac.New(sha256.New, []byte(secret))
	_, err := sig.Write([]byte(fmt.Sprintf("GET/realtime%d", timestamp)))

	if err != nil {
		ErrorLogger.Fatal("Was unable to generate signature for socket authentication!!!")
	}

	signature := hex.EncodeToString(sig.Sum(nil))

	var msgKey []interface{}
	msgKey = append(msgKey, key, timestamp, signature)

	return Message{"authKeyExpires", msgKey}
}
