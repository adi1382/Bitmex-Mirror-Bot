package hostClient

import "github.com/adi1382/Bitmex-Mirror-Bot/websocket"

func (c *HostClient) ActivePositions() websocket.PositionSlice {

	InfoLogger.Printf("Fetching active positons for subClient %s\n", c.ApiKey)

	c.positionsLock.Lock()
	defer c.positionsLock.Unlock()
	return c.activePositions
}
