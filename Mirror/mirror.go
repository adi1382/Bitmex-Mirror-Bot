package Mirror

import (
	"github.com/adi1382/Bitmex-Mirror-Bot/hostClient"
	"github.com/adi1382/Bitmex-Mirror-Bot/subClient"
	"go.uber.org/atomic"
	"go.uber.org/zap"
	"sync"
)

type Mirror struct {
	host            *hostClient.HostClient
	subs            []*subClient.SubClient
	restartRequired *atomic.Bool
	logger          *zap.Logger
	subsLock        sync.Mutex
	hostLock        sync.Mutex
	wg              *sync.WaitGroup
}

func NewMirror(restartRequired *atomic.Bool, logger *zap.Logger, wg *sync.WaitGroup) *Mirror {
	mirror := Mirror{
		restartRequired: restartRequired,
		logger:          logger,
		wg:              wg,
	}
	return &mirror
}

func (m *Mirror) StartMirroring() {
	m.hostLock.Lock()
	m.subsLock.Lock()
	defer m.subsLock.Unlock()
	defer m.hostLock.Unlock()

	m.host.SubscribeTopics("order", "position", "margin")
	for i := range m.subs {
		m.subs[i].SubscribeTopics("order", "position", "margin")
	}
}

func (m *Mirror) InitializeAll() {
	m.InitializeHost()
	m.InitializeSubs()
}

func (m *Mirror) InitializeHost() {
	m.host.Initialize()
}

func (m *Mirror) InitializeSubs() {
	m.subsLock.Lock()
	defer m.subsLock.Unlock()

	for i := range m.subs {
		m.subs[i].Initialize()
	}
}

func (m *Mirror) SetHost(host *hostClient.HostClient) {
	m.hostLock.Lock()
	defer m.hostLock.Unlock()
	m.host = host
}

func (m *Mirror) AddSub(sub *subClient.SubClient) {
	m.subsLock.Lock()
	defer m.subsLock.Unlock()
	m.subs = append(m.subs, sub)
}

func (m *Mirror) AddAndStartSub(sub *subClient.SubClient) {
	m.subsLock.Lock()
	defer m.subsLock.Unlock()
	sub.Initialize()
	sub.SubscribeTopics("order", "position", "margin")
	m.subs = append(m.subs, sub)
}

//func (m *Mirror) SubChecker() {
//	go func() {
//		for {
//			if m.RestartCounter.Load() > 0 {
//				break
//			}
//
//			m.remover()
//
//			time.Sleep(time.Second * 5)
//		}
//	}()
//}

func (m *Mirror) remover() {
	m.subsLock.Lock()
	for i := range m.subs {
		if !m.subs[i].RunningStatus() {
			m.subs = append(m.subs[:i], m.subs[i+1:]...)
			break
		}
	}
	m.subsLock.Unlock()
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
//				for _, v := range m.subs {
//					v.DropConnection()
//				}
//				m.host.CloseConnection()
//				m.subs = make([]*subClient.SubClient, 0, 5)
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
//				for _, v := range m.subs {
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
//					m.host.Push(&message)
//				} else {
//					for _, v := range m.subs {
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
//				for _, v := range m.subs {
//					v.DropConnection()
//				}
//				m.host.CloseConnection()
//				m.subs = make([]*subClient.SubClient, 0, 5)
//				//subClient.AllClientsLock.Unlock()
//				RestartCounter.Add(1)
//				subClient.InfoLogger.Println("Distributor closed")
//				break L
//			}
//
//		}
//	}
//}
