package hostClient

import (
	"fmt"
	"github.com/adi1382/Bitmex-Mirror-Bot/bitmex"
	"github.com/adi1382/Bitmex-Mirror-Bot/swagger"
	"github.com/adi1382/Bitmex-Mirror-Bot/websocket"
	"go.uber.org/atomic"
	"go.uber.org/zap"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func NewHostClient(apiKey, apiSecret string, test bool, ch chan<- interface{}, marginUpdateTime int64,
	RestartCounter *atomic.Uint32, logger *zap.Logger) *HostClient {
	c := HostClient{
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
	c.WebsocketTopic = "hostAccount"
	c.restartCounter = RestartCounter

	c.active.Store(true)

	c.marginUpdateTime = marginUpdateTime
	c.chWriteToWSClient = ch
	c.chReadFromWSClient = make(chan []byte, 100)
	c.logger = logger

	c.logger.Info("New SubClient Created",
		zap.String("websocketTopic", c.WebsocketTopic),
		zap.String("apiKey", apiKey))

	return &c
}

type HostClient struct {
	active             atomic.Bool
	marginUpdateTime   int64
	test               bool
	marginUpdated      atomic.Bool
	partials           atomic.Uint32
	marginBalance      atomic.Float64
	activeOrders       websocket.OrderSlice
	activePositions    websocket.PositionSlice
	currentMargin      websocket.MarginSlice
	ordersLock         sync.Mutex
	positionsLock      sync.Mutex
	marginLock         sync.Mutex
	ApiKey             string
	apiSecret          string
	WebsocketTopic     string
	chWriteToWSClient  chan<- interface{}
	chReadFromWSClient chan []byte
	Rest               *swagger.APIClient
	restartCounter     *atomic.Uint32
	logger             *zap.Logger
}

func (c *HostClient) Initialize() {
	c.startSocketConnection()
	c.socketAuthentication()
	go c.marginUpdate()
	go c.dataHandler()

	c.logger.Info("New hostClient Initialized",
		zap.String("apiKey", c.ApiKey),
		zap.String("websocketTopic", c.WebsocketTopic))
}

func (c *HostClient) CloseConnection() {
	c.active.Store(false)
	var message []interface{}
	message = append(message, 2, c.ApiKey, c.WebsocketTopic)
	c.chWriteToWSClient <- message
	c.chReadFromWSClient <- []byte("quit")
	c.restartCounter.Add(1)

	c.logger.Info("Closed Connection for hostClient",
		zap.String("apiKey", c.ApiKey),
		zap.String("websocketTopic", c.WebsocketTopic))
}

func (c *HostClient) WaitForPartial() {
	for {

		if c.partials.Load() >= 3 {
			break
		} else if !c.RunningStatus() {
			break
		}
		time.Sleep(time.Nanosecond * 1)
	}
}

func (c *HostClient) marginUpdate() {

	c.marginUpdated.Store(false)

	for {
		if !c.RunningStatus() {
			break
		}

		marginBalance := c.restMargin()

		c.logger.Info("Updating marginBalance on hostAccount",
			zap.Float64("marginBalance", marginBalance),
			zap.String("apiKey", c.ApiKey),
			zap.String("websocketTopic", c.WebsocketTopic))

		c.marginBalance.Store(marginBalance)

		c.marginUpdated.Store(true)

		resetTime := time.Now().Add(time.Second * time.Duration(c.marginUpdateTime))

		time.Sleep(time.Second * 5)

		for {
			time.Sleep(time.Second * 5)
			if time.Now().Unix() > resetTime.Unix() {
				break
			} else if !c.RunningStatus() {
				break
			}
		}
	}
}

func (c *HostClient) GetMarginBalance() float64 {

	for {
		if c.marginUpdated.Load() {
			return c.marginBalance.Load()
		} else if !c.RunningStatus() {
			return c.marginBalance.Load()
		}
		time.Sleep(time.Nanosecond * 1)
	}
}

func (c *HostClient) startSocketConnection() {
	c.logger.Info("Starting socket connection on hostClient",
		zap.String("apiKey", c.ApiKey),
		zap.String("websocketTopic", c.WebsocketTopic))
	var message []interface{}
	message = append(message, 1, c.ApiKey, c.WebsocketTopic)
	c.chWriteToWSClient <- message
}

func (c *HostClient) socketAuthentication() {
	var message []interface{}

	c.logger.Info("Authenticating websocket connection on hostClient",
		zap.String("apiKey", c.ApiKey),
		zap.String("websocketTopic", c.WebsocketTopic))

	message = append(message, 0, c.ApiKey, c.WebsocketTopic, websocket.GetAuthMessage(c.ApiKey, c.apiSecret))
	c.chWriteToWSClient <- message
}

func (c *HostClient) SubscribeTopics(tables ...string) {
	var message []interface{}
	command := websocket.Message{Op: "subscribe"}

	for _, v := range tables {
		command.AddArgument(v)
	}

	c.logger.Info("Subscribing Tables on hostClient",
		zap.String("apiKey", c.ApiKey),
		zap.String("websocketTopic", c.WebsocketTopic))

	message = append(message, 0, c.ApiKey, c.WebsocketTopic, command)
	c.chWriteToWSClient <- message
}

func (c *HostClient) UnsubscribeTopics(tables ...string) {

	c.logger.Info("Unsubscribing Tables on hostClient",
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

func (c *HostClient) Push(message *[]byte) {
	c.chReadFromWSClient <- *message
}

func (c *HostClient) dataHandler() {
	//fmt.Println("Data Handler started for subClient ", c.ApiKey)
	defer func() {
		c.logger.Info("Data Handler Closed for subClient ",
			zap.String("apiKey", c.ApiKey),
			zap.String("websocketTopic", c.WebsocketTopic))
		//fmt.Println("Data Handler Closed for subClient ", c.ApiKey)
	}()
	for {

		if !c.RunningStatus() {
			break
		}

		message := <-c.chReadFromWSClient
		c.logger.Debug("Received new message in Data Handler for subClient ",
			zap.String("apiKey", c.ApiKey),
			zap.String("websocketTopic", c.WebsocketTopic))

		strResponse := string(message)
		if strResponse == "quit" {
			break
		}

		if strings.Contains(strResponse, "Access Token expired for subscription") {
			c.restartCounter.Add(1)
			//atomic.AddInt64(c.restartCounter, 1)

			c.logger.Error("Expiration Error",
				zap.String("errorMessage", string(message)),
				zap.String("apiKey", c.ApiKey),
				zap.String("websocketTopic", c.WebsocketTopic))

			//fmt.Println(string(message))
		}

		if strings.Contains(strResponse, "Invalid API Key") {
			fmt.Println("API key ", c.ApiKey, " is invalid.")

			c.logger.Error("api key invalid for hostCLient",
				zap.String("errMessage", strResponse),
				zap.String("apiKey", c.ApiKey),
				zap.String("websocketTopic", c.WebsocketTopic))

			time.Sleep(time.Second * 10)
			os.Exit(-1)
		}

		if strings.Contains(strResponse, "This key is disabled") {

			c.logger.Error("apiKey is disabled on hostClient",
				zap.String("errorMessage", strResponse),
				zap.String("apiKey", c.ApiKey),
				zap.String("websocketTopic", c.WebsocketTopic))

			time.Sleep(time.Second * 10)
			os.Exit(-1)
		}

		prefix := fmt.Sprintf(`[0,"%s","%s",`, c.ApiKey, c.WebsocketTopic)
		suffix := fmt.Sprintf("]")
		strResponse = strings.TrimPrefix(strResponse, prefix)
		strResponse = strings.TrimSuffix(strResponse, suffix)
		if !strings.Contains(string(message), "table") {
			continue
		}

		response, table := bitmex.DecodeMessage([]byte(strResponse), c.logger)

		c.logger.Debug("Updating table on hostClient",
			zap.String("table", table),
			zap.String("apiKey", c.ApiKey),
			zap.String("websocketTopic", c.WebsocketTopic))

		// Potential Race Condition when fetching
		if table == "order" {

			orderResponse := response.(bitmex.OrderResponse)

			c.ordersLock.Lock()
			if orderResponse.Action == "partial" {
				c.partials.Add(1)
				//atomic.AddInt64(&(c.partials), 1)
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

func (c *HostClient) RunningStatus() bool {
	return c.active.Load()
}

func (c *HostClient) CurrentMargin() websocket.MarginSlice {

	c.logger.Debug("Fetching Current margin for hostAccount ",
		zap.String("apiKey", c.ApiKey),
		zap.String("websocketTopic", c.WebsocketTopic))

	c.marginLock.Lock()
	defer c.marginLock.Unlock()
	return c.currentMargin
}

func (c *HostClient) restMargin() float64 {

	c.logger.Info("Updating current margin on hostClient",
		zap.String("apiKey", c.ApiKey),
		zap.String("websocketTopic", c.WebsocketTopic))

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
			c.logger.Error("API key Invalid/Disabled on host")
			fmt.Println("API key Invalid/Disabled on host, closing in 10 seconds")
			time.Sleep(time.Second * 10)
			os.Exit(-1)
			//c.CloseConnection()
			//return -404
			//break function
		case 3:
			fmt.Println("Restart the bot")
			return -404
		}

	}

}

func (c *HostClient) SwaggerError(err error, response *http.Response) int {

	if err != nil {

		//fmt.Println(err)
		c.logger.Error("Error on subClient",
			zap.String("apiKey", c.ApiKey),
			zap.String("websocketTopic", c.WebsocketTopic))

		if strings.Contains(err.Error(), "401") || strings.Contains(err.Error(), "403") {
			return 2
		}

		k, ok := err.(swagger.GenericSwaggerError)
		if ok {
			k, ok := k.Model().(swagger.ModelError)

			if ok {
				e := k.Error_
				c.logger.Sugar().Error(e.Message.Value, "///", e.Name.Value)
				c.logger.Sugar().Error(string(err.(swagger.GenericSwaggerError).Body()))
				c.logger.Sugar().Error(err.(swagger.GenericSwaggerError).Error())
				c.logger.Sugar().Error(err.Error())

				//fmt.Println(e)
				//panic(err)

				// success, retry, remove, restart
				// 0 - success
				// 1 - retry
				// 2 - remove
				// 3 - restart

				if response.StatusCode < 300 {
					return 0
				}

				if response.StatusCode > 300 {
					c.logger.Sugar().Error(*response)
				}

				if response.StatusCode == 400 {
					c.logger.Sugar().Error(e.Message, e.Name)

					if e.Message.Valid {
						if strings.Contains(e.Message.Value, "Account has insufficient Available Balance") {
							return 2
						} else if strings.Contains(e.Message.Value, "Account is suspended") {
							return 2
						} else if strings.Contains(e.Message.Value, "Account has no") {
							return 2
						} else if strings.Contains(e.Message.Value, "Invalid account") {
							return 2
						} else if strings.Contains(e.Message.Value, "Invalid amend: orderQty, leavesQty, price, stopPx unchanged") {
							time.Sleep(time.Millisecond * 500)
						}
					}

				} else if response.StatusCode == 401 {
					//fmt.Printf("Sub Account removed: %v\n")
					return 2
				} else if response.StatusCode == 403 {
					return 2
				} else if response.StatusCode == 404 {
					return 0
				} else if response.StatusCode == 429 {
					c.logger.Sugar().Error("\n\n\nReceived 429 too many errors")
					c.logger.Sugar().Error(e.Name, e.Message)
					a, _ := strconv.Atoi(response.Header["X-Ratelimit-Reset"][0])
					reset := int64(a) - time.Now().Unix()
					c.logger.Sugar().Error("Time to reset: %v\n", reset)
					c.logger.Sugar().Error("Slept for %v seconds.\n", reset)
					time.Sleep(time.Second * time.Duration(reset))
					return 1
				} else if response.StatusCode == 503 {
					time.Sleep(time.Millisecond * 500)
					return 1
				}
			}
		}
	}
	return 0
}
