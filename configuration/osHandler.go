package configuration

import (
	"encoding/json"
	"fmt"
	"github.com/adi1382/Bitmex-Mirror-Bot/tools"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

func getConfigModifiedTime() time.Time {
	configLock.Lock()
	defer configLock.Unlock()

	for {
		fileStat, err := os.Stat(path)

		if err == nil {
			return fileStat.ModTime()
		} else {
			if os.IsNotExist(err) {
				generateDummyConfig()
				continue
			} else {
				fmt.Println("Could not get stats of config file.")
				configLock.Unlock()
				tools.EnterToExit("Could not get stats of config file.")
			}
		}
	}
}

func generateDummyConfig() {
	dummyConfig := Config{
		Settings: Settings{
			Testnet:            false,
			RatioUpdateRate:    3600,
			CalibrationRate:    10,
			LimitFilledTimeout: 10,
		},
		HostAccount: HostAccount{
			ApiKey: "",
			Secret: "",
		},
		SubAccounts: map[string]SubAccount{},
	}

	for i := 1; i <= 5; i++ {
		dummyConfig.SubAccounts["account"+strconv.Itoa(i)] = SubAccount{
			Enabled:           false,
			CopyLeverage:      true,
			BalanceProportion: false,
			FixedProportion:   1,
			ApiKey:            "",
			Secret:            "",
			AccountName:       "",
		}
	}

	marshaledJSON, err := json.MarshalIndent(dummyConfig, "", "	")

	if err != nil {
		fmt.Println("JSON Marshaling error")
		tools.EnterToExit("JSON Marshaling error")
	}

	pathParts := strings.Split(path, "/")

	_, err = os.Stat(pathParts[0])

	if os.IsNotExist(err) {
		err = os.MkdirAll(pathParts[0], 0750)
		if err != nil {
			fmt.Println("Unable to create config folder.")
			tools.EnterToExit("Unable to create config folder.")
		}
	}

	_, err = os.Stat(path)

	if os.IsNotExist(err) {
		err = ioutil.WriteFile(path, marshaledJSON, 0644)
		if err != nil {
			fmt.Println("Unable to write to config file")
			tools.EnterToExit("Unable to write to config file")
		}
	}
}
