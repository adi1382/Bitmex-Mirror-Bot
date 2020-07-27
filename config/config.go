package config

import (
	"encoding/json"
	"fmt"
	"github.com/adi1382/Bitmex-Mirror-Bot/tools"
	"os"
	"time"
)

type Settings struct {
	Testnet            bool
	RatioUpdateRate    int
	CalibrationRate    int
	LimitFilledTimeout int
}

type HostAccount struct {
	ApiKey    string
	ApiSecret string
}

type SubAccount struct {
	Enabled           bool
	CopyLeverage      bool
	BalanceProportion bool
	FixedProportion   float32
	ApiKey            string
	Secret            string
}

type Config struct {
	Settings    Settings
	HostAccount HostAccount
	SubAccounts map[string]SubAccount
}

func LoadConfig(path string) Config {
	config := LoadMasterConfig(path)
	checkConfig(&config)
	return config
}

func checkConfig(config *Config) {
	apiKeys := make([]string, 0, 5)
	hostAPI := config.HostAccount.ApiKey

	if hostAPI == "" {
		fmt.Println("No API key for Host account")
		fmt.Println("Closing Bot in 5 seconds....")
		time.Sleep(time.Second * 5)
		os.Exit(-1)
	}

	for _, v := range config.SubAccounts {
		if v.Enabled {

			if v.ApiKey == "" {
				fmt.Println("No API key for enabled account")
				fmt.Println("Closing Bot in 5 seconds....")
				time.Sleep(time.Second * 5)
				os.Exit(-1)
			}

			for _, aK := range apiKeys {

				if hostAPI == v.ApiKey {
					fmt.Println("Host account duplicate in sub account.")
					fmt.Println("Closing Bot in 5 seconds....")
					time.Sleep(time.Second * 5)
					os.Exit(-1)
				}

				if aK == v.ApiKey {
					fmt.Println("Duplicate API keys detected")
					fmt.Println("Closing Bot in 5 seconds....")
					time.Sleep(time.Second * 5)
					os.Exit(-1)
				}
			}
			apiKeys = append(apiKeys, v.ApiKey)
		}
	}
}

func LoadMasterConfig(path string) Config {
	file, err := os.Open(path)
	tools.CheckErr(err)
	defer file.Close()
	decoder := json.NewDecoder(file)
	config := Config{}
	err = decoder.Decode(&config)
	tools.CheckErr(err)

	return config
}
