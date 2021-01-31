package subClient

import (
	"crypto/rand"
	"fmt"
	"github.com/adi1382/Bitmex-Mirror-Bot/bitmex"
	"github.com/adi1382/Bitmex-Mirror-Bot/swagger"
	"github.com/adi1382/Bitmex-Mirror-Bot/websocket"
	"go.uber.org/zap"
	"log"
	"strings"
	"time"
)

func (c *SubClient) dataHandler() {
	c.wg.Add(1)
	defer c.wg.Done()
	//fmt.Println("Data Handler started for subClient ", c.ApiKey)
	defer func() {
		c.logger.Info("Data Handler Closed for subClient",
			zap.String("apiKey", c.ApiKey),
			zap.String("websocketTopic", c.WebsocketTopic))
		//fmt.Println("Data Handler Closed for subClient ", c.ApiKey)
	}()
	for {

		if !c.RunningStatus() {
			return
		}

		message := <-c.chReadFromWSClient

		c.logger.Debug("New Message in Data Handler for subClient",
			zap.String("apiKey", c.ApiKey),
			zap.String("websocketTopic", c.WebsocketTopic))

		strMessage := string(message)
		//fmt.Println(strMessage)
		if strMessage == "quit" {
			c.active.Store(false)
			return
		}

		//if strings.Contains(strMessage, "Access Token expired for subscription") {
		//	c.logger.Info("RestartRequiredToTrue")
		//	c.restartRequired.Store(true)
		//
		//	c.logger.Error("Expiration Error",
		//		zap.String("errorMessage", strMessage),
		//		zap.String("apiKey", c.ApiKey),
		//		zap.String("websocketTopic", c.WebsocketTopic))
		//	//fmt.Println(string(message))
		//}
		//
		//if strings.Contains(strMessage, "Invalid API Key") {
		//	fmt.Println("API key ", c.ApiKey, " is invalid.")
		//
		//	c.logger.Error("api key invalid for subClient",
		//		zap.String("errMessage", strMessage),
		//		zap.String("apiKey", c.ApiKey),
		//		zap.String("websocketTopic", c.WebsocketTopic))
		//
		//	c.CloseConnection()
		//}
		//
		//if strings.Contains(strMessage, "This key is disabled") {
		//	c.logger.Error("apiKey is disabled on subClient",
		//		zap.String("errorMessage", strMessage),
		//		zap.String("apiKey", c.ApiKey),
		//		zap.String("websocketTopic", c.WebsocketTopic))
		//
		//	fmt.Println("API key ", c.ApiKey, " is disabled.")
		//
		//	c.CloseConnection()
		//}

		c.extractCoreMessage(&strMessage)
		shouldClose, whyClose := c.socketError(&strMessage)

		if shouldClose {
			if strings.Contains(whyClose, "Signature not valid") ||
				strings.Contains(whyClose, "This key is disabled") ||
				strings.Contains(whyClose, "Invalid API Key") {
				c.CloseConnection(whyClose)
			} else {
				c.restartRequired.Store(true)
				c.logger.Error("Socket Error", zap.String("error", whyClose))
				return
			}
		}

		if !c.RunningStatus() {
			return
		}

		if !c.isConnectedToSocket.Load() {
			c.checkIfConnected(strMessage)
		} else if !c.isAuthenticatedToSocket.Load() {
			c.checkIfAuthenticated(strMessage)
		}

		if !strings.Contains(strMessage, "table") {
			continue
		}

		response, table := bitmex.DecodeMessage([]byte(strMessage), c.logger, c.restartRequired)
		if c.restartRequired.Load() {
			return
		}

		c.logger.Debug("Updating table on subClient",
			zap.String("table", table),
			zap.String("apiKey", c.ApiKey),
			zap.String("websocketTopic", c.WebsocketTopic))

		// Potential Race Condition when fetching
		if table == "order" {

			orderResponse := response.(bitmex.OrderResponse)

			c.ordersLock.Lock()
			if orderResponse.Action == "partial" {
				c.partials.Add(1)
				c.activeOrders.OrderPartial(&orderResponse.Data)
			} else if orderResponse.Action == "insert" {
				c.activeOrders.OrderInsert(&orderResponse.Data)
			} else if orderResponse.Action == "update" {
				c.activeOrders.OrderUpdate(&orderResponse.Data)
			} else if orderResponse.Action == "delete" {
				c.activeOrders.OrderDelete(&orderResponse.Data)
			}
			c.ordersLock.Unlock()

		} else if table == "position" {
			positionResponse := response.(bitmex.PositionResponse)

			c.positionsLock.Lock()
			if positionResponse.Action == "partial" {
				c.partials.Add(1)
				c.activePositions.PositionPartial(&positionResponse.Data)
			} else if positionResponse.Action == "insert" {
				c.activePositions.PositionInsert(&positionResponse.Data)
			} else if positionResponse.Action == "update" {
				c.activePositions.PositionUpdate(&positionResponse.Data)
			} else if positionResponse.Action == "delete" {
				c.activePositions.PositionDelete(&positionResponse.Data)
			}
			c.positionsLock.Unlock()

		} else if table == "margin" {
			marginResponse := response.(bitmex.MarginResponse)

			c.marginLock.Lock()
			if marginResponse.Action == "partial" {
				c.partials.Add(1)
				c.currentMargin.MarginPartial(&marginResponse.Data)
			} else if marginResponse.Action == "insert" {
				c.currentMargin.MarginInsert(&marginResponse.Data)
			} else if marginResponse.Action == "update" {
				c.currentMargin.MarginUpdate(&marginResponse.Data)
			} else if marginResponse.Action == "delete" {
				c.currentMargin.MarginDelete(&marginResponse.Data)
			}
			c.marginLock.Unlock()
		}
	}
}

