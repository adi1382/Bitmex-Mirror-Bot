package subClient

import (
	"github.com/adi1382/Bitmex-Mirror-Bot/swagger"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (c *SubClient) SwaggerError(err error, response *http.Response) int {

	if err != nil {

		//fmt.Println(err)
		ErrorLogger.Println("Error on subClient", c.ApiKey)

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

				//fmt.Println(e)
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
