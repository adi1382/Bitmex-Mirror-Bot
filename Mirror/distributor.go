package Mirror

import (
	"encoding/json"
	"go.uber.org/zap"
	"strings"
)

func (m *Mirror) SocketResponseDistributor(c <-chan []byte) {
	m.wg.Add(1)
	defer m.wg.Done()

	defer func() {
		m.logger.Info("Socket Distributor Closed")
	}()

	defer func() {
		m.host.CloseConnection()

		m.subsLock.Lock()
		for _, v := range m.subs {
			v.CloseConnection("socket distributor closed")
		}
		m.subsLock.Unlock()
	}()

	for {
		message := <-c
		if m.restartRequired.Load() {
			break
		}

		if string(message) == "quit" {
			m.logger.Info("Close message received in socket distributor")
			m.restartRequired.Store(true)
			break
		}

		var u []interface{}

		err := json.Unmarshal(message, &u)

		if err != nil {

			m.logger.Error("Error while performing Unmarshal in socket distributor",
				zap.String("message", string(message)), zap.Error(err))

			m.restartRequired.Store(true)
			break
		}

		//tools.CheckErr(err)

		key := u[1].(string)

		socketTopic := u[2].(string)

		go func() {
			m.subsLock.Lock()
			for _, v := range m.subs {

				if !v.RunningStatus() {
					continue
				}

				if !(strings.Contains(string(message), `"table":"instrument"`) || !strings.Contains(string(message), "table")) {
					if socketTopic == "hostAccount" {
						go v.HostUpdatePush(&message)
					}
				}
			}
			m.subsLock.Unlock()
		}()

		go func() {
			if socketTopic == "hostAccount" {
				m.host.Push(&message)
			} else {
				m.subsLock.Lock()
				for _, v := range m.subs {

					if !v.RunningStatus() {
						continue
					}

					if v.ApiKey == key {
						v.Push(&message)
						break
					}
				}
				m.subsLock.Unlock()
			}
		}()
	}
}
