package Mirror

import (
	"encoding/json"
	"fmt"
	"github.com/adi1382/Bitmex-Mirror-Bot/hostClient"
	"github.com/adi1382/Bitmex-Mirror-Bot/subClient"
	"go.uber.org/atomic"
	"strings"
	"sync"
	"time"
)

type Mirror struct {
	Host           *hostClient.HostClient
	Subs           []*subClient.SubClient
	RestartCounter *atomic.Uint32
	subLock        sync.Mutex
}

func (m *Mirror) SocketResponseDistributor(c <-chan []byte, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	defer func() {
		fmt.Println("Socket Distributor Closed")
	}()

	defer func() {
		m.Host.CloseConnection()
		for _, v := range m.Subs {
			v.CloseConnection()
		}
	}()

	for {
		message := <-c

		if m.RestartCounter.Load() > 0 {
			break
		}

		if string(message) == "quit" {
			m.RestartCounter.Add(1)
			break
		}

		var u []interface{}

		err := json.Unmarshal(message, &u)

		if err != nil {
			m.RestartCounter.Add(1)
			break
		}

		//tools.CheckErr(err)

		key := u[1].(string)

		socketTopic := u[2].(string)

		go func() {
			for _, v := range m.Subs {

				if !v.RunningStatus() {
					continue
				}

				if !(strings.Contains(string(message), `"table":"instrument"`) || !strings.Contains(string(message), "table")) {
					if socketTopic == "hostAccount" {
						go v.HostUpdatePush(&message)
					}
				}
			}
		}()

		go func() {
			if socketTopic == "hostAccount" {
				m.Host.Push(&message)
			} else {
				for _, v := range m.Subs {

					if !v.RunningStatus() {
						continue
					}

					if v.ApiKey == key {
						v.Push(&message)
						break
					}
				}
			}
		}()
	}
}

func (m *Mirror) SubChecker() {
	go func() {
		for {
			m.remover()

			if m.RestartCounter.Load() > 0 {
				break
			}

			time.Sleep(time.Second * 5)
		}
	}()
}

func (m *Mirror) remover() {
	for i := range m.Subs {
		if !m.Subs[i].RunningStatus() {
			m.Subs = append(m.Subs[:i], m.Subs[i+1:]...)
			break
		}
	}
}

//func (m *Mirror) SocketResponseDistributor(c <-chan []byte, RestartCounter *atomic.Uint32, wg *sync.WaitGroup) {
//	wg.Add(1)
//	defer wg.Done()
//L:
//	for {
//		time.Sleep(time.Nanosecond)
//
//		select {
//		case message := <-c:
//			if RestartCounter.Load() > 0 {
//				//subClient.AllClientsLock.Lock()
//				for _, v := range m.Subs {
//					v.DropConnection()
//				}
//				m.Host.CloseConnection()
//				m.Subs = make([]*subClient.SubClient, 0, 5)
//				//subClient.AllClientsLock.Unlock()
//				RestartCounter.Add(1)
//				subClient.InfoLogger.Println("Distributor closed")
//				break L
//			}
//
//			var u []interface{}
//
//			err := json.Unmarshal(message, &u)
//
//			tools.CheckErr(err)
//
//			key := u[1].(string)
//
//			socketTopic := u[2].(string)
//
//			go func() {
//				//subClient.AllClientsLock.Lock()
//				for _, v := range m.Subs {
//					if v.WebsocketTopic == "hostAccount" {
//						continue
//					}
//					if !(strings.Contains(string(message), `"table":"instrument"`) || !strings.Contains(string(message), "table")) {
//						if socketTopic == "hostAccount" {
//							go v.HostUpdatePush(&message)
//						}
//					}
//				}
//				//subClient.AllClientsLock.Unlock()
//			}()
//
//			go func() {
//				//subClient.AllClientsLock.Lock()
//				if socketTopic == "hostAccount" {
//					m.Host.Push(&message)
//				} else {
//					for _, v := range m.Subs {
//						if v.ApiKey == key {
//							v.Push(&message)
//							break
//						}
//					}
//				}
//				//subClient.AllClientsLock.Unlock()
//			}()
//		default:
//			if RestartCounter.Load() > 0 {
//				//subClient.AllClientsLock.Lock()
//				for _, v := range m.Subs {
//					v.DropConnection()
//				}
//				m.Host.CloseConnection()
//				m.Subs = make([]*subClient.SubClient, 0, 5)
//				//subClient.AllClientsLock.Unlock()
//				RestartCounter.Add(1)
//				subClient.InfoLogger.Println("Distributor closed")
//				break L
//			}
//
//		}
//	}
//}
