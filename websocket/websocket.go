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
	"net/url"
	"sync"
	"time"
)

var socketWriterLock sync.Mutex

type Message struct {
	Op   string        `json:"op,omitempty"`
	Args []interface{} `json:"args,omitempty"`
}

func (m *Message) AddArgument(argument string) {
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
		logger.Info("Connecting to socket",
			zap.String("url", u.String()))
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
}

func ReadFromWSToChannel(
	c *websocket.Conn,
	chRead chan<- []byte,
	incomingLogger, logger *zap.Logger,
	RestartRequired *atomic.Bool,
	wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

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

func WriteFromChannelToWS(
	c *websocket.Conn,
	chWrite <-chan interface{},
	OutgoingLogger, logger *zap.Logger,
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

		socketWriterLock.Lock()
		err = c.WriteMessage(websocket.TextMessage, message.([]byte))
		socketWriterLock.Unlock()

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
	sig.Write([]byte(fmt.Sprintf("GET/realtime%d", timestamp)))

	signature := hex.EncodeToString(sig.Sum(nil))

	var msgKey []interface{}
	msgKey = append(msgKey, key, timestamp, signature)

	return Message{"authKeyExpires", msgKey}
}
