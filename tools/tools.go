package tools

import (
	"fmt"
	"github.com/adi1382/Bitmex-Mirror-Bot/swagger"
	"github.com/sparrc/go-ping"
	"go.uber.org/atomic"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
)

func init() {
	_, err := os.Stat("logs")

	if os.IsNotExist(err) {
		errDir := os.MkdirAll("logs", 0750)
		if errDir != nil {
			ErrorLogger.Fatal(err)
		}
	}

	file, err := os.OpenFile("logs/logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		ErrorLogger.Fatal(err)
	}

	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
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

func WebsocketErr(err error, RestartCounter *atomic.Uint32) {
	if err != nil {
		RestartCounter.Add(1)
	}
}

func SwaggerError123(err error, response *http.Response) int {

	if err != nil {

		fmt.Println(err)

		if strings.Contains(err.Error(), "401") || strings.Contains(err.Error(), "403") {
			return 2
		}

		k, ok := err.(swagger.GenericSwaggerError)
		if ok {
			k, ok := k.Model().(swagger.ModelError)

			if ok {
				e := k.Error_
				ErrorLogger.Println(e.Message.Value, "///", e.Name.Value)
				ErrorLogger.Println(string(err.(swagger.GenericSwaggerError).Body()))
				ErrorLogger.Println(err.(swagger.GenericSwaggerError).Error())
				ErrorLogger.Println(err.Error())

				fmt.Println(e)
				//panic(err)

				// success, retry, remove, restart
				// 0 - success
				// 1 - retry
				// 2 - remove
				// 3 - restart

				if response.StatusCode < 300 {
					return 0
				}

				if response.StatusCode > 300 {
					ErrorLogger.Println(*response)
				}

				if response.StatusCode == 400 {
					ErrorLogger.Println(e.Message, e.Name)

					if e.Message.Valid {
						if strings.Contains(e.Message.Value, "Account has insufficient Available Balance") {
							return 2
						} else if strings.Contains(e.Message.Value, "Account is suspended") {
							return 2
						} else if strings.Contains(e.Message.Value, "Account has no") {
							return 2
						} else if strings.Contains(e.Message.Value, "Invalid account") {
							return 2
						} else if strings.Contains(e.Message.Value, "Invalid amend: orderQty, leavesQty, price, stopPx unchanged") {
							time.Sleep(time.Millisecond * 500)
						}
					}

				} else if response.StatusCode == 401 {
					//fmt.Printf("Sub Account removed: %v\n")
					return 2
				} else if response.StatusCode == 403 {
					return 2
				} else if response.StatusCode == 404 {
					return 0
				} else if response.StatusCode == 429 {
					ErrorLogger.Printf("\n\n\nReceived 429 too many errors")
					ErrorLogger.Println(e.Name, e.Message)
					a, _ := strconv.Atoi(response.Header["X-Ratelimit-Reset"][0])
					reset := int64(a) - time.Now().Unix()
					ErrorLogger.Printf("Time to reset: %v\n", reset)
					ErrorLogger.Printf("Slept for %v seconds.\n", reset)
					time.Sleep(time.Second * time.Duration(reset))
					return 1
				} else if response.StatusCode == 503 {
					time.Sleep(time.Millisecond * 500)
					return 1
				}
			}
		}
	}
	return 0
}

//func WriteGob(filePath string, object interface{}) error {
//	file, err := os.Create(filePath)
//	if err == nil {
//		encoder := gob.NewEncoder(file)
//		encoder.Encode(object)
//	}
//	file.Close()
//	return err
//}
//
//func ReadGob(filePath string, object interface{}) error {
//	file, err := os.Open(filePath)
//	if err == nil {
//		decoder := gob.NewDecoder(file)
//		err = decoder.Decode(object)
//	}
//	file.Close()
//	return err
//}
