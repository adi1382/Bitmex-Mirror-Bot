package server

import (
	"encoding/json"
	"github.com/adi1382/Bitmex-Mirror-Bot/configuration"
	"github.com/adi1382/Bitmex-Mirror-Bot/tools"
	"go.uber.org/atomic"
	"go.uber.org/zap"
	"html/template"
	"net/http"
	"strconv"
)

var (
	logger          *zap.Logger
	botStatus       *tools.RunningStatus
	restartRequired *atomic.Bool
)

func SetServerLogger(loggerMain *zap.Logger, botStatusMain *tools.RunningStatus, restartRequiredMain *atomic.Bool) {
	logger = loggerMain
	botStatus = botStatusMain
	restartRequired = restartRequiredMain
}

func ConfigHandler(w http.ResponseWriter, r *http.Request) {

	type configHandler struct {
		BotStatus bool
		Config    configuration.Config
	}

	tmpl := template.Must(template.ParseFiles("templates/index.gohtml"))

	if r.Method == http.MethodPost {
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
			logger.Error("Unable to Marshal with indent", zap.Error(err))
		}

		configuration.WriteConfig(&marshaledJSON, logger)
		botStatus.IsRunning.Store(true)
		botStatus.Message.Store("OK")
	}

	err := tmpl.Execute(w, configHandler{
		Config:    configuration.ReadConfig(logger, botStatus, restartRequired),
		BotStatus: botStatus.IsRunning.Load(),
	})
	if err != nil {
		logger.Error("GET config Error", zap.Error(err))
	}
	return
}
