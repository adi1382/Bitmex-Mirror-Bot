package main

import (
	"encoding/json"
	"fmt"
	"github.com/adi1382/Bitmex-Mirror-Bot/Mirror"
	"github.com/adi1382/Bitmex-Mirror-Bot/configuration"
	"github.com/adi1382/Bitmex-Mirror-Bot/hostClient"
	"github.com/adi1382/Bitmex-Mirror-Bot/subClient"
	"github.com/adi1382/Bitmex-Mirror-Bot/tools"
	"github.com/adi1382/Bitmex-Mirror-Bot/websocket"
	"github.com/google/uuid"
	"go.uber.org/atomic"
	"go.uber.org/zap"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

var (
	sessionID       zap.Field
	botStatus       *tools.RunningStatus
	restartRequired *atomic.Bool
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
}

func main() {

	logger, _ := tools.NewLogger("Mirror", "debug", sessionID)
	socketIncomingLogger, _ := tools.NewLogger("SocketIncoming", "debug", sessionID)
	socketOutgoingLogger, _ := tools.NewLogger("SocketOutgoing", "debug", sessionID)
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

	go func() {
		for {

			if botStatus.IsRunning.Load() {
				restartRequired.Store(false)
				fmt.Println("Trader Initiated..")
				for {
				}
				trader(logger, socketIncomingLogger, socketOutgoingLogger)
				fmt.Printf("\n\n\n")
			}
			time.Sleep(time.Nanosecond)
		}
	}()

	//con := configuration.ReadConfig(logger, botStatus, restartRequired)
	//con.SubAccounts["account1"].AccountName

	tmpl := template.Must(template.ParseFiles("templates/index.gohtml"))
	configHandler := func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodPost {
			fmt.Println("POST")
			err := r.ParseForm()

			var newConfig configuration.Config
			newConfig.SubAccounts = make(map[string]configuration.SubAccount, 0)

			exchangeType, err := strconv.Atoi(r.Form["ExchangeType"][0])
			ratioUpdateRate, err := strconv.ParseFloat(r.Form["RatioUpdateRate"][0], 64)
			calibrationRate, err := strconv.ParseFloat(r.Form["CalibrationRate"][0], 64)
			limitFilledTimeout, err := strconv.ParseFloat(r.Form["LimitFilledTimeout"][0], 64)
			hostApiKey := r.Form["HostApiKey"][0]
			hostSecret := r.Form["HostSecret"][0]

			if exchangeType == 0 {
				newConfig.Settings.Testnet = true
			} else {
				newConfig.Settings.Testnet = false
			}

			newConfig.Settings.RatioUpdateRate = ratioUpdateRate
			newConfig.Settings.CalibrationRate = calibrationRate
			newConfig.Settings.LimitFilledTimeout = limitFilledTimeout
			newConfig.HostAccount.ApiKey = hostApiKey
			newConfig.HostAccount.Secret = hostSecret

			{
				i := 1

				for {
					var newSubAccount configuration.SubAccount

					sub := "sub" + strconv.Itoa(i) + "_"

					_, ok := r.Form[sub+"Status"]

					if !ok {
						break
					}

					configUpdateErrorCheck := func(err error) bool {
						if err != nil {
							logger.Error("Invalid Config Update", zap.Error(err))
							return true
						}
						return false
					}

					isEnabled, err := strconv.Atoi(r.Form[sub+"Status"][0])
					if configUpdateErrorCheck(err) {
						return
					}

					copyLeverage, err := strconv.Atoi(r.Form[sub+"CopyLeverage"][0])
					if configUpdateErrorCheck(err) {
						return
					}

					balanceProportional, err := strconv.Atoi(r.Form[sub+"BalanceProportional"][0])
					if configUpdateErrorCheck(err) {
						return
					}

					fixedProportion, err := strconv.ParseFloat(r.Form[sub+"FixedProportion"][0], 64)
					if configUpdateErrorCheck(err) {
						return
					}

					apiKey := r.Form[sub+"ApiKey"][0]
					secret := r.Form[sub+"Secret"][0]
					accountName := r.Form[sub+"AccountName"][0]

					if isEnabled == 1 {
						newSubAccount.Enabled = true
					} else {
						newSubAccount.Enabled = false
					}

					if copyLeverage == 1 {
						newSubAccount.CopyLeverage = true
					} else {
						newSubAccount.CopyLeverage = false
					}

					if balanceProportional == 1 {
						newSubAccount.BalanceProportion = true
					} else {
						newSubAccount.BalanceProportion = false
					}

					newSubAccount.FixedProportion = fixedProportion
					newSubAccount.ApiKey = apiKey
					newSubAccount.Secret = secret
					newSubAccount.AccountName = accountName
					newSubAccount.AccountNumber = i

					newConfig.SubAccounts["account"+strconv.Itoa(i)] = newSubAccount
					i++
				}
			}

			marshaledJSON, err := json.MarshalIndent(newConfig, "", "	")

			if err != nil {
				logger.Error("Unable to Marhsal with indent", zap.Error(err))
			}

			configuration.WriteConfig(&marshaledJSON, logger)
		}

		err := tmpl.Execute(w, configuration.ReadConfig(logger, botStatus, restartRequired))
		if err != nil {
			logger.Error("GET config Error", zap.Error(err))
		}
		return
	}

	http.HandleFunc("/", configHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("templates/static"))))
	http.Handle("/logs/", http.StripPrefix("/logs/", http.FileServer(http.Dir("logs"))))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		logger.Error("Listen And Serve Error", zap.Error(err))
	}

}

