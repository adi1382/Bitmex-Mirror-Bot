package subClient

import (
	"fmt"
	"github.com/adi1382/Bitmex-Mirror-Bot/swagger"
	"github.com/adi1382/Bitmex-Mirror-Bot/websocket"
)

func (c *SubClient) CurrentMargin() websocket.MarginSlice {
	InfoLogger.Println("Fetching Current margin for subClient ", c.ApiKey)
	c.marginLock.Lock()
	defer c.marginLock.Unlock()
	return c.currentMargin
}

func (c *SubClient) RestMargin() float64 {
	InfoLogger.Println("Fetching Current margin for subClient ", c.ApiKey)
	var currency swagger.UserGetMarginOpts

L:
	for {
		Margin, response, err := c.Rest.UserApi.UserGetMargin(&currency)
		switch c.SwaggerError(err, response) {
		case 0:
			return Margin.MarginBalance.Value
		case 1:
			continue L
		case 2:
			fmt.Println("Remove the current sub subClient")
			c.CloseConnection()
			return -404
			//break function
		case 3:
			fmt.Println("Restart the bot")
			return -404
		}

	}

}
