package tools

import (
	"bufio"
	"fmt"
	"github.com/adi1382/Bitmex-Mirror-Bot/swagger"
	"github.com/sparrc/go-ping"
	"go.uber.org/atomic"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
	"time"
)

type RunningStatus struct {
	IsRunning *atomic.Bool
	Message   *atomic.String
}

func NewBotStatus() *RunningStatus {
	return &RunningStatus{
		IsRunning: atomic.NewBool(false),
		Message:   atomic.NewString("First Start Pending"),
	}
}

func EnterToExit(errMessage string) {
	fmt.Println(errMessage)
	fmt.Print("\n\nPress 'Enter' to exit")
	_, _ = bufio.NewReader(os.Stdin).ReadBytes('\n')
	os.Exit(0)
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
