package subClient

import (
	"fmt"
	"github.com/adi1382/Bitmex-Mirror-Bot/bitmex"
	"github.com/adi1382/Bitmex-Mirror-Bot/hostClient"
	"github.com/adi1382/Bitmex-Mirror-Bot/swagger"
	"github.com/adi1382/Bitmex-Mirror-Bot/websocket"
	"go.uber.org/atomic"
	"go.uber.org/zap"
	"os"
	"strings"
	"sync"
	"time"
)

func NewSubClient(
	apiKey, apiSecret string,
	test, balanceProportion bool,
	fixedRatio, marginUpdateTime, calibrationTime, limitFilledTimeout float64,
	ch chan<- interface{},
	RestartCounter *atomic.Uint32,
	hostClient *hostClient.HostClient,
	logger *zap.Logger) *SubClient {

	c := SubClient{
		ApiKey:    apiKey,
		apiSecret: apiSecret,
		test:      test,
	}
	if c.test {
		c.Rest = swagger.NewAPIClient(swagger.NewTestnetConfiguration())
	} else {
		c.Rest = swagger.NewAPIClient(swagger.NewConfiguration())
	}
	c.Rest.InitializeAuth(c.ApiKey, c.apiSecret)
	c.WebsocketTopic = ""

	c.restartCounter = RestartCounter
	c.active.Store(true)
	c.marginUpdateTime = marginUpdateTime
	c.chWriteToWSClient = ch
	c.chReadFromWSClient = make(chan []byte, 100)
	c.BalanceProportion = balanceProportion
	c.FixedRatio = fixedRatio
	c.calibrationTime = calibrationTime
	c.LimitFilledTimeout = limitFilledTimeout
	c.hostClient = hostClient
	c.logger = logger
	c.initiateRest()

	c.logger.Info("New SubClient Created",
		zap.String("websocketTopic", c.WebsocketTopic),
		zap.String("apiKey", apiKey))

	// balanceProportion bool, fixedRatio float64,
	//hostClient *SubClient, calibrationTime int64, LimitFilledTimeout int64

	return &c
}

type SubClient struct {
	active                  atomic.Bool
	isConnectedToSocket     bool
	isAuthenticatedToSocket bool
	marginUpdateTime        float64
	BalanceProportion       bool
	FixedRatio              float64
	test                    bool
	marginUpdated           atomic.Bool
	partials                atomic.Uint32
	marginBalance           atomic.Float64
	LimitFilledTimeout      float64
	activeOrders            websocket.OrderSlice
	activePositions         websocket.PositionSlice
	currentMargin           websocket.MarginSlice
	ordersLock              sync.Mutex
	positionsLock           sync.Mutex
	marginLock              sync.Mutex
	ApiKey                  string
	apiSecret               string
	WebsocketTopic          string
	chWriteToWSClient       chan<- interface{}
	chReadFromWSClient      chan []byte
	Rest                    *swagger.APIClient
	hostClient              *hostClient.HostClient
	calibrationTime         float64
	hostUpdatesFetcher      chan []byte
	restartCounter          *atomic.Uint32
	logger                  *zap.Logger
}

func (c *SubClient) Initialize() {

	c.logger.Info(
		"Initializing sub client",
		zap.String("websocketTopic", c.WebsocketTopic),
		zap.String("apiKey", c.ApiKey))

	c.hostUpdatesFetcher = make(chan []byte, 100)
	c.active.Store(true)
	c.chReadFromWSClient = make(chan []byte, 100)
	c.StartConnection()
	c.Authorize()
	go c.marginUpdate()
	go c.dataHandler()
	go c.OrderHandler()

	c.logger.Info(
		"sub client initialized",
		zap.String("websocketTopic", c.WebsocketTopic),
		zap.String("apiKey", c.ApiKey))
}

//////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////
func (c *SubClient) CloseConnection() {

	c.logger.Info("close connection request for sub client",
		zap.String("apiKey", c.ApiKey),
		zap.String("websocketTopic", c.WebsocketTopic))
	c.active.Store(false)
	var message []interface{}
	message = append(message, 2, c.ApiKey, c.WebsocketTopic)
	c.chWriteToWSClient <- message
	c.chReadFromWSClient <- []byte("quit")
	c.removeCurrentClient()

}

func (c *SubClient) DropConnection() {
	if c.RunningStatus() {
		c.logger.Info("drop connection request for sub client",
			zap.String("apiKey", c.ApiKey),
			zap.String("websocketTopic", c.WebsocketTopic))
		c.active.Store(false)
		c.chReadFromWSClient <- []byte("quit")
	}
}

//////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////

func (c *SubClient) WaitForPartial() {
	c.logger.Debug("waiting for partials",
		zap.String("apiKey", c.ApiKey),
		zap.String("websocketTopic", c.WebsocketTopic))

	for {
		if c.partials.Load() >= 3 {
			break
		}
	}
	c.logger.Debug("partials received",
		zap.String("apiKey", c.ApiKey),
		zap.String("websocketTopic", c.WebsocketTopic))
}

