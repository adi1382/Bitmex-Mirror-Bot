package subClient

import (
	"fmt"
	"github.com/adi1382/Bitmex-Mirror-Bot/bitmex"
	"github.com/adi1382/Bitmex-Mirror-Bot/hostClient"
	"github.com/adi1382/Bitmex-Mirror-Bot/swagger"
	"github.com/adi1382/Bitmex-Mirror-Bot/websocket"
	"go.uber.org/atomic"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

var (
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
)

func init() {
	_, err := os.Stat("logs")

	if os.IsNotExist(err) {
		errDir := os.MkdirAll("logs", 0750)
		if errDir != nil {
			ErrorLogger.Fatal(err)
		}
	}

	file, err := os.OpenFile("logs/logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		ErrorLogger.Fatal(err)
	}

	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func NewSubClient(
	apiKey, apiSecret string,
	test, balanceProportion bool,
	fixedRatio float64,
	ch chan<- interface{},
	marginUpdateTime, calibrationTime, limitFilledTimeout int64,
	RestartCounter *atomic.Uint32,
	hostClient *hostClient.HostClient) *SubClient {

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
	c.initiateRest()

	InfoLogger.Println("New SubClient Initialized | ", c.WebsocketTopic, " | ", apiKey)

	// balanceProportion bool, fixedRatio float64,
	//hostClient *SubClient, calibrationTime int64, LimitFilledTimeout int64

	return &c
}

type SubClient struct {
	active             atomic.Bool
	marginUpdateTime   int64
	BalanceProportion  bool
	FixedRatio         float64
	test               bool
	marginUpdated      atomic.Bool
	partials           atomic.Uint32
	marginBalance      atomic.Float64
	LimitFilledTimeout int64
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
	hostClient         *hostClient.HostClient
	calibrationTime    int64
	hostUpdatesFetcher chan []byte
	restartCounter     *atomic.Uint32
}

func (c *SubClient) Initialize() {

	InfoLogger.Println("New SubClient Initializing | ", c.WebsocketTopic, " | ", c.ApiKey)

	c.hostUpdatesFetcher = make(chan []byte, 100)
	c.active.Store(true)
	c.chReadFromWSClient = make(chan []byte, 100)
	c.StartConnection()
	c.Authorize()
	go c.marginUpdate()
	go c.dataHandler()
	go c.OrderHandler()

	InfoLogger.Println("New SubClient Initialized | ", c.WebsocketTopic, " | ", c.ApiKey)
}

//////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////
func (c *SubClient) CloseConnection() {

	InfoLogger.Println("Close connection request initiated for subClient", c.ApiKey)
	c.active.Store(false)
	var message []interface{}
	message = append(message, 2, c.ApiKey, c.WebsocketTopic)
	c.chWriteToWSClient <- message
	c.chReadFromWSClient <- []byte("quit")
	c.removeCurrentClient()

}

func (c *SubClient) DropConnection() {
	if c.RunningStatus() {
		InfoLogger.Println("Drop connection request initiated for subClient ", c.ApiKey)
		c.active.Store(false)
		c.chReadFromWSClient <- []byte("quit")
	}
}

//////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////

func (c *SubClient) WaitForPartial() {
	InfoLogger.Println("Waiting For partials")
	for {
		if c.partials.Load() >= 3 {
			break
		}
	}
	InfoLogger.Println("Partials received")
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
		InfoLogger.Println("Margin Balance is: ", marginBalance)

		InfoLogger.Println("Updating margin on ", c.ApiKey)
		//c.marginBalance = c.currentMargin[0].MarginBalance.Value
		c.marginBalance.Store(marginBalance)
		c.marginUpdated.Store(true)
		InfoLogger.Println("Margin updated on ", c.ApiKey)
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
	InfoLogger.Println("Fetching balance of ", c.ApiKey)
	for {
		if c.marginUpdated.Load() {
			return c.marginBalance.Load()
		} else {
			time.Sleep(time.Nanosecond)
		}
	}
}

func (c *SubClient) initiateRest() {
	InfoLogger.Println("Initiating Rest object on subClient ", c.ApiKey)
	if c.test {
		c.Rest = swagger.NewAPIClient(swagger.NewTestnetConfiguration())
	} else {
		c.Rest = swagger.NewAPIClient(swagger.NewConfiguration())
	}
	c.Rest.InitializeAuth(c.ApiKey, c.apiSecret)
	InfoLogger.Println("Initiated Rest object on subClient ", c.ApiKey)
}

func (c *SubClient) StartConnection() {
	InfoLogger.Println("Initiating connection of subClient ", c.ApiKey, " to sockets.")
	var message []interface{}
	message = append(message, 1, c.ApiKey, c.WebsocketTopic)
	c.chWriteToWSClient <- message
	InfoLogger.Println("Initiated connection of subClient ", c.ApiKey, " to sockets.")
}

func (c *SubClient) Authorize() {
	var message []interface{}
	InfoLogger.Println("Authenticating websocket connection of subClient ", c.ApiKey)
	message = append(message, 0, c.ApiKey, c.WebsocketTopic, websocket.GetAuthMessage(c.ApiKey, c.apiSecret))
	c.chWriteToWSClient <- message
}

func (c *SubClient) SubscribeTopics(tables ...string) {
	var message []interface{}
	command := websocket.Message{Op: "subscribe"}

	for _, v := range tables {
		command.AddArgument(v)
	}

	InfoLogger.Println("Subscribing tables ", tables, " on subClient ", c.ApiKey)

	message = append(message, 0, c.ApiKey, c.WebsocketTopic, command)
	c.chWriteToWSClient <- message
}

func (c *SubClient) UnsubscribeTopics(tables ...string) {

	InfoLogger.Println("Unsubscribing tables ", tables, " on subClient ", c.ApiKey)

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

		if strings.Contains(string(message), "Access Token expired for subscription") {
			c.restartCounter.Add(1)
			ErrorLogger.Println(string(message))
			//fmt.Println(string(message))
		}

		if strings.Contains(string(message), "Invalid API Key") {
			fmt.Println("API key ", c.ApiKey, " is invalid.")
			ErrorLogger.Println(string(message))
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

		if strings.Contains(string(message), "This key is disabled") {
			ErrorLogger.Println(string(message))
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

	InfoLogger.Println("Removing subClient ", c.ApiKey, " from all clients")

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
