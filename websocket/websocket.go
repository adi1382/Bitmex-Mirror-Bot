package websocket

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"go.uber.org/atomic"
	"go.uber.org/zap"
	"log"
	"net/url"
	"os"
	"sync"
	"time"
)

var (
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
)

func init() {

	//_, err := os.Stat("logs")
	//
	//if os.IsNotExist(err) {
	//	errDir := os.MkdirAll("logs", 0750)
	//	if errDir != nil {
	//		ErrorLogger.Fatal(err)
	//	}
	//
	//}

	file, err := os.OpenFile("logs/logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		ErrorLogger.Fatal(err)
	}

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

func Connect(test bool, logger *zap.Logger) (*websocket.Conn, error) {

	defer logger.Sync()

	var host string

	if test {
		host = "testnet.bitmex.com"
	} else {
		host = "www.bitmex.com"
	}

	for i := 0; ; i++ {
		u := url.URL{Scheme: "wss", Host: host, Path: "/realtimemd"}
		InfoLogger.Println("\n\n\nConnecting to: ", u.String())
		conn, httpResponse, err := websocket.DefaultDialer.Dial(u.String(), nil)

		if err != nil {
			logger.Sugar().Error(
				"Unable to establish websocket connection",
				conn, httpResponse, err, zap.Int("attempt", i))

			if i > 5 {
				return nil, fmt.Errorf("cannot connect to Bitmex's websocket")
			}
			time.Sleep(time.Second)
			continue
		}

		return conn, err
	}
	//if err != nil{
	//	panic(err)
	//}
}

func GetNewConnection() *websocket.Conn {
	return &websocket.Conn{}
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

func ReadFromWSToChannel(
	c *websocket.Conn,
	chRead chan<- []byte,
	incomingLogger *zap.Logger,
	logger *zap.Logger,
	RestartRequired *atomic.Bool) {

	logger.Info("Socket Status", zap.String("status", "Socket Reader Routine Started"))

	defer func() {
		logger.Info("Socket Status", zap.String("status", "Socket Reader Routine Closed"))
	}()

	for {
		messageType, message, err := c.ReadMessage()

		if err != nil {
			RestartRequired.Store(true)

			if websocket.IsUnexpectedCloseError(err) {
				logger.Error("Unexpected websocket closure", zap.Error(err))
			}

			if websocket.IsCloseError(err) {
				logger.Info("Expected websocket closure", zap.Error(err))
			}

			fmt.Println("Expected: ", websocket.IsCloseError(err))
			fmt.Println("Unexpected: ", websocket.IsUnexpectedCloseError(err))

			chRead <- []byte("quit")
			break
		}

		if messageType == 1 {
			incomingLogger.Info(
				"INCOMING MESSAGE",
				zap.Int("ChLength",
					len(chRead)),
				zap.String("message", string(message)))

			chRead <- message
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

func WriteFromChannelToWS(
	c *websocket.Conn,
	chWrite <-chan interface{},
	OutgoingLogger *zap.Logger,
	logger *zap.Logger,
	restartRequired *atomic.Bool,
	wg *sync.WaitGroup) {

	wg.Add(1)
	defer wg.Done()

	logger.Info("Socket writer started")
	defer func() {
		logger.Info("Socket writer closed")
	}()

	for {
		message := <-chWrite
		s, ok := message.(string)
		if ok {
			if s == "quit" {
				restartRequired.Store(true)
				break
			}
		}

		message, err := json.Marshal(message)
		if err != nil {
			restartRequired.Store(true)
			logger.Sugar().Error("Error in SocketWriter while Marshaling message", message, zap.Error(err))
			break
		}

		OutgoingLogger.Info("OUTGOING MESSAGE",
			zap.Int("ChLen", len(chWrite)),
			zap.String("message", string(message.([]byte))))

		err = c.WriteMessage(websocket.TextMessage, message.([]byte))

		if err != nil {
			logger.Error("Error while writing message to socket", zap.Error(err))
			restartRequired.Store(true)
			break
		}
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
