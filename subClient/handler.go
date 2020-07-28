package subClient

import (
	"crypto/rand"
	"fmt"
	"github.com/adi1382/Bitmex-Mirror-Bot/bitmex"
	"github.com/adi1382/Bitmex-Mirror-Bot/swagger"
	"github.com/adi1382/Bitmex-Mirror-Bot/websocket"
	"log"
	"strings"
	"time"
)

func (c *SubClient) OrderHandler() {

	c.WaitForPartial()
	c.hostClient.WaitForPartial()

	defer func() {
		InfoLogger.Println("Order Handler Closed for subClient ", c.ApiKey)
		//fmt.Println("Order Handler Closed for subClient ", c.ApiKey)
	}()

	//fmt.Println("Order Handler Started for subClient ", c.ApiKey)

	c.calibrate()

	calibrateBool := true
	calibrateBoolReset := time.Now().Add(time.Second * time.Duration(c.calibrationTime))

	go func() {

		defer func() {
			InfoLogger.Println("Calibrator timer closed for subClient ", c.ApiKey)
			//fmt.Println("Calibrator timer closed for subClient ", c.ApiKey)
		}()

		//fmt.Println("Calibrator timer Started for subClient ", c.ApiKey)

		for {
			time.Sleep(time.Nanosecond)

			if !c.RunningStatus() {
				break
			}

			if calibrateBool == false && calibrateBoolReset.Unix() < time.Now().Unix() {
				calibrateBoolReset = time.Now().Add(time.Second * time.Duration(c.calibrationTime))
				calibrateBool = true
			}

			//if calibrateBool == false {
			//	time.Sleep(time.Second*time.Duration(c.calibrationTime))
			//	calibrateBool = true
			//}
		}
	}()

	for {

		time.Sleep(time.Nanosecond)

		if !c.RunningStatus() {
			break
		}

		select {

		case message := <-c.hostUpdatesFetcher:
			c.mirroring(&message, &calibrateBoolReset, &calibrateBool)
			continue
		default:
			if calibrateBool {
				c.calibrate()
				calibrateBool = false
				continue
			}
			continue
		}
	}
}

func (c *SubClient) mirroring(message *[]byte, calibrateBoolReset *time.Time, calibrateBool *bool) {

	InfoLogger.Println("Starting Mirror for subClient ", c.ApiKey)

	strResponse := string(*message)
	prefix := fmt.Sprintf(`[0,"%s","%s",`, c.hostClient.ApiKey, "hostAccount")
	suffix := fmt.Sprintf("]")
	strResponse = strings.TrimPrefix(strResponse, prefix)
	strResponse = strings.TrimSuffix(strResponse, suffix)

	var ratio float64
	InfoLogger.Println("Calculating ratio on subClient ", c.ApiKey)
	if c.BalanceProportion {
		ratio = c.GetMarginBalance() / c.hostClient.GetMarginBalance()
	} else {
		ratio = c.FixedRatio
	}

	response, table := bitmex.DecodeMessage([]byte(strResponse))

	InfoLogger.Println("Manipulating ", table, " from mirror on subClient ", c.ApiKey)

	// Potential Race Condition when fetching
	if table == "order" {
		orderResponse, ok := response.(bitmex.OrderResponse)

		if !ok {
			ErrorLogger.Println("Invalid Interface Conversion of ", orderResponse)
			panic("Invalid Conversion")
		}

		if orderResponse.Action == "insert" {

			InfoLogger.Println("New Order Inserted for SubClient ", c.ApiKey)

			//fmt.Println("Host Margin Balance: ", hostClient.GetMarginBalance())
			//fmt.Println("Sub Margin Balance: ", subClient.GetMarginBalance())

			orders := make([]map[string]interface{}, 0, 5)

			for h := range orderResponse.Data {

				ord := make(map[string]interface{})
				ord["clOrdID"] = random() + orderResponse.Data[h].OrderID.Value[8:]
				if orderResponse.Data[h].Symbol.Valid {
					ord["symbol"] = orderResponse.Data[h].Symbol.Value
				}
				if orderResponse.Data[h].Side.Valid {
					ord["side"] = orderResponse.Data[h].Side.Value
				}
				if orderResponse.Data[h].LeavesQty.Valid {
					ord["orderQty"] = int(orderResponse.Data[h].LeavesQty.Value * ratio)
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

				InfoLogger.Println("New Order placed for symbol ", symbol, " on subClient ", c.ApiKey)

				c.OrderNewBulk(&placeNewOrders)

			}

		} else if orderResponse.Action == "update" {
			InfoLogger.Println("New Order update received for SubClient ", c.ApiKey)

			amendOrders := make([]map[string]interface{}, 0, 5)

			activeOrders := c.ActiveOrders()

			var toCancel []string
			for h := range orderResponse.Data {

				if orderResponse.Data[h].OrdStatus.Valid {
					if orderResponse.Data[h].OrdStatus.Value == "Filled" || orderResponse.Data[h].OrdStatus.Value == "PartiallyFilled" {
						InfoLogger.Println("Order " + orderResponse.Data[h].OrdStatus.Value)
						*calibrateBool = false
						*calibrateBoolReset = time.Now().Add(time.Second * time.Duration(c.LimitFilledTimeout))
					}
				}

				if !orderResponse.Data[h].OrdStatus.Valid {
					if orderResponse.Data[h].Price.Valid || orderResponse.Data[h].StopPx.Valid || orderResponse.Data[h].LeavesQty.Valid || orderResponse.Data[h].PegOffsetValue.Valid {

						InfoLogger.Println("Amended order detected for SubClient ", c.ApiKey)

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

				InfoLogger.Println("Order amended for symbol ", symbol, " on subClient ", c.ApiKey)

				c.OrderAmendBulk(&amendOldOrders)
			}

			InfoLogger.Println(len(toCancel), " Orders canceled for subClient ", c.ApiKey)
			c.OrderCancelBulk(&toCancel)
		}

	} else if table == "position" {

		InfoLogger.Println("Position Update from mirror for subClient ", c.ApiKey)

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
