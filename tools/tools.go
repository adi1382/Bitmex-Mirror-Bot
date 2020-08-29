package tools

import (
	"bufio"
	"fmt"
	"github.com/adi1382/Bitmex-Mirror-Bot/swagger"
	"github.com/sparrc/go-ping"
	"github.com/spf13/viper"
	"go.uber.org/atomic"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strconv"
	"strings"
	"time"
)

type RunningStatus struct {
	IsRunning *atomic.Bool
	Message   *atomic.String
}

func NewBotStatus() *RunningStatus {
	return &RunningStatus{
		IsRunning: atomic.NewBool(true),
		Message:   atomic.NewString("OK"),
	}
}

func EnterToExit() {
	fmt.Print("\n\nPress 'Enter' to exit")
	_, _ = bufio.NewReader(os.Stdin).ReadBytes('\n')
	os.Exit(0)
}

func ReadConfig(
	logger *zap.Logger,
	config *viper.Viper,
	botStatus *RunningStatus,
	restartRequired *atomic.Bool) {

	defer logger.Sync()

	err := config.ReadInConfig() // Find and read the config file

	if err != nil {
		fmt.Println("The recent changes that were made to the config file have made it inaccessible.")
		fmt.Println("Kindly reconfigure the configuration and restart the program.")
		logger.Error("Unable to Read config file", zap.Error(err))

		botStatus.IsRunning.Store(false)
		botStatus.Message.Store("invalid configuration")
		//tools.EnterToExit()
	} else {
		isConfigValid, str := CheckConfig(config)

		if !isConfigValid {
			fmt.Println("Invalid Configuration")
			fmt.Println(str)
			EnterToExit()
		}

		if !botStatus.IsRunning.Load() {
			botStatus.IsRunning.Store(true)
			botStatus.Message.Store("OK")
		}

		restartRequired.Store(true)
	}
}

func NewLogger(fileName, level string, sessionID zap.Field) (*zap.Logger, error) {

	logLevel := zap.DebugLevel

	if strings.EqualFold(level, "INFO") {
		logLevel = zap.InfoLevel
	} else if strings.EqualFold(level, "WARN") {
		logLevel = zap.WarnLevel
	} else if strings.EqualFold(level, "ERROR") {
		logLevel = zap.ErrorLevel
	}

	config := zap.Config{
		Level:       zap.NewAtomicLevelAt(logLevel),
		Development: true,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:      "json",
		EncoderConfig: zap.NewProductionEncoderConfig(),
		//OutputPaths:      []string{"stderr"},
		//ErrorOutputPaths: []string{"stderr"},
		OutputPaths:      []string{"./logs/" + fileName + ".log"},
		ErrorOutputPaths: []string{"./logs/" + fileName + ".log"},
	}

	config.EncoderConfig.TimeKey = "TimeUTC"

	config.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.UTC().Format("2006-01-02T15:04:05Z0700"))
		// 2019-08-13T04:39:11Z
	}

	logger, err := config.Build()

	if err != nil {
		panic(err)
	}

	logger = logger.With(sessionID)
	return logger, err
}

