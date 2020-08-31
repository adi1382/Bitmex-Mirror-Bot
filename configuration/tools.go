package configuration

import (
	"encoding/json"
	"fmt"
	"github.com/adi1382/Bitmex-Mirror-Bot/tools"
	"go.uber.org/atomic"
	"go.uber.org/zap"
	"io/ioutil"
	"os"
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

func ValidateConfig(config *Config) (bool, string) {
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

	return true, "OK"
}

func OnConfigChange(functionCall func()) {
	var currentModifiedTime time.Time
	var previousModifiedTime time.Time

	previousModifiedTime = getConfigModifiedTime()
	currentModifiedTime = previousModifiedTime

	go func() {
		for {

			time.Sleep(time.Nanosecond)
			currentModifiedTime = getConfigModifiedTime()

			if currentModifiedTime != previousModifiedTime {
				functionCall()
				previousModifiedTime = currentModifiedTime
			}

		}
	}()
}
