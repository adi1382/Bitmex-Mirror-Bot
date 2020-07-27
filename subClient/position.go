package subClient

import (
	"github.com/adi1382/Bitmex-Mirror-Bot/tools"
	"github.com/adi1382/Bitmex-Mirror-Bot/websocket"
)

func (c *SubClient) UpdateLeverage(symbol string, leverage float64) {

	InfoLogger.Printf("Updating leverage of %s to %f for subClient %s\n", symbol, leverage, c.ApiKey)

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

	InfoLogger.Printf("Fetching active positons for subClient %s\n", c.ApiKey)

	c.positionsLock.Lock()
	defer c.positionsLock.Unlock()
	return c.activePositions
}

func (c *SubClient) TransferMargin(symbol string, margin int) {

	InfoLogger.Printf("Transfering margin on %s by %d for subClient %s\n", symbol, margin, c.ApiKey)

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
