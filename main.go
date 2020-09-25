package main

import (
	"fmt"
	rice "github.com/GeertJohan/go.rice"
	"github.com/adi1382/Bitmex-Mirror-Bot/configuration"
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

	logger, _ = tools.NewLogger("Mirror", "debug", sessionID)
	socketIncomingLogger, _ = tools.NewLogger("SocketIncoming", "debug", sessionID)
	socketOutgoingLogger, _ = tools.NewLogger("SocketOutgoing", "debug", sessionID)
	resourceLogger, _ := tools.NewLogger("ResourceLogger", "debug", sessionID)

	go tools.NewMonitor(1, resourceLogger)

	fmt.Println("started")

	restartRequired.Store(false)

	configuration.OnConfigChange(func() {
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
