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

	logger.Info("Ping Pong protocol Initiated")

	pongWait := time.Second * 10
	err := conn.SetReadDeadline(time.Now().Add(pongWait))

	if err != nil {
		logger.Error("Error for PingPong", zap.Error(conn.Close()))
		RestartRequired.Store(true)
		return
	}

	conn.SetPongHandler(func(string) error { err = conn.SetReadDeadline(time.Now().Add(pongWait)); return err })

	isPingSent := atomic.NewBool(false)

	pingSender := func() {
		socketWriterLock.Lock()
		err := conn.WriteMessage(websocket.PingMessage, []byte{})
		socketWriterLock.Unlock()
		isPingSent.Store(true)
		if err != nil {
			logger.Error("Error for PingPong, closing socket connection", zap.Error(conn.Close()))
			RestartRequired.Store(true)
		}
	}

	go func() {
		pingSender()
		timer := time.NewTimer(time.Nanosecond)
		for {
			if isPingSent.Load() {
				isPingSent.Store(false)
				timer = time.AfterFunc(time.Second*5, pingSender)
			}

			time.Sleep(time.Nanosecond)

			if RestartRequired.Load() {
				logger.Warn("Did timer stopped?", zap.Bool("PingPong Timer status", timer.Stop()))
				logger.Info("Closing Socket", zap.Error(conn.Close()))
				break
			}
		}
	}()
}
