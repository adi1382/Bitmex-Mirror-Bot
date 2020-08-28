package subClient

import (
	"github.com/adi1382/Bitmex-Mirror-Bot/websocket"
	"go.uber.org/zap"
)

func (c *SubClient) connectToSocket() {

	c.logger.Info("Starting Connection on subClient to sockets",
		zap.String("apiKey", c.ApiKey),
		zap.String("websocketTopic", c.WebsocketTopic))

	var message []interface{}
	message = append(message, 1, c.ApiKey, c.WebsocketTopic)
	c.chWriteToWSClient <- message

	// Can't use this as m.subsLock unable to get unlocked
	//for {
	//	fmt.Println("checking")
	//	if c.isConnectedToSocket.Load() {
	//		break
	//	}
	//	time.Sleep(time.Second)
	//}

}

func (c *SubClient) authorize() {
	var message []interface{}

	c.logger.Info("Authenticating websocket connection of subClient",
		zap.String("apiKey", c.ApiKey),
		zap.String("websocketTopic", c.WebsocketTopic))

	message = append(message, 0, c.ApiKey, c.WebsocketTopic, websocket.GetAuthMessage(c.ApiKey, c.apiSecret))
	c.chWriteToWSClient <- message

	//for {
	//	if c.isAuthenticatedToSocket.Load() {
	//		break
	//	}
	//}
}

func (c *SubClient) subscribeTopics(tables ...string) {
	var message []interface{}
	command := websocket.Message{Op: "subscribe"}

	for _, v := range tables {
		command.AddArgument(v)
	}

	c.logger.Info("Subscribing Tables on SubClient",
		zap.String("apiKey", c.ApiKey),
		zap.String("websocketTopic", c.WebsocketTopic))

	message = append(message, 0, c.ApiKey, c.WebsocketTopic, command)
	c.chWriteToWSClient <- message
}

func (c *SubClient) unsubscribeTopics(tables ...string) {

	c.logger.Info("Unsubscribing Tables on SubClient",
		zap.String("apiKey", c.ApiKey),
		zap.String("websocketTopic", c.WebsocketTopic))

	var message []interface{}
	command := websocket.Message{Op: "unsubscribe"}

	for _, v := range tables {
		command.AddArgument(v)
	}

	message = append(message, 0, c.ApiKey, c.WebsocketTopic, command)
	c.chWriteToWSClient <- message
}
