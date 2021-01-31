package configuration

import (
	"encoding/json"
	"fmt"
	"github.com/adi1382/Bitmex-Mirror-Bot/swagger"
	"github.com/adi1382/Bitmex-Mirror-Bot/tools"
	"go.uber.org/atomic"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func ReadConfig(logger *zap.Logger, botStatus *tools.RunningStatus, restartRequired *atomic.Bool) Config {

	configLock.Lock()
	defer configLock.Unlock()

	var jsonFile *os.File
	var err error

	for {
		jsonFile, err = os.Open(ConfigPath)

		if err == nil {
			break
		} else {
			if os.IsNotExist(err) {
				generateDummyConfig()
				continue
			} else {
				fmt.Println("Could not open config file.")
				tools.EnterToExit("Could not open config file.")
			}
		}
	}

	config := Config{}

	byteValue, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		restartRequired.Store(true)
		botStatus.IsRunning.Store(false)
		botStatus.Message.Store("Unable to read the contents of the configuration file.")
		logger.Error("Unable to read the contents of the configuration file.")
		fmt.Println(err)
		tools.EnterToExit("Unable to read the contents of the configuration file.")
	}

	err = json.Unmarshal(byteValue, &config)

	if err != nil {
		restartRequired.Store(true)
		botStatus.IsRunning.Store(false)
		botStatus.Message.Store("Invalid configuration file.")
		logger.Error("Invalid configuration file.")
	}

	return config
}

func ValidateConfig(config *Config, logger *zap.Logger) (bool, string) {
	var errString string

	if config.Settings.RatioUpdateRate == 0 {
		errString = "ratioUpdateRate cannot be zero."
		return false, errString
	}

	if config.Settings.CalibrationRate == 0 {
		errString = "calibrationRate cannot be zero."
		return false, errString
	}

	if config.Settings.LimitFilledTimeout == 0 {
		errString = "limitFilledTimeout cannot be zero."
		return false, errString
	}

	if config.HostAccount.ApiKey == "" {
		errString = "Host API key cannot be empty."
		return false, errString
	}

	if config.HostAccount.Secret == "" {
		errString = "Host Secret cannot be empty."
		return false, errString
	}

	subKeys := make([]string, 0, 5)

	for key := range config.SubAccounts {
		subKeys = append(subKeys, key)
	}

	for i := range subKeys {
		if config.SubAccounts[subKeys[i]].Enabled {
			if config.SubAccounts[subKeys[i]].ApiKey == "" {
				errString = "API key empty for enabled subClient"
				return false, errString
			}

			if config.SubAccounts[subKeys[i]].Secret == "" {
				errString = "Secret empty for enabled subClient"
				return false, errString
			}

			if config.SubAccounts[subKeys[i]].FixedProportion == 0 {
				errString = "Fixed Proportion cannot be zero for enabled subClient"
				return false, errString
			}
		}
	}

	rest := new(swagger.APIClient)

	if config.Settings.Testnet {
		rest = swagger.NewAPIClient(swagger.NewTestnetConfiguration())
	} else {
		rest = swagger.NewAPIClient(swagger.NewConfiguration())
	}
	rest.InitializeAuth(config.HostAccount.ApiKey, config.HostAccount.Secret)
	var currency swagger.UserGetMarginOpts

L:
	for {

		_, response, err := rest.UserApi.UserGetMargin(&currency)
		switch SwaggerError(err, response, logger, config.HostAccount.ApiKey, "") {
		case 0:
			return true, "OK"
		case 1:
			continue L
		case 2:
			errString = "API key Invalid/Disabled on host"
			return false, errString
			//c.CloseConnection()
			//return -404
			//break function
		case 3:
			fmt.Println("Restart the bot")
			return false, "Unknown error"
		}

	}

	//return true, "OK"
}

func SwaggerError(err error, response *http.Response, logger *zap.Logger, apiKey, websocketTopic string) int {

	defer logger.Sync()

	if err != nil {

		//fmt.Println(err)
		logger.Error("Error on hostClient",
			zap.String("apiKey", apiKey),
			zap.String("websocketTopic", websocketTopic))

		if strings.Contains(err.Error(), "401") || strings.Contains(err.Error(), "403") {
			return 2
		}

		k, ok := err.(swagger.GenericSwaggerError)
		if ok {
			k, ok := k.Model().(swagger.ModelError)

			if ok {
				e := k.Error_
				logger.Sugar().Error(e.Message.Value, "///", e.Name.Value)
				logger.Sugar().Error(string(err.(swagger.GenericSwaggerError).Body()))
				logger.Sugar().Error(err.(swagger.GenericSwaggerError).Error())
				logger.Sugar().Error(err.Error())

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

				if response.StatusCode > 300 && response.StatusCode < 400 {
					logger.Sugar().Error(*response)
					logger.Error("NEW ERROR!!!!!",
						zap.Int("statusCode", response.StatusCode),
						zap.String("name", e.Name.Value),
						zap.String("message", e.Message.Value))
					return 3
				}

				if response.StatusCode == 400 {
					logger.Sugar().Error(e.Message, e.Name)

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
							return 0
						}
					}

				} else if response.StatusCode == 401 {
					return 2
				} else if response.StatusCode == 403 {
					return 2
				} else if response.StatusCode == 404 {
					return 0
				} else if response.StatusCode == 429 {
					logger.Sugar().Error("\n\n\nReceived 429 too many errors")
					logger.Sugar().Error(e.Name, e.Message)
					a, _ := strconv.Atoi(response.Header["X-Ratelimit-Reset"][0])
					reset := int64(a) - time.Now().Unix()
					logger.Sugar().Error("Time to reset: %v\n", reset)
					logger.Sugar().Error("Slept for %v seconds.\n", reset)
					time.Sleep(time.Second * time.Duration(reset))
					time.Sleep(time.Millisecond * 500)
					return 1
				} else if response.StatusCode == 503 {
					time.Sleep(time.Millisecond * 500)
					return 1
				} else {
					logger.Sugar().Error(*response)
					logger.Error("NEW ERROR!!!!!",
						zap.Int("statusCode", response.StatusCode),
						zap.String("name", e.Name.Value),
						zap.String("message", e.Message.Value))
					return 3
				}
			}
		}
	}
	return 0
}

func OnConfigChange(functionCall func()) {
	var currentModifiedTime time.Time
	var previousModifiedTime time.Time

	previousModifiedTime = getConfigModifiedTime()
	currentModifiedTime = previousModifiedTime

	go func() {
		for {

			time.Sleep(time.Second)
			currentModifiedTime = getConfigModifiedTime()

			if currentModifiedTime != previousModifiedTime {
				functionCall()
				previousModifiedTime = currentModifiedTime
			}

		}
	}()
}
