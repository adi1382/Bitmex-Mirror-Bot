package main

import (
	"fmt"
	"github.com/adi1382/Bitmex-Mirror-Bot/configuration"
	"github.com/adi1382/Bitmex-Mirror-Bot/server"
	"github.com/adi1382/Bitmex-Mirror-Bot/tools"
	"github.com/adi1382/Bitmex-Mirror-Bot/trader"
	"github.com/google/uuid"
	"go.uber.org/atomic"
	"go.uber.org/zap"
	"net/http"
	"os"
)

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

	//tmpl := template.Must(template.ParseFiles("templates/index.gohtml"))

	http.HandleFunc("/", server.ConfigHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("templates/static"))))
	http.Handle("/logs/", http.StripPrefix("/logs/", http.FileServer(http.Dir("logs"))))
	http.Handle("/config/", http.StripPrefix("/config/", http.FileServer(http.Dir("config"))))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		logger.Error("Listen And Serve Error", zap.Error(err))
	}

}
