package subClient

import (
	"github.com/adi1382/Bitmex-Mirror-Bot/hostClient"
	"github.com/adi1382/Bitmex-Mirror-Bot/swagger"
	"github.com/adi1382/Bitmex-Mirror-Bot/websocket"
	"go.uber.org/atomic"
	"go.uber.org/zap"
	"sync"
	"time"
)

type SubClient struct {
	active                  atomic.Bool
	calibrateBool           atomic.Bool
	isConnectedToSocket     atomic.Bool
	isAuthenticatedToSocket atomic.Bool
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
	restartRequired         *atomic.Bool
	logger                  *zap.Logger
	wg                      *sync.WaitGroup
}

func (c *SubClient) Initialize() {

	c.logger.Info(
		"Initializing sub client",
		zap.String("websocketTopic", c.WebsocketTopic),
		zap.String("apiKey", c.ApiKey))

	c.hostUpdatesFetcher = make(chan []byte, 100)
	c.active.Store(true)
	c.chReadFromWSClient = make(chan []byte, 100)
	go c.marginUpdate()
	go c.dataHandler()
	go c.OrderHandler()

	go func() {
		for {
			if c.restartRequired.Load() {
				c.active.Store(false)
				return
			} else if !c.RunningStatus() {
				return
			}
			<-time.After(time.Nanosecond)
		}
	}()

	c.logger.Info(
		"sub client initialized",
		zap.String("websocketTopic", c.WebsocketTopic),
		zap.String("apiKey", c.ApiKey))
}

func (c *SubClient) Start() {
	c.connectToSocket()
	c.authorize()
	c.subscribeTopics("order", "position", "margin")
}

//////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////
func (c *SubClient) CloseConnection(reason string) {
	var message []interface{}
	message = append(message, 2, c.ApiKey, c.WebsocketTopic)
	c.chWriteToWSClient <- message

	c.logger.Info("closing connection for sub client",
		zap.String("apiKey", c.ApiKey),
		zap.String("websocketTopic", c.WebsocketTopic),
		zap.String("reason", reason))
	c.chReadFromWSClient <- []byte("quit")
	c.isConnectedToSocket.Store(false)
	c.isAuthenticatedToSocket.Store(false)
	c.active.Store(false)
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
		} else if !c.RunningStatus() {
			c.logger.Debug("partials receiving canceled",
				zap.String("apiKey", c.ApiKey),
				zap.String("websocketTopic", c.WebsocketTopic))
			return
		}
	}
	c.logger.Debug("partials received",
		zap.String("apiKey", c.ApiKey),
		zap.String("websocketTopic", c.WebsocketTopic))
}

func (c *SubClient) marginUpdate() {
	c.wg.Add(1)
	defer c.wg.Done()

	c.marginUpdated.Store(false)
	c.WaitForPartial()
	//fmt.Println("Margin Update Started for subClient ", c.ApiKey)
	for {
		if !c.RunningStatus() {
			return
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
				return
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
		} else if !c.RunningStatus() {
			return 0
		}
		<-time.After(time.Nanosecond)
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

func (c *SubClient) Push(message *[]byte) {
	c.chReadFromWSClient <- *message
}

func (c *SubClient) HostUpdatePush(message *[]byte) {
	c.hostUpdatesFetcher <- *message
}

func (c *SubClient) RunningStatus() bool {
	return c.active.Load()
}

func (c *SubClient) removeCurrentClient() {

	c.logger.Info("Removed subClient",
		zap.String("apiKey", c.ApiKey),
		zap.String("websocketTopic", c.WebsocketTopic))

}
