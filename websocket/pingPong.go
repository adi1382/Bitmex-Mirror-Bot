package websocket

import (
	"github.com/gorilla/websocket"
	"go.uber.org/atomic"
	"go.uber.org/zap"
	"sync"
	"time"
)

func PingPong(conn *websocket.Conn, RestartRequired *atomic.Bool, logger *zap.Logger, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	InfoLogger.Println("Ping Pong protocol Initiated")

	pongWait := time.Second * 10
	err := conn.SetReadDeadline(time.Now().Add(pongWait))

	if err != nil {
		logger.Error("Error for PingPong", zap.Error(conn.Close()))
		RestartRequired.Store(true)
		return
	}

	conn.SetPongHandler(func(string) error { err = conn.SetReadDeadline(time.Now().Add(pongWait)); return err })

	go func() {
		for {
			if RestartRequired.Load() {
				break
			}

			err := conn.WriteMessage(9, []byte{})
			if err != nil {
				logger.Error("Error for PingPong", zap.Error(conn.Close()))
				RestartRequired.Store(true)
				break
			}
			resetTime := time.Now().Add(time.Second * 5)
			for {
				time.Sleep(time.Nanosecond)
				if time.Now().Unix() > resetTime.Unix() {
					break
				} else if RestartRequired.Load() {
					break
				}
			}
		}
	}()
}
