package hostClient

import (
	"fmt"
	"github.com/adi1382/Bitmex-Mirror-Bot/bitmex"
	"github.com/adi1382/Bitmex-Mirror-Bot/swagger"
	"github.com/adi1382/Bitmex-Mirror-Bot/websocket"
	"go.uber.org/atomic"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
)

func init() {
	//_, err := os.Stat("logs")
	//
	//if os.IsNotExist(err) {
	//	errDir := os.MkdirAll("logs", 0750)
	//	if errDir != nil {
	//		ErrorLogger.Fatal(err)
	//	}
	//}

	file, err := os.OpenFile("logs/logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		ErrorLogger.Fatal(err)
	}

	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func NewHostClient(apiKey, apiSecret string, test bool, ch chan<- interface{}, marginUpdateTime int64,
	RestartCounter *atomic.Uint32) *HostClient {
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
	InfoLogger.Println("New HostClient Initialized | ", c.WebsocketTopic, " | ", apiKey)

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
}

func (c *HostClient) Initialize() {
	c.startSocketConnection()
	c.socketAuthentication()
	go c.marginUpdate()
	go c.dataHandler()

	InfoLogger.Println("New Client Initialized | ", c.WebsocketTopic, " | ", c.ApiKey)
}

func (c *HostClient) CloseConnection() {
	c.active.Store(false)
	var message []interface{}
	message = append(message, 2, c.ApiKey, c.WebsocketTopic)
	c.chWriteToWSClient <- message
	c.chReadFromWSClient <- []byte("quit")
	c.restartCounter.Add(1)
	InfoLogger.Println("Closed connection for hostClient", c.ApiKey)
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
		InfoLogger.Println("Margin Balance is: ", marginBalance)

		InfoLogger.Println("Updating margin on ", c.ApiKey)

		c.marginBalance.Store(marginBalance)

		c.marginUpdated.Store(true)

		InfoLogger.Println("Margin updated on ", c.ApiKey)

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
	InfoLogger.Println("Initiating connection of subClient ", c.ApiKey, " to sockets.")
	var message []interface{}
	message = append(message, 1, c.ApiKey, c.WebsocketTopic)
	c.chWriteToWSClient <- message
	InfoLogger.Println("Initiated connection of subClient ", c.ApiKey, " to sockets.")
}

func (c *HostClient) socketAuthentication() {
	var message []interface{}
	InfoLogger.Println("Authenticating websocket connection of subClient ", c.ApiKey)
	message = append(message, 0, c.ApiKey, c.WebsocketTopic, websocket.GetAuthMessage(c.ApiKey, c.apiSecret))
	c.chWriteToWSClient <- message
}

func (c *HostClient) SubscribeTopics(tables ...string) {
	var message []interface{}
	command := websocket.Message{Op: "subscribe"}

	for _, v := range tables {
		command.AddArgument(v)
	}

	InfoLogger.Println("Subscribing tables ", tables, " on subClient ", c.ApiKey)

	message = append(message, 0, c.ApiKey, c.WebsocketTopic, command)
	c.chWriteToWSClient <- message
}

func (c *HostClient) UnsubscribeTopics(tables ...string) {

	InfoLogger.Println("Unsubscribing tables ", tables, " on subClient ", c.ApiKey)

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
		InfoLogger.Println("Data Handler Closed for subClient ", c.ApiKey)
		//fmt.Println("Data Handler Closed for subClient ", c.ApiKey)
	}()
	for {

		if !c.RunningStatus() {
			break
		}

		message := <-c.chReadFromWSClient
		InfoLogger.Println("Received new message in Data Handler for subClient ", c.ApiKey)
		strResponse := string(message)
		if strResponse == "quit" {
			break
		}

		if strings.Contains(strResponse, "Access Token expired for subscription") {
			c.restartCounter.Add(1)
			//atomic.AddInt64(c.restartCounter, 1)
			ErrorLogger.Println(strResponse)
			//fmt.Println(string(message))
		}

		if strings.Contains(strResponse, "Invalid API Key") {
			fmt.Println("API key ", c.ApiKey, " is invalid.")
			ErrorLogger.Println(strResponse)
			ErrorLogger.Println("API key ", c.ApiKey, " is invalid.")
			if c.WebsocketTopic == "hostAccount" {
				fmt.Println("Host Account API key is Invalid. Closing the bot in 10 seconds.")
				ErrorLogger.Println("Host Account API key is Invalid. Closing the bot in 10 seconds.")
				time.Sleep(time.Second * 10)
				os.Exit(-1)
			} else {
				c.CloseConnection()
			}
		}

		if strings.Contains(strResponse, "This key is disabled") {
			ErrorLogger.Println(strResponse)
			fmt.Println("API key ", c.ApiKey, " is disabled.")
			ErrorLogger.Println("API key ", c.ApiKey, " is disabled.")
			if c.WebsocketTopic == "hostAccount" {
				fmt.Println("Host Account API key is disabled. Closing the bot in 10 seconds.")
				ErrorLogger.Println("Host Account API key is disabled. Closing the bot in 10 seconds.")
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

		response, table := bitmex.DecodeMessage([]byte(strResponse))

		InfoLogger.Println("Manipulating ", table, " table on subClient ", c.ApiKey)

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
	InfoLogger.Println("Fetching Current margin for subClient ", c.ApiKey)
	c.marginLock.Lock()
	defer c.marginLock.Unlock()
	return c.currentMargin
}

func (c *HostClient) restMargin() float64 {
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
			fmt.Println("API key Invalid/Disabled on Host")
			InfoLogger.Println("API key Invalid/Disabled on Host")
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
		ErrorLogger.Println("Error on subClient", c.ApiKey)

		if strings.Contains(err.Error(), "401") || strings.Contains(err.Error(), "403") {
			return 2
		}

		k, ok := err.(swagger.GenericSwaggerError)
		if ok {
			k, ok := k.Model().(swagger.ModelError)

			if ok {
				e := k.Error_
				ErrorLogger.Println(e.Message.Value, "///", e.Name.Value)
				ErrorLogger.Println(string(err.(swagger.GenericSwaggerError).Body()))
				ErrorLogger.Println(err.(swagger.GenericSwaggerError).Error())
				ErrorLogger.Println(err.Error())

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
					ErrorLogger.Println(*response)
				}

				if response.StatusCode == 400 {
					ErrorLogger.Println(e.Message, e.Name)

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
					ErrorLogger.Printf("\n\n\nReceived 429 too many errors")
					ErrorLogger.Println(e.Name, e.Message)
					a, _ := strconv.Atoi(response.Header["X-Ratelimit-Reset"][0])
					reset := int64(a) - time.Now().Unix()
					ErrorLogger.Printf("Time to reset: %v\n", reset)
					ErrorLogger.Printf("Slept for %v seconds.\n", reset)
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
