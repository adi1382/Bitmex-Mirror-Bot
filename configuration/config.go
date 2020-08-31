package configuration

import (
	"sync"
)

const ConfigPath = "config/config.json"

var configLock sync.Mutex

type Settings struct {
	Testnet            bool    `json:"testnet"`
	RatioUpdateRate    float64 `json:"ratioUpdateRate"`
	CalibrationRate    float64 `json:"calibrationRate"`
	LimitFilledTimeout float64 `json:"limitFilledTimeout"`
}

type HostAccount struct {
	ApiKey string `json:"apiKey"`
	Secret string `json:"secret"`
}

type SubAccount struct {
	Enabled           bool    `json:"enabled"`
	CopyLeverage      bool    `json:"copyLeverage"`
	BalanceProportion bool    `json:"balanceProportion"`
	FixedProportion   float64 `json:"fixedProportion"`
	ApiKey            string  `json:"apiKey"`
	Secret            string  `json:"secret"`
	AccountName       string  `json:"accountName"`
	AccountNumber     int     `json:"accountNumber"`
}

type Config struct {
	Settings    Settings              `json:"settings"`
	HostAccount HostAccount           `json:"hostAccount"`
	SubAccounts map[string]SubAccount `json:"subAccounts"`
}
