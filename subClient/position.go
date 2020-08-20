package subClient

import (
	"github.com/adi1382/Bitmex-Mirror-Bot/tools"
	"github.com/adi1382/Bitmex-Mirror-Bot/websocket"
	"go.uber.org/zap"
)

func (c *SubClient) UpdateLeverage(symbol string, leverage float64) {

	c.logger.Info("Updating leverage via rest for subClient",
		zap.String("symbol", symbol),
		zap.Float64("leverage", leverage),
		zap.String("apiKey", c.ApiKey),
		zap.String("websocketTopic", c.WebsocketTopic))

	res, _, err := c.Rest.PositionApi.PositionUpdateLeverage(symbol, leverage)
	tools.CheckErr(err)

	c.positionsLock.Lock()
	defer c.positionsLock.Unlock()

	for idx := range c.activePositions {
		if res.Symbol.Value == c.activePositions[idx].Symbol.Value {
			c.activePositions[idx].Leverage.Value = res.Leverage.Value
		}
	}

	tools.CheckErr(err)
}

func (c *SubClient) ActivePositions() websocket.PositionSlice {

	c.logger.Debug("Fetching active positions for subClient",
		zap.String("apiKey", c.ApiKey),
		zap.String("websocketTopic", c.WebsocketTopic))

	c.positionsLock.Lock()
	defer c.positionsLock.Unlock()
	return c.activePositions
}

func (c *SubClient) TransferMargin(symbol string, margin int) {

	c.logger.Info("Transferring Margin Via Rest",
		zap.String("symbol", symbol),
		zap.Int("margin", margin),
		zap.String("apiKey", c.ApiKey),
		zap.String("websocketTopic", c.WebsocketTopic))

	res, _, err := c.Rest.PositionApi.PositionTransferIsolatedMargin(symbol, margin)
	tools.CheckErr(err)

	c.positionsLock.Lock()
	defer c.positionsLock.Unlock()

	for idx := range c.activePositions {
		if c.activePositions[idx].Symbol.Value == res.Symbol.Value {
			c.activePositions[idx].PosMargin.Value = res.PosMargin.Value
		}
	}

}