func CheckConfig(config *viper.Viper) (bool, string) {
	//fmt.Println("Check config", viper.Sub("Settings").GetDuration("CalibrationRate") * time.Second)
	//if !strings.EqualFold(viper.Sub("Settings").GetString("Testnet"), "true") &&
	//	!strings.EqualFold(viper.Sub("Settings").GetString("Testnet"), "false") {
	//	return false
	//}

	//fmt.Println(viper.Sub("Settings"))
	var errStr string

	if config.Sub("Settings") != nil {
		settings := config.Sub("Settings")

		_, err := strconv.ParseBool(settings.GetString("Testnet"))
		if err != nil {
			errStr = "Invalid configuration (Testnet)"
			return false, errStr
		}

		_, err = strconv.ParseFloat(settings.GetString("RatioUpdateRate"), 64)
		if err != nil {
			errStr = "Invalid configuration (RatioUpdateRate)"
			return false, errStr
		}

		_, err = strconv.ParseFloat(settings.GetString("CalibrationRate"), 64)
		if err != nil {
			errStr = "Invalid configuration (CalibrationRate)"
			return false, errStr
		}

		_, err = strconv.ParseFloat(settings.GetString("LimitFilledTimeout"), 64)
		if err != nil {
			errStr = "Invalid configuration (LimitFilledTimeout)"
			return false, errStr
		}
	} else {
		errStr = "Settings does not exists"
		return false, errStr
	}

	if config.Sub("HostAccount") != nil {
		hostAccount := config.Sub("HostAccount")

		if hostAccount.GetString("ApiKey") == "" {
			errStr = "HostAccount ApiKey does not exists"
			return false, errStr
		}

		if hostAccount.GetString("Secret") == "" {
			errStr = "HostAccount Secret does not exists"
			return false, errStr
		}

	} else {
		errStr = "HostAccount does not exists"
		return false, errStr
	}

	allKeys := make([]string, 0, 5)

	if config.Sub("SubAccounts") != nil {
		subAccounts := config.Sub("SubAccounts").AllSettings()
		for key := range subAccounts {
			allKeys = append(allKeys, key)
		}
	} else {
		errStr = "Invalid Configuration (SubAccounts)"
		return false, errStr
	}

	for i := range allKeys {
		subAccount := config.Sub("SubAccounts").Sub(allKeys[i])

		_, ok := subAccount.AllSettings()[strings.ToLower("AccountName")]

		if !ok {
			errStr = "AccountName option does not exists for " + allKeys[i]
			return false, errStr
		}

		isEnabled, err := strconv.ParseBool(subAccount.GetString("Enabled"))
		if err != nil {
			errStr = "Invalid configuration for 'Enabled' on " + allKeys[i]
			return false, errStr
		}

		if isEnabled {

			if subAccount.GetString("ApiKey") == "" {
				errStr = "Invalid API key for " + allKeys[i]
				return false, errStr
			}

			if subAccount.GetString("Secret") == "" {
				errStr = "Invalid Secret key for " + allKeys[i]
				return false, errStr
			}

			_, err := strconv.ParseBool(subAccount.GetString("CopyLeverage"))
			if err != nil {
				errStr = "Invalid configuration for 'CopyLeverage' on " + allKeys[i]
				return false, errStr
			}

			_, err = strconv.ParseBool(subAccount.GetString("BalanceProportion"))
			if err != nil {
				errStr = "Invalid configuration for 'BalanceProportion' on " + allKeys[i]
				return false, errStr
			}

			_, err = strconv.ParseFloat(subAccount.GetString("FixedProportion"), 64)
			if err != nil {
				errStr = "Invalid configuration for 'FixedProportion' on " + allKeys[i]
				return false, errStr
			}
		}
	}

	return true, "OK"
}

func CheckConnection(baseUrl *string) {
	for {
		pinger, err := ping.NewPinger(*baseUrl)
		if err != nil {
			fmt.Println("Unable to connect to ", *baseUrl)
			time.Sleep(time.Second * 5)
			continue
		}
		pinger.SetPrivileged(true)
		pinger.Count = 1
		pinger.Timeout = time.Second * 5
		//pinger.OnFinish()
		pinger.Run()                 // blocks until finished
		stats := pinger.Statistics() // get send/receive/rtt stats
		if stats.PacketsRecv < stats.PacketsSent {
			fmt.Println("Unable to connect to ", *baseUrl)
			time.Sleep(time.Second * 5)
			continue
		}
		break
	}
}

func CheckErr(err error) {
	if err != nil {
		fmt.Println(err)
		e := err.(swagger.GenericSwaggerError).Model().(swagger.ModelError).Error_
		fmt.Println(e.Name, e.Message)
		panic(err)
	}
}
