package hostClient

import (
	"github.com/adi1382/Bitmex-Mirror-Bot/websocket"
	"go.uber.org/zap"
)

func (c *HostClient) ActivePositions() websocket.PositionSlice {

	c.logger.Debug("Fetching active positions for hostClient %s\n",
		zap.String("apiKey", c.ApiKey),
		zap.String("websocketTopic", c.WebsocketTopic))

	c.positionsLock.Lock()
	defer c.positionsLock.Unlock()
	return c.activePositions
}
