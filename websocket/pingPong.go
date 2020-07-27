package websocket

import (
	"github.com/adi1382/Bitmex-Mirror-Bot/tools"
	"github.com/gorilla/websocket"
	"go.uber.org/atomic"
	"time"
)

func PingPong(conn *websocket.Conn, RestartCounter *atomic.Uint32) {

	InfoLogger.Println("Ping Pong protocol Initiated")

	pongWait := time.Second * 10
	err := conn.SetReadDeadline(time.Now().Add(pongWait))
	tools.WebsocketErr(err, RestartCounter)
	conn.SetPongHandler(func(string) error { err = conn.SetReadDeadline(time.Now().Add(pongWait)); return err })

	go func() {
		for {
			if RestartCounter.Load() > 0 {
				//fmt.Println("Ping Pong Closed")
				RestartCounter.Add(1)
				break
			}
			err := conn.WriteMessage(9, []byte{})
			if err != nil {
				RestartCounter.Add(1)
				break
			}
			resetTime := time.Now().Add(time.Second * 5)
			for {
				time.Sleep(time.Nanosecond)
				if time.Now().Unix() > resetTime.Unix() {
					break
				} else {
					if RestartCounter.Load() > 0 {
						break
					}
				}
			}
			//time.Sleep(time.Second*5)
		}
	}()
}