func (c *SubClient) marginUpdate() {

	defer func() {
		//fmt.Println("Margin Update Stopped for subClient ", c.ApiKey)
	}()
	c.marginUpdated.Store(false)
	c.WaitForPartial()
	//fmt.Println("Margin Update Started for subClient ", c.ApiKey)
	for {
		if !c.RunningStatus() {
			break
		}

		marginBalance := c.RestMargin()
		c.logger.Debug("Updating Margin Balance",
			zap.Float64("balance", marginBalance),
			zap.String("apiKey", c.ApiKey),
			zap.String("websocketTopic", c.WebsocketTopic))

		//c.marginBalance = c.currentMargin[0].MarginBalance.Value
		c.marginBalance.Store(marginBalance)
		c.marginUpdated.Store(true)

		c.logger.Debug("Margin Balance Updated",
			zap.Float64("balance", marginBalance),
			zap.String("apiKey", c.ApiKey),
			zap.String("websocketTopic", c.WebsocketTopic))

		//fmt.Println("Margin updated on ", c.ApiKey)
		// Calibration time
		//time.Sleep(time.Second * time.Duration(c.marginUpdateTime))

		time.Sleep(time.Second * 5)

		resetTime := time.Now().Add(time.Second * time.Duration(c.marginUpdateTime))
		for {
			time.Sleep(time.Nanosecond)
			if time.Now().Unix() > resetTime.Unix() {
				break
			} else if !c.RunningStatus() {
				break
			}
		}
	}
}

func (c *SubClient) GetMarginBalance() float64 {
	c.logger.Debug("Fetching Margin Balance",
		zap.String("apiKey", c.ApiKey),
		zap.String("websocketTopic", c.WebsocketTopic))
	for {
		if c.marginUpdated.Load() {
			return c.marginBalance.Load()
		} else {
			time.Sleep(time.Nanosecond)
		}
	}
}

func (c *SubClient) initiateRest() {

	c.logger.Info("Initiating Rest Object on subClient",
		zap.String("apiKey", c.ApiKey),
		zap.String("websocketTopic", c.WebsocketTopic))

	if c.test {
		c.Rest = swagger.NewAPIClient(swagger.NewTestnetConfiguration())
	} else {
		c.Rest = swagger.NewAPIClient(swagger.NewConfiguration())
	}
	c.Rest.InitializeAuth(c.ApiKey, c.apiSecret)

	c.logger.Info("Rest Object Initiated on subClient",
		zap.String("apiKey", c.ApiKey),
		zap.String("websocketTopic", c.WebsocketTopic))
}

func (c *SubClient) StartConnection() {

	c.logger.Info("Starting Connection on subClient to sockets",
		zap.String("apiKey", c.ApiKey),
		zap.String("websocketTopic", c.WebsocketTopic))

	var message []interface{}
	message = append(message, 1, c.ApiKey, c.WebsocketTopic)
	c.chWriteToWSClient <- message

}

func (c *SubClient) Authorize() {
	var message []interface{}

	c.logger.Info("Authenticating websocket connection of subClient",
		zap.String("apiKey", c.ApiKey),
		zap.String("websocketTopic", c.WebsocketTopic))

	message = append(message, 0, c.ApiKey, c.WebsocketTopic, websocket.GetAuthMessage(c.ApiKey, c.apiSecret))
	c.chWriteToWSClient <- message
}

