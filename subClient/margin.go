package subClient

import (
	"fmt"
	"github.com/adi1382/Bitmex-Mirror-Bot/swagger"
	"github.com/adi1382/Bitmex-Mirror-Bot/websocket"
	"go.uber.org/zap"
)

func (c *SubClient) CurrentMargin() websocket.MarginSlice {
	c.logger.Debug("Fetching Current margin for subClient",
		zap.String("apiKey", c.ApiKey),
		zap.String("websocketTopic", c.WebsocketTopic))

	c.marginLock.Lock()
	defer c.marginLock.Unlock()
	return c.currentMargin
}

func (c *SubClient) RestMargin() float64 {
	c.logger.Debug("Updating current margin via rest request",
		zap.String("apiKey", c.ApiKey),
		zap.String("websocketTopic", c.WebsocketTopic))
	var currency swagger.UserGetMarginOpts

L:
	for {
		Margin, response, err := c.Rest.UserApi.UserGetMargin(&currency)

		if err != nil {
			c.logger.Sugar().Error("Error while fetching margin",
				zap.Error(err),
				response)
		}

		switch c.SwaggerError(err, response) {
		case 0:
			return Margin.MarginBalance.Value
		case 1:
			continue L
		case 2:
			fmt.Println("Remove the current sub subClient")
			c.CloseConnection("Error while fetching margin")
			return -404
			//break function
		case 3:
			fmt.Println("Restart the bot")
			return -404
		}

	}

}
