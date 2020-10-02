package tools

import (
	"bufio"
	"fmt"
	"github.com/adi1382/Bitmex-Mirror-Bot/constants"
	"github.com/adi1382/Bitmex-Mirror-Bot/swagger"
	"github.com/adi1382/Bitmex-Mirror-Bot/wmic"
	"github.com/sparrc/go-ping"
	"go.uber.org/atomic"
	"os"
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

func CheckLicense() bool {
	if constants.HashedKey == wmic.GetHashedKey() {
		return true
	}
	return false
}

func CheckErr(err error) {
	if err != nil {
		fmt.Println(err)
		e := err.(swagger.GenericSwaggerError).Model().(swagger.ModelError).Error_
		fmt.Println(e.Name, e.Message)
		panic(err)
	}
}
