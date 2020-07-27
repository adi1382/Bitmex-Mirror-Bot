package hostClient

import "github.com/adi1382/Bitmex-Mirror-Bot/websocket"

func (c *HostClient) ActiveOrders() websocket.OrderSlice {
	c.ordersLock.Lock()
	defer c.ordersLock.Unlock()
	return c.activeOrders
}