func trader(logger, socketIncomingLogger, socketOutgoingLogger *zap.Logger) {

	defer func() {
		fmt.Println("Trader closed..")
	}()

	config := configuration.ReadConfig(logger, botStatus, restartRequired)
	isOK, message := configuration.ValidateConfig(&config)

	if !isOK {
		fmt.Println("Invalid Configuration File")
		fmt.Println(message)
		botStatus.Message.Store(message)
		botStatus.IsRunning.Store(false)
	}

	if !botStatus.IsRunning.Load() {
		return
	}

	var wg sync.WaitGroup

	mirror := Mirror.NewMirror(restartRequired, logger, &wg)

	logger.Info("logging started")

	// Connect to WS
	conn, err := websocket.Connect(config.Settings.Testnet, logger)

	if err != nil {
		botStatus.IsRunning.Store(false)
		botStatus.Message.Store("Unable to establish websocket connection")
		fmt.Println(err)
		tools.EnterToExit("Unable to establish websocket connection")
	}

	// Listen write WS
	chReadFromWS := make(chan []byte, 100)
	go websocket.ReadFromWSToChannel(conn, chReadFromWS, socketIncomingLogger, logger, restartRequired, &wg)

	// Write to WS
	chWriteToWS := make(chan interface{}, 100)
	go websocket.WriteFromChannelToWS(conn, chWriteToWS, socketOutgoingLogger, logger, restartRequired, &wg)
	//go viperConfig.SocketResponseDistributor(chReadFromWS, RestartCounter, &wg)

	websocket.PingPong(conn, restartRequired, logger, &wg)

	//RestartCounter.Store(0)

	//hostClient.Initialize(cfg.HostAccount.ApiKey, cfg.HostAccount.ApiSecret,
	//	"hostAccount", chWriteToWS, false,
	//	0, cfg.Settings.Testnet, cfg.Settings.RatioUpdateRate, &dummyClient, cfg.Settings.CalibrationRate,
	//	cfg.Settings.LimitFilledTimeout, &RestartCounter)
	//hostClient.subscribeTopics("order", "position", "margin")

	host := hostClient.NewHostClient(
		config.HostAccount.ApiKey,
		config.HostAccount.Secret,
		config.Settings.Testnet,
		chWriteToWS,
		config.Settings.RatioUpdateRate,
		restartRequired,
		logger,
		&wg,
		botStatus)

	mirror.SetHost(host)

	subKeys := make([]string, 0, 5)
	for key := range config.SubAccounts {
		subKeys = append(subKeys, key)
	}

	//subAccounts := viperConfig.Sub("SubAccounts")
	for i := range subKeys {

		if config.SubAccounts[subKeys[i]].Enabled {
			sub := subClient.NewSubClient(
				config.SubAccounts[subKeys[i]].ApiKey,
				config.SubAccounts[subKeys[i]].Secret,
				config.Settings.Testnet,
				config.SubAccounts[subKeys[i]].BalanceProportion,
				config.SubAccounts[subKeys[i]].FixedProportion,
				config.Settings.RatioUpdateRate,
				config.Settings.CalibrationRate,
				config.Settings.LimitFilledTimeout,
				chWriteToWS,
				restartRequired,
				host,
				logger,
				&wg)
			mirror.AddSub(sub)
		}
	}

	go mirror.SocketResponseDistributor(chReadFromWS)

	mirror.InitializeAll()

	mirror.StartMirroring()

	fmt.Println("Running...")

	go func() {
		wg.Add(1)
		defer wg.Done()
		for {
			if restartRequired.Load() {
				_ = conn.Close()
				chWriteToWS <- "quit"
				break
			}
			<-time.After(time.Nanosecond)
		}
	}()

	wg.Wait()
	logger.Info("All wait groups completed")
}
