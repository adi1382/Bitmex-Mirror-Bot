package Mirror

import (
	"encoding/json"
	"github.com/adi1382/Bitmex-Mirror-Bot/hostClient"
	"github.com/adi1382/Bitmex-Mirror-Bot/subClient"
	"github.com/adi1382/Bitmex-Mirror-Bot/tools"
	"go.uber.org/atomic"
	"strings"
	"sync"
	"time"
)

type Mirror struct {
	Host *hostClient.HostClient
	Subs []*subClient.SubClient
}

func (m *Mirror) SocketResponseDistributor(c <-chan []byte, RestartCounter *atomic.Uint32, wg *sync.WaitGroup) {
L:
	for {
		time.Sleep(time.Nanosecond)

		select {
		case message := <-c:
			if RestartCounter.Load() > 0 {
				//subClient.AllClientsLock.Lock()
				for _, v := range m.Subs {
					v.DropConnection()
				}
				m.Host.CloseConnection()
				m.Subs = make([]*subClient.SubClient, 0, 5)
				//subClient.AllClientsLock.Unlock()
				RestartCounter.Add(1)
				subClient.InfoLogger.Println("Distributor closed")
				break L
			}

			var u []interface{}

			err := json.Unmarshal(message, &u)

			tools.CheckErr(err)

			key := u[1].(string)

			socketTopic := u[2].(string)

			go func() {
				//subClient.AllClientsLock.Lock()
				for _, v := range m.Subs {
					if v.WebsocketTopic == "hostAccount" {
						continue
					}
					if !(strings.Contains(string(message), `"table":"instrument"`) || !strings.Contains(string(message), "table")) {
						if socketTopic == "hostAccount" {
							go v.HostUpdatePush(&message)
						}
					}
				}
				//subClient.AllClientsLock.Unlock()
			}()

			go func() {
				//subClient.AllClientsLock.Lock()
				if socketTopic == "hostAccount" {
					m.Host.Push(&message)
				} else {
					for _, v := range m.Subs {
						if v.ApiKey == key {
							v.Push(&message)
							break
						}
					}
				}
				//subClient.AllClientsLock.Unlock()
			}()
		default:
			if RestartCounter.Load() > 0 {
				//subClient.AllClientsLock.Lock()
				for _, v := range m.Subs {
					v.DropConnection()
				}
				m.Host.CloseConnection()
				m.Subs = make([]*subClient.SubClient, 0, 5)
				//subClient.AllClientsLock.Unlock()
				RestartCounter.Add(1)
				subClient.InfoLogger.Println("Distributor closed")
				break L
			}

		}
	}
	wg.Done()
}