func (c *SubClient) SubscribeTopics(tables ...string) {
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

func (c *SubClient) UnsubscribeTopics(tables ...string) {

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

func (c *SubClient) Push(message *[]byte) {
	c.chReadFromWSClient <- *message
}

func (c *SubClient) HostUpdatePush(message *[]byte) {
	c.hostUpdatesFetcher <- *message
}

func (c *SubClient) dataHandler() {
	//fmt.Println("Data Handler started for subClient ", c.ApiKey)
	defer func() {
		c.logger.Info("Data Handler Closed for subClient",
			zap.String("apiKey", c.ApiKey),
			zap.String("websocketTopic", c.WebsocketTopic))
		//fmt.Println("Data Handler Closed for subClient ", c.ApiKey)
	}()
	for {

		if !c.RunningStatus() {
			break
		}

		message := <-c.chReadFromWSClient

		c.logger.Debug("New Message in Data Handler for subClient",
			zap.String("apiKey", c.ApiKey),
			zap.String("websocketTopic", c.WebsocketTopic))

		strResponse := string(message)
		if strResponse == "quit" {
			break
		}

		if strings.Contains(string(message), "Access Token expired for subscription") {
			c.restartCounter.Add(1)

			c.logger.Error("Expiration Error",
				zap.String("errorMessage", string(message)),
				zap.String("apiKey", c.ApiKey),
				zap.String("websocketTopic", c.WebsocketTopic))
			//fmt.Println(string(message))
		}

		if strings.Contains(string(message), "Invalid API Key") {
			fmt.Println("API key ", c.ApiKey, " is invalid.")

			c.logger.Error("api key invalid for subClient",
				zap.String("errMessage", strResponse),
				zap.String("apiKey", c.ApiKey),
				zap.String("websocketTopic", c.WebsocketTopic))

			if c.WebsocketTopic == "hostAccount" {
				fmt.Println("host Account API key is Invalid. Closing the bot in 10 seconds.")
				c.logger.Error("host Account API key is Invalid. Closing the bot in 10 seconds.")
				time.Sleep(time.Second * 10)
				os.Exit(-1)
			} else {
				c.CloseConnection()
			}
		}

		if strings.Contains(string(message), "This key is disabled") {
			c.logger.Error("apiKey is disabled on subClient",
				zap.String("errorMessage", strResponse),
				zap.String("apiKey", c.ApiKey),
				zap.String("websocketTopic", c.WebsocketTopic))

			fmt.Println("API key ", c.ApiKey, " is disabled.")

			if c.WebsocketTopic == "hostAccount" {
				fmt.Println("host Account API key is disabled. Closing the bot in 10 seconds.")
				c.logger.Error("host Account API key is disabled. Closing the bot in 10 seconds.")
				time.Sleep(time.Second * 10)
				os.Exit(-1)
			} else {
				c.CloseConnection()
			}
		}

		prefix := fmt.Sprintf(`[0,"%s","%s",`, c.ApiKey, c.WebsocketTopic)
		suffix := fmt.Sprintf("]")
		strResponse = strings.TrimPrefix(strResponse, prefix)
		strResponse = strings.TrimSuffix(strResponse, suffix)
		if !strings.Contains(string(message), "table") {
			continue
		}

		response, table := bitmex.DecodeMessage([]byte(strResponse), c.logger)

		c.logger.Debug("Updating table on subClient",
			zap.String("table", table),
			zap.String("apiKey", c.ApiKey),
			zap.String("websocketTopic", c.WebsocketTopic))

		// Potential Race Condition when fetching
		if table == "order" {

			orderResponse := response.(bitmex.OrderResponse)

			c.ordersLock.Lock()
			if orderResponse.Action == "partial" {
				c.partials.Add(1)
				c.activeOrders.OrderPartial(&orderResponse.Data)
			} else if orderResponse.Action == "insert" {
				c.activeOrders.OrderInsert(&orderResponse.Data)
			} else if orderResponse.Action == "update" {
				c.activeOrders.OrderUpdate(&orderResponse.Data)
			} else if orderResponse.Action == "delete" {
				c.activeOrders.OrderDelete(&orderResponse.Data)
			}
			c.ordersLock.Unlock()

		} else if table == "position" {
			positionResponse := response.(bitmex.PositionResponse)

			c.positionsLock.Lock()
			if positionResponse.Action == "partial" {
				c.partials.Add(1)
				c.activePositions.PositionPartial(&positionResponse.Data)
			} else if positionResponse.Action == "insert" {
				c.activePositions.PositionInsert(&positionResponse.Data)
			} else if positionResponse.Action == "update" {
				c.activePositions.PositionUpdate(&positionResponse.Data)
			} else if positionResponse.Action == "delete" {
				c.activePositions.PositionDelete(&positionResponse.Data)
			}
			c.positionsLock.Unlock()

		} else if table == "margin" {
			marginResponse := response.(bitmex.MarginResponse)

			c.marginLock.Lock()
			if marginResponse.Action == "partial" {
				c.partials.Add(1)
				c.currentMargin.MarginPartial(&marginResponse.Data)
			} else if marginResponse.Action == "insert" {
				c.currentMargin.MarginInsert(&marginResponse.Data)
			} else if marginResponse.Action == "update" {
				c.currentMargin.MarginUpdate(&marginResponse.Data)
			} else if marginResponse.Action == "delete" {
				c.currentMargin.MarginDelete(&marginResponse.Data)
			}
			c.marginLock.Unlock()
		}
	}
}

func (c *SubClient) RunningStatus() bool {
	return c.active.Load()
}

func (c *SubClient) removeCurrentClient() {

	c.logger.Info("Removing subClient from all clients",
		zap.String("apiKey", c.ApiKey),
		zap.String("websocketTopic", c.WebsocketTopic))

}

//func GetAllClients() []*SubClient {
//	AllClientsLock.Lock()
//	defer AllClientsLock.Unlock()
//	return AllClients
//}

//func ResetAllClients() {
//	AllClientsLock.Lock()
//	defer AllClientsLock.Unlock()
//
//	fmt.Println("Removing all clients")
//	//for i := range AllClients {
//	//	AllClients[i].closeConnection()
//	//}
//	fmt.Println("All clients removed")
//
//	AllClients = AllClients[:0]
//}
