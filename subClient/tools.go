package subClient

import (
	"encoding/json"
	"fmt"
	"github.com/adi1382/Bitmex-Mirror-Bot/swagger"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (c *SubClient) checkErr(err error) {
	if err != nil {
		c.logger.Error("New Error Detected",
			zap.Error(err))
		c.restartRequired.Store(true)
	}
}

func (c *SubClient) socketError(strMessage *string) (bool, string) {
	if !strings.Contains(*strMessage, "error") {
		return false, ""
	} else {
		var obj map[string]interface{}
		var returnString string
		err := json.Unmarshal([]byte(*strMessage), &obj)

		if err != nil {
			c.logger.Error("Cannot unmarshal socket error message",
				zap.Error(err),
				zap.String("message", *strMessage))
			c.restartRequired.Store(true)
			return true, "cannot unmarshal socket error message"
		}

		returnInterface, ok := obj["error"]

		if ok {
			returnString, ok = returnInterface.(string)

			if !ok {
				c.restartRequired.Store(true)
				c.logger.Error("error message is not of type string",
					zap.String("message", *strMessage))
				return true, "error message is not of type string"
			}

		} else {
			c.logger.Error("NEW ERROR!!",
				zap.String("socketErrMessage", *strMessage))

			c.restartRequired.Store(true)
			return true, "NEW SOCKET ERROR"
		}

		if ok {
			return true, returnString
		} else {
			c.restartRequired.Store(true)
			return true, ""
		}
	}
}

func (c *SubClient) extractCoreMessage(strMessage *string) {
	prefix := fmt.Sprintf(`[0,"%s","%s",`, c.ApiKey, c.WebsocketTopic)
	suffix := "]"
	*strMessage = strings.TrimPrefix(*strMessage, prefix)
	*strMessage = strings.TrimSuffix(*strMessage, suffix)
}

func (c *SubClient) checkIfConnected(strMessage string) {
	if strings.Contains(strMessage, "Welcome to the BitMEX Realtime API.") {
		c.isConnectedToSocket.Store(true)

		c.logger.Info("Successfully connected to websockets",
			zap.String("apiKey", c.ApiKey),
			zap.String("websocketTopic", c.WebsocketTopic))
	} else {
		c.CloseConnection("First message on subClient is something different")
		c.restartRequired.Store(true)
		c.logger.Error("NEW ERROR!! First message on subClient is something different",
			zap.String("error", strMessage))
	}
}

func (c *SubClient) checkIfAuthenticated(strMessage string) {
	if strings.Contains(strMessage, `"success":true`) && strings.Contains(strMessage, "authKeyExpires") {
		c.isAuthenticatedToSocket.Store(true)

		c.logger.Info("Successfully authenticated to websockets",
			zap.String("apiKey", c.ApiKey),
			zap.String("websocketTopic", c.WebsocketTopic))
	} else if strings.Contains(strMessage, "Signature not valid.") {
		c.CloseConnection("Signature not valid.")
		c.logger.Error("API Secret is Invalid.",
			zap.String("apiKey", c.ApiKey),
			zap.String("websocketTopic", c.WebsocketTopic))
	}
}

func (c *SubClient) SwaggerError(err error, response *http.Response) int {

	if err != nil {

		//fmt.Println(err)
		c.logger.Error("Error on subClient",
			zap.String("apiKey", c.ApiKey),
			zap.String("websocketTopic", c.WebsocketTopic))

		if strings.Contains(err.Error(), "401") || strings.Contains(err.Error(), "403") {
			return 2
		}

		k, ok := err.(swagger.GenericSwaggerError)
		if ok {
			k, ok := k.Model().(swagger.ModelError)

			if ok {
				e := k.Error_
				c.logger.Sugar().Error(e.Message.Value, "///", e.Name.Value)
				c.logger.Sugar().Error(string(err.(swagger.GenericSwaggerError).Body()))
				c.logger.Sugar().Error(err.(swagger.GenericSwaggerError).Error())
				c.logger.Sugar().Error(err.Error())

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
					c.logger.Sugar().Error(*response)
				}

				if response.StatusCode == 400 {
					c.logger.Sugar().Error(e.Message, e.Name)

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
					c.logger.Sugar().Error("\n\n\nReceived 429 too many errors")
					c.logger.Sugar().Error(e.Name, e.Message)
					a, _ := strconv.Atoi(response.Header["X-Ratelimit-Reset"][0])
					reset := int64(a) - time.Now().Unix()
					c.logger.Sugar().Error("Time to reset: %v\n", reset)
					c.logger.Sugar().Error("Slept for %v seconds.\n", reset)
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
