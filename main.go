package main

import (
	"fmt"
	rice "github.com/GeertJohan/go.rice"
	"github.com/adi1382/Bitmex-Mirror-Bot/configuration"
	"github.com/adi1382/Bitmex-Mirror-Bot/logging"
	"github.com/adi1382/Bitmex-Mirror-Bot/server"
	"github.com/adi1382/Bitmex-Mirror-Bot/tools"
	"github.com/adi1382/Bitmex-Mirror-Bot/trader"
	"github.com/google/uuid"
	"go.uber.org/atomic"
	"go.uber.org/zap"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"
)

var port = 5000

var (
	sessionID            zap.Field
	botStatus            *tools.RunningStatus
	restartRequired      *atomic.Bool
	logger               *zap.Logger
	socketIncomingLogger *zap.Logger
	socketOutgoingLogger *zap.Logger
)

func init() {

	//fmt.Println(time.Now().Add(7 * 24 * time.Hour).Unix())
	if !tools.CheckLicense() {
		fmt.Println("License Validation Failed")
		tools.EnterToExit("Contact support@dappertrader.com to renew license.")
	}

	const expireTime = 1605543249
	timeLeft := (expireTime - time.Now().Unix()) / 3600
	fmt.Printf("\nTime left for license expiration %d hours\n", timeLeft)

	//fmt.Println(time.Now().Add(14 * 24 * time.Hour).Unix())
	//fmt.Println(time.Now().Add(time.Minute*1).Unix())
	//fmt.Println(time.Now().Add(time.Hour*24).Unix())
	//fmt.Println(((expireTime - time.Now().Unix()) / 3600) / 24)

	if time.Now().Unix() > expireTime {
		fmt.Println("License Expired!")
		time.Sleep(time.Second * 10)
		os.Exit(-1)
	}

	go func() {
		for {
			if time.Now().Unix() > expireTime {
				fmt.Println("License Expired!")
				time.Sleep(time.Second * 10)
				os.Exit(-1)
			}
			time.Sleep(5 * time.Second)
		}
	}()

	_, err := os.Stat("logs")

	if os.IsNotExist(err) {
		err = os.MkdirAll("logs", 0750)
		if err != nil {
			fmt.Println("Unable to create log folder.")
			tools.EnterToExit("Unable to create log folder.")
		}
	} else if err != nil {
		fmt.Println("Unable to perform logging operations due to the following error(s)")
		fmt.Println(err)
		tools.EnterToExit("Unable to perform logging operations due to the following error(s)")
	}

	botStatus = tools.NewBotStatus()

	restartRequired = atomic.NewBool(false)
	sessionID = zap.String("sessionID", uuid.New().String())

	logger = logging.NewLogger("Mirror", "debug", sessionID)
	socketIncomingLogger = logging.NewLogger("SocketIncoming", "debug", sessionID)
	socketOutgoingLogger = logging.NewLogger("SocketOutgoing", "debug", sessionID)
	resourceLogger := logging.NewLogger("ResourceLogger", "debug", sessionID)

	go tools.NewMonitor(60, resourceLogger)

	restartRequired.Store(false)

	configuration.OnConfigChange(
		func() {
			restartRequired.Store(true)
			botStatus.IsRunning.Store(true)
			botStatus.Message.Store("OK")
			logger.Info("Configuration File Updated")
		})

	server.SetServerLogger(logger, botStatus, restartRequired)
}

func main() {
	go trader.BotHandler(logger, socketIncomingLogger, socketOutgoingLogger, botStatus, restartRequired)
	router := http.NewServeMux()
	router.HandleFunc("/", server.ConfigHandler)
	router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(rice.MustFindBox("templates/static").HTTPBox())))
	router.Handle("/logs/", http.StripPrefix("/logs/", http.FileServer(http.Dir("logs"))))
	router.Handle("/config/", http.StripPrefix("/config/", http.FileServer(http.Dir("config"))))

	for {
		l, err := net.Listen("tcp", ":"+strconv.Itoa(port))

		if err == nil {
			_ = l.Close()
			break
		} else {
			port++
		}
	}

	fmt.Printf("\nDT Bitmex Mirror is running on http://127.0.0.1:%d/\n", port)

	err := http.ListenAndServe(":"+strconv.Itoa(port), router)
	if err != nil {
		logger.Error("Listen And Serve Error", zap.Error(err))
		fmt.Println(err.Error())
	}
}