func (c *SubClient) OrderHandler() {
	var lenOfOrdersCounter int

	c.wg.Add(1)
	defer c.wg.Done()

	c.WaitForPartial()
	c.hostClient.WaitForPartial()

	defer func() {
		c.logger.Info("Order handler closed for subClient",
			zap.String("apiKey", c.ApiKey),
			zap.String("websocketTopic", c.WebsocketTopic))
		//fmt.Println("Order Handler Closed for subClient ", c.ApiKey)
	}()

	//fmt.Println("Order Handler Started for subClient ", c.ApiKey)

	c.calibrate()

	c.calibrateBool.Store(true)
	//calibrateBool := true
	calibrateBoolReset := time.Now().Add(time.Second * time.Duration(c.calibrationTime))
	//calibrateTrigger := time.After(time.Second*time.Duration(c.calibrationTime))

	go func() {

		c.wg.Add(1)
		defer c.wg.Done()

		defer func() {
			c.logger.Info("Calibrator timer closed for subClient ",
				zap.String("apiKey", c.ApiKey),
				zap.String("websocketTopic", c.WebsocketTopic))
			//fmt.Println("Calibrator timer closed for subClient ", c.ApiKey)
		}()

		//fmt.Println("Calibrator timer Started for subClient ", c.ApiKey)

		for {
			time.Sleep(time.Millisecond * 100)

			if !c.RunningStatus() {
				return
			}

			if !c.calibrateBool.Load() && calibrateBoolReset.Unix() < time.Now().Unix() {
				calibrateBoolReset = time.Now().Add(time.Second * time.Duration(c.calibrationTime))
				c.calibrateBool.Store(true)
			}

			//if calibrateBool == false {
			//	time.Sleep(time.Second*time.Duration(c.calibrationTime))
			//	calibrateBool = true
			//}
		}
	}()

	for {
		time.Sleep(time.Millisecond)

		if !c.RunningStatus() {
			return
		}

		select {

		case message := <-c.hostUpdatesFetcher:
			c.mirroring(&message, &calibrateBoolReset)
			continue
		default:
			if c.calibrateBool.Load() {
				c.calibrate()

				if len(c.hostClient.ActiveOrders()) != len(c.getActiveOrders()) {
					lenOfOrdersCounter++
					if lenOfOrdersCounter > 3 {
						c.orderCancelAll()
						c.restartRequired.Store(true)
						return
					}
				} else {
					lenOfOrdersCounter = 0
				}

				c.calibrateBool.Store(false)
				continue
			}
			continue
		}
	}
}

