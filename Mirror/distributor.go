package Mirror

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"
)

func (m *Mirror) SocketResponseDistributor(c <-chan []byte, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	defer func() {
		fmt.Println("Socket Distributor Closed")
	}()

	defer func() {
		m.Host.CloseConnection()

		m.subLock.Lock()
		for _, v := range m.Subs {
			v.CloseConnection()
		}
		m.subLock.Unlock()
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
			m.subLock.Lock()
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
			m.subLock.Unlock()
		}()

		go func() {
			if socketTopic == "hostAccount" {
				m.Host.Push(&message)
			} else {
				m.subLock.Lock()
				for _, v := range m.Subs {

					if !v.RunningStatus() {
						continue
					}

					if v.ApiKey == key {
						v.Push(&message)
						break
					}
				}
				m.subLock.Unlock()
			}
		}()
	}
}
