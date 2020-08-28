package subClient

import (
	"github.com/adi1382/Bitmex-Mirror-Bot/hostClient"
	"github.com/adi1382/Bitmex-Mirror-Bot/swagger"
	"go.uber.org/atomic"
	"go.uber.org/zap"
	"sync"
)

func NewSubClient(
	apiKey, apiSecret string,
	test, balanceProportion bool,
	fixedRatio, marginUpdateTime, calibrationTime, limitFilledTimeout float64,
	ch chan<- interface{},
	restartRequired *atomic.Bool,
	hostClient *hostClient.HostClient,
	logger *zap.Logger,
	wg *sync.WaitGroup) *SubClient {

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
	c.wg = wg
	c.Rest.InitializeAuth(c.ApiKey, c.apiSecret)
	c.WebsocketTopic = ""

	c.restartRequired = restartRequired
	c.active.Store(true)
	c.isConnectedToSocket.Store(false)
	c.isAuthenticatedToSocket.Store(false)
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