func (c *SubClient) mirroring(message *[]byte, calibrateBoolReset *time.Time) {

	c.logger.Debug("Starting Mirror for subClient",
		zap.String("apiKey", c.ApiKey),
		zap.String("websocketTopic", c.WebsocketTopic))

	strResponse := string(*message)
	prefix := fmt.Sprintf(`[0,"%s","%s",`, c.hostClient.ApiKey, "hostAccount")
	suffix := "]"
	strResponse = strings.TrimPrefix(strResponse, prefix)
	strResponse = strings.TrimSuffix(strResponse, suffix)

	var ratio float64

	c.logger.Debug("Calculating ratio on subClient",
		zap.String("apiKey", c.ApiKey),
		zap.String("websocketTopic", c.WebsocketTopic))

	if c.BalanceProportion {
		ratio = c.GetMarginBalance() / c.hostClient.GetMarginBalance()
	} else {
		ratio = c.FixedRatio
	}

	response, table := bitmex.DecodeMessage([]byte(strResponse), c.logger, c.restartRequired)
	if c.restartRequired.Load() {
		return
	}

	c.logger.Debug("Updating table from mirror on subClient",
		zap.String("table", table),
		zap.String("apiKey", c.ApiKey),
		zap.String("websocketTopic", c.WebsocketTopic))

	// Potential Race Condition when fetching
	if table == "order" {
		orderResponse, ok := response.(bitmex.OrderResponse)

		if !ok {
			c.logger.Debug("Updating table from mirror on subClient",
				zap.String("table", table),
				zap.String("apiKey", c.ApiKey),
				zap.String("websocketTopic", c.WebsocketTopic))

			c.logger.Sugar().Error("Invalid Interface Conversion",
				orderResponse,
				zap.String("apiKey", c.ApiKey),
				zap.String("websocketTopic", c.WebsocketTopic))

			c.logger.Info("RestartRequiredToTrue")
			c.restartRequired.Store(true)
			return
		}

		if orderResponse.Action == "insert" {
			c.logger.Debug("New Order Inserted in subClient",
				zap.String("apiKey", c.ApiKey),
				zap.String("websocketTopic", c.WebsocketTopic))

			//fmt.Println("host Margin Balance: ", hostClient.GetMarginBalance())
			//fmt.Println("Sub Margin Balance: ", subClient.GetMarginBalance())

			orders := make([]map[string]interface{}, 0, 5)

			for h := range orderResponse.Data {

				ord := make(map[string]interface{})
				ord["clOrdID"] = random() + orderResponse.Data[h].OrderID.Value[8:]
				if orderResponse.Data[h].Symbol.Valid {
					ord["symbol"] = orderResponse.Data[h].Symbol.Value
				}

				if orderResponse.Data[h].OrdType.Value == "Market" {
					ord["ordType"] = "Market"
					if orderResponse.Data[h].Side.Valid {
						ord["side"] = orderResponse.Data[h].Side.Value
					}

					if orderResponse.Data[h].LeavesQty.Valid {
						ord["orderQty"] = int(orderResponse.Data[h].OrderQty.Value * ratio)
					}
				} else {
					if orderResponse.Data[h].Side.Valid {
						ord["side"] = orderResponse.Data[h].Side.Value
					}

					if orderResponse.Data[h].LeavesQty.Valid {
						ord["orderQty"] = int(orderResponse.Data[h].OrderQty.Value * ratio)
					}
					if orderResponse.Data[h].Price.Valid {
						ord["price"] = orderResponse.Data[h].Price.Value
					}
					if orderResponse.Data[h].DisplayQty.Valid {
						ord["displayQty"] = int(orderResponse.Data[h].DisplayQty.Value * ratio)
					}
					if orderResponse.Data[h].StopPx.Valid {
						ord["stopPx"] = orderResponse.Data[h].StopPx.Value
					}
					if orderResponse.Data[h].PegOffsetValue.Valid {
						ord["pegOffsetValue"] = orderResponse.Data[h].PegOffsetValue.Value
					}
					if orderResponse.Data[h].PegPriceType.Valid {
						ord["pegPriceType"] = orderResponse.Data[h].PegPriceType.Value
					}
					if orderResponse.Data[h].OrdType.Valid {
						ord["ordType"] = orderResponse.Data[h].OrdType.Value
					}
					if orderResponse.Data[h].TimeInForce.Valid {
						ord["timeInForce"] = orderResponse.Data[h].TimeInForce.Value
					}
					if orderResponse.Data[h].ExecInst.Valid {
						ord["execInst"] = orderResponse.Data[h].ExecInst.Value
					}
					if orderResponse.Data[h].Text.Valid {
						ord["text"] = orderResponse.Data[h].Text.Value
					}
				}
				orders = append(orders, ord)
			}

			var symbols []string

			for _, order := range orders {
				if !isIn(order["symbol"].(string), symbols) {
					symbols = append(symbols, order["symbol"].(string))
				}
			}

			for _, symbol := range symbols {
				placeNewOrders := make([]interface{}, 0, 5)
				for _, order := range orders {
					if order["symbol"].(string) == symbol {
						placeNewOrders = append(placeNewOrders, order)
					}
				}

				c.logger.Debug("New Order Placed on subClient",
					zap.String("symbol", symbol),
					zap.String("apiKey", c.ApiKey),
					zap.String("websocketTopic", c.WebsocketTopic))
				c.orderNewBulk(&placeNewOrders)

			}

		} else if orderResponse.Action == "update" {

			c.logger.Debug("New Order Update received for subClient",
				zap.String("apiKey", c.ApiKey),
				zap.String("websocketTopic", c.WebsocketTopic))

			amendOrders := make([]map[string]interface{}, 0, 5)

			activeOrders := c.getActiveOrders()

			var toCancel []string
			for h := range orderResponse.Data {

				if orderResponse.Data[h].OrdStatus.Valid {
					if orderResponse.Data[h].OrdStatus.Value == "Filled" || orderResponse.Data[h].OrdStatus.Value == "PartiallyFilled" {
						c.logger.Debug("OrderStatus Update subClient",
							zap.String("orderStatus", orderResponse.Data[h].OrdStatus.Value),
							zap.String("apiKey", c.ApiKey),
							zap.String("websocketTopic", c.WebsocketTopic))

						c.calibrateBool.Store(false)
						*calibrateBoolReset = time.Now().Add(time.Second * time.Duration(c.LimitFilledTimeout))
					}
				}

				if !orderResponse.Data[h].OrdStatus.Valid {
					if orderResponse.Data[h].Price.Valid || orderResponse.Data[h].StopPx.Valid || orderResponse.Data[h].LeavesQty.Valid || orderResponse.Data[h].PegOffsetValue.Valid {

						c.logger.Debug("Amended Order Detected for SubClient",
							zap.String("apiKey", c.ApiKey),
							zap.String("websocketTopic", c.WebsocketTopic))

						subOrders := getSubOrder(orderResponse.Data[h].OrderID.Value, activeOrders)

						if len(subOrders) == 0 {
							continue
						}

						subOrder := subOrders[0]

						if len(subOrders) > 1 {
							extraOrders := subOrders[1:]
							for i := range extraOrders {
								toCancel = append(toCancel, extraOrders[i].OrderID.Value)
							}
						}

						if (orderResponse.Data[h].OrdType.Value == "StopLimit" || orderResponse.Data[h].OrdType.Value == "LimitIfTouched") &&
							orderResponse.Data[h].Triggered.Value != "" {

							if orderResponse.Data[h].Price.Value != subOrder.Price.Value ||
								int(orderResponse.Data[h].LeavesQty.Value*ratio) != int(subOrder.LeavesQty.Value) {

								amend := make(map[string]interface{})

								amend["symbol"] = subOrder.Symbol.Value
								amend["orderID"] = subOrder.OrderID.Value
								if orderResponse.Data[h].Price.Valid {
									amend["price"] = orderResponse.Data[h].Price.Value
								}
								if orderResponse.Data[h].LeavesQty.Valid {
									amend["orderQty"] = int(orderResponse.Data[h].LeavesQty.Value * ratio)
								}

								amendOrders = append(amendOrders, amend)
							}

						} else {
							if orderResponse.Data[h].Price.Value != subOrder.Price.Value ||
								int(orderResponse.Data[h].LeavesQty.Value*ratio) != int(subOrder.LeavesQty.Value) ||
								orderResponse.Data[h].StopPx.Value != subOrder.StopPx.Value ||
								orderResponse.Data[h].PegOffsetValue.Value != subOrder.PegOffsetValue.Value {

								amend := make(map[string]interface{})

								amend["symbol"] = subOrder.Symbol.Value
								amend["orderID"] = subOrder.OrderID.Value
								if orderResponse.Data[h].Price.Valid {
									amend["price"] = orderResponse.Data[h].Price.Value
								}
								if orderResponse.Data[h].LeavesQty.Valid {
									amend["orderQty"] = int(orderResponse.Data[h].LeavesQty.Value * ratio)
								}
								if orderResponse.Data[h].StopPx.Valid {
									amend["stopPx"] = orderResponse.Data[h].StopPx.Value
								}
								if orderResponse.Data[h].PegOffsetValue.Valid {
									amend["pegOffsetValue"] = orderResponse.Data[h].PegOffsetValue.Value
								}
								amendOrders = append(amendOrders, amend)
							}
						}

					}
				} else if orderResponse.Data[h].OrdStatus.Valid {

					subOrders := getSubOrder(orderResponse.Data[h].OrderID.Value, activeOrders)

					if len(subOrders) == 0 {
						continue
					}
					subOrder := subOrders[0]

					if orderResponse.Data[h].OrdStatus.Value == "Canceled" {
						toCancel = append(toCancel, subOrder.OrderID.Value)
					}

					if len(subOrders) > 1 {
						extraOrders := subOrders[1:]
						for i := range extraOrders {
							toCancel = append(toCancel, extraOrders[i].OrderID.Value)
						}
					}
				}
			}

			symbols := make([]string, 0, 5)
			for _, order := range amendOrders {
				if !isIn(order["symbol"].(string), symbols) {
					symbols = append(symbols, order["symbol"].(string))
				}
			}
			for _, symbol := range symbols {
				amendOldOrders := make([]interface{}, 0, 5)
				for _, order := range amendOrders {
					if order["symbol"].(string) == symbol {
						_, ok := order["symbol"]
						if ok {
							delete(order, "symbol")
							amendOldOrders = append(amendOldOrders, order)
						}
						order["symbol"] = symbol
					}
				}

				c.logger.Debug("Order Amended on subClient",
					zap.String("symbol", symbol),
					zap.String("apiKey", c.ApiKey),
					zap.String("websocketTopic", c.WebsocketTopic))

				c.orderAmendBulk(&amendOldOrders)
			}

			c.logger.Debug("Order Cancel request on subClient",
				zap.Int("noOfOrders", len(toCancel)),
				zap.String("apiKey", c.ApiKey),
				zap.String("websocketTopic", c.WebsocketTopic))
			c.orderCancelBulk(&toCancel)
		}

	} else if table == "position" {

		c.logger.Debug("Position update from mirror on subClient",
			zap.String("apiKey", c.ApiKey),
			zap.String("websocketTopic", c.WebsocketTopic))
		positionResponse := response.(bitmex.PositionResponse)

		if positionResponse.Action == "update" {

			for i := range positionResponse.Data {

				if positionResponse.Data[i].CrossMargin.Valid {
					if positionResponse.Data[i].CrossMargin.Value {
						c.UpdateLeverage(positionResponse.Data[i].Symbol.Value, 0)

					} else if positionResponse.Data[i].Leverage.Valid {
						c.UpdateLeverage(positionResponse.Data[i].Symbol.Value, positionResponse.Data[i].Leverage.Value)

					} else {
						activePositions := c.hostClient.ActivePositions()
						for i := range activePositions {
							if activePositions[i].Symbol.Value == positionResponse.Data[i].Symbol.Value {
								c.UpdateLeverage(positionResponse.Data[i].Symbol.Value, activePositions[i].Leverage.Value)
							}
						}
					}
				} else if positionResponse.Data[i].Leverage.Valid {
					c.UpdateLeverage(positionResponse.Data[i].Symbol.Value, positionResponse.Data[i].Leverage.Value)

				}
			}
		}

	}
}

func getSubOrder(id string, orders websocket.OrderSlice) []swagger.Order {
	returnValue := make([]swagger.Order, 0, 5)
	for i := range orders {
		if orders[i].ClOrdID.Value[8:] == id[8:] {
			returnValue = append(returnValue, orders[i])
		}
	}

	return returnValue
}

func random() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid[0:8]
}

func isIn(key string, xi []string) bool {
	for _, v := range xi {
		if key == v {
			return true
		}
	}
	return false
}
