package subClient

import (
	"github.com/adi1382/Bitmex-Mirror-Bot/hostClient"
	"strings"
)

func (c *SubClient) calibrate() {
	for {
	L:
		for {
			select {
			case <-c.hostUpdatesFetcher:
			default:
				break L
			}
		}

		doMirror(c.hostClient, c)

		if len(c.hostUpdatesFetcher) > 0 {
			continue
		} else {
			break
		}
	}
}

//func (c *SubClient) StartMirroring(orderHandlerMutex *sync.Mutex) {
//
//	for {
//
//		if !c.RunningStatus() {
//			fmt.Println("Closed Mirror")
//			break
//		}
//
//		fmt.Println("Calibration Sequence initiated 1..........................")
//		orderHandlerMutex.Lock()
//
//		//for len(c.hostUpdatesFetcher) > 0 {
//		//	fmt.Println(len(c.hostUpdatesFetcher))
//		//	fmt.Println("emptying fetcher")
//		//	<-c.hostUpdatesFetcher
//		//}
//
//		L:
//			for {
//				fmt.Println("Length channel: ", len(c.hostUpdatesFetcher))
//				select {
//				case <-c.hostUpdatesFetcher:
//				default:
//					break L
//				}
//			}
//		fmt.Println("Calibration Sequence initiated..........................")
//		time.Sleep(time.Second*5)
//
//		doMirror(c.hostClient, c)
//
//		if len(c.hostUpdatesFetcher) != 0 {
//			fmt.Println("looped")
//			orderHandlerMutex.Unlock()
//			continue
//		}
//
//		fmt.Println("Calibration Sequence completed..........................")
//		orderHandlerMutex.Unlock()
//		time.Sleep(time.Second*time.Duration(c.calibrationTime))
//	}
//
//}

func doMirror(hostClient *hostClient.HostClient, subClient *SubClient) {
	var ratio float64
	if subClient.BalanceProportion {
		ratio = subClient.GetMarginBalance() / hostClient.GetMarginBalance()
	} else {
		ratio = subClient.FixedRatio
	}

	hostPositions := hostClient.ActivePositions()
	hostOrders := hostClient.ActiveOrders()
	subPositions := subClient.ActivePositions()
	subOrders := subClient.ActiveOrders()

	// Sequence of events
	// 1. Cancel not required Orders
	// 2. Copy Leverage
	// 3. Calibrate Position
	// 4. Calibrate Position Margin
	// 5. Place/amend required Orders

	cancelOrderIds := make([]string, 0, 2)
	ordersToPlace := make([]map[string]interface{}, 0, 5)
	ordersToAmend := make([]map[string]interface{}, 0, 5)

	for hIdx := range hostOrders {

		orderFound := false

		// Iterating through all the sub accounts
		for sIdx := range subOrders {

			// If order is found on the sub account
			if hostOrders[hIdx].OrderID.Value[8:] == subOrders[sIdx].ClOrdID.Value[8:] {
				if !orderFound {
					orderFound = true
					//check order for any required amend

					amend := make(map[string]interface{})

					if (hostOrders[hIdx].OrdType.Value == "StopLimit" || hostOrders[hIdx].OrdType.Value == "LimitIfTouched") &&
						hostOrders[hIdx].Triggered.Value != "" {

						if hostOrders[hIdx].Price.Value != subOrders[sIdx].Price.Value ||
							int(hostOrders[hIdx].LeavesQty.Value*ratio) != int(subOrders[sIdx].LeavesQty.Value) {

							amend["symbol"] = subOrders[sIdx].Symbol.Value
							amend["orderID"] = subOrders[sIdx].OrderID.Value
							if hostOrders[hIdx].Price.Valid {
								amend["price"] = hostOrders[hIdx].Price.Value
							}
							if hostOrders[hIdx].LeavesQty.Valid {
								amend["orderQty"] = int(hostOrders[hIdx].LeavesQty.Value * ratio)
							}

							ordersToAmend = append(ordersToAmend, amend)
						}

					} else {

						if hostOrders[hIdx].Price.Value != subOrders[sIdx].Price.Value ||
							int(hostOrders[hIdx].LeavesQty.Value*ratio) != int(subOrders[sIdx].LeavesQty.Value) ||
							hostOrders[hIdx].StopPx.Value != subOrders[sIdx].StopPx.Value ||
							hostOrders[hIdx].PegOffsetValue.Value != subOrders[sIdx].PegOffsetValue.Value {

							amend["symbol"] = subOrders[sIdx].Symbol.Value
							amend["orderID"] = subOrders[sIdx].OrderID.Value
							if hostOrders[hIdx].Price.Valid {
								amend["price"] = hostOrders[hIdx].Price.Value
							}
							if hostOrders[hIdx].LeavesQty.Valid {
								amend["orderQty"] = int(hostOrders[hIdx].LeavesQty.Value * ratio)
							}
							if hostOrders[hIdx].StopPx.Valid {
								amend["stopPx"] = hostOrders[hIdx].StopPx.Value
							}
							if hostOrders[hIdx].PegOffsetValue.Valid {
								amend["pegOffsetValue"] = hostOrders[hIdx].PegOffsetValue.Value
							}
							ordersToAmend = append(ordersToAmend, amend)
						}

					}

				} else {
					// remove all the duplicated sub orders
					cancelOrderIds = append(cancelOrderIds, subOrders[sIdx].OrderID.Value)
				}
			}
		}

		// if host order is not found on the sub account
		if !orderFound {
			// if not found place it

			if hostOrders[hIdx].LeavesQty.Value != 0 {

				if int(hostOrders[hIdx].LeavesQty.Value*ratio) != 0 {

					ord := make(map[string]interface{})

					if (hostOrders[hIdx].OrdType.Value == "StopLimit" || hostOrders[hIdx].OrdType.Value == "LimitIfTouched") &&
						hostOrders[hIdx].Triggered.Value != "" {

						ord["clOrdID"] = random() + hostOrders[hIdx].OrderID.Value[8:]
						if hostOrders[hIdx].Symbol.Valid {
							ord["symbol"] = hostOrders[hIdx].Symbol.Value
						}
						if hostOrders[hIdx].Side.Valid {
							ord["side"] = hostOrders[hIdx].Side.Value
						}
						if hostOrders[hIdx].LeavesQty.Valid {
							ord["orderQty"] = int(hostOrders[hIdx].LeavesQty.Value * ratio)
						}
						if hostOrders[hIdx].Price.Valid {
							ord["price"] = hostOrders[hIdx].Price.Value
						}
						if hostOrders[hIdx].DisplayQty.Valid {
							ord["displayQty"] = int(hostOrders[hIdx].DisplayQty.Value * ratio)
						}

						///////////////////

						if hostOrders[hIdx].TimeInForce.Valid {
							ord["timeInForce"] = hostOrders[hIdx].TimeInForce.Value
						}
						if hostOrders[hIdx].ExecInst.Valid {

							if strings.Contains(hostOrders[hIdx].ExecInst.Value, "Close") {
								ord["execInst"] = "ReduceOnly"
							}
						}
						if hostOrders[hIdx].Text.Valid {
							ord["text"] = hostOrders[hIdx].Text.Value
						}

						ordersToPlace = append(ordersToPlace, ord)

					} else {

						ord["clOrdID"] = random() + hostOrders[hIdx].OrderID.Value[8:]
						if hostOrders[hIdx].Symbol.Valid {
							ord["symbol"] = hostOrders[hIdx].Symbol.Value
						}
						if hostOrders[hIdx].Side.Valid {
							ord["side"] = hostOrders[hIdx].Side.Value
						}
						if hostOrders[hIdx].LeavesQty.Valid {
							ord["orderQty"] = int(hostOrders[hIdx].LeavesQty.Value * ratio)
						}
						if hostOrders[hIdx].Price.Valid {
							ord["price"] = hostOrders[hIdx].Price.Value
						}
						if hostOrders[hIdx].DisplayQty.Valid {
							ord["displayQty"] = int(hostOrders[hIdx].DisplayQty.Value * ratio)
						}
						if hostOrders[hIdx].StopPx.Valid {
							ord["stopPx"] = hostOrders[hIdx].StopPx.Value
						}
						if hostOrders[hIdx].PegOffsetValue.Valid {
							ord["pegOffsetValue"] = hostOrders[hIdx].PegOffsetValue.Value
						}
						if hostOrders[hIdx].PegPriceType.Valid {
							ord["pegPriceType"] = hostOrders[hIdx].PegPriceType.Value
						}
						if hostOrders[hIdx].OrdType.Valid {
							ord["ordType"] = hostOrders[hIdx].OrdType.Value
						}
						if hostOrders[hIdx].TimeInForce.Valid {
							ord["timeInForce"] = hostOrders[hIdx].TimeInForce.Value
						}
						if hostOrders[hIdx].ExecInst.Valid {
							ord["execInst"] = hostOrders[hIdx].ExecInst.Value
						}
						if hostOrders[hIdx].Text.Valid {
							ord["text"] = hostOrders[hIdx].Text.Value
						}
						ordersToPlace = append(ordersToPlace, ord)

					}

					//fmt.Println(ordersToPlace)
				}

			}

		}
	}

	for sIdx := range subOrders {

		orderFound := false

		for hIdx := range hostOrders {
			if hostOrders[hIdx].OrderID.Value[8:] == subOrders[sIdx].ClOrdID.Value[8:] {
				orderFound = true
			}
		}

		if !orderFound {
			cancelOrderIds = append(cancelOrderIds, subOrders[sIdx].OrderID.Value)
		}
	}

	// Step 1. Cancel Orders
	subClient.OrderCancelBulk(&cancelOrderIds)

	//////////////////////////////////////////////////////////

	{ // Leverage Update
		for hIdx := range hostPositions {

			positionFound := false

			for sIdx := range subPositions {

				if hostPositions[hIdx].Symbol.Value == subPositions[sIdx].Symbol.Value {
					positionFound = true
					if hostPositions[hIdx].CrossMargin.Value {
						if !subPositions[sIdx].CrossMargin.Value {
							subClient.UpdateLeverage(hostPositions[hIdx].Symbol.Value, 0)
						}
					} else {
						if subPositions[sIdx].CrossMargin.Value {
							subClient.UpdateLeverage(hostPositions[hIdx].Symbol.Value, hostPositions[hIdx].Leverage.Value)
						} else if hostPositions[hIdx].Leverage.Value != subPositions[sIdx].Leverage.Value {
							subClient.UpdateLeverage(hostPositions[hIdx].Symbol.Value, hostPositions[hIdx].Leverage.Value)
						}
					}
				}
			}

			if !positionFound {
				if hostPositions[hIdx].CrossMargin.Value {
					subClient.UpdateLeverage(hostPositions[hIdx].Symbol.Value, 0)
				} else {
					subClient.UpdateLeverage(hostPositions[hIdx].Symbol.Value, hostPositions[hIdx].Leverage.Value)
				}
			}
		}
	}

	////////////////////////// Calibrate Position ////////////////////////////////

	{
		for hIdx := range hostPositions {

			positionFound := false

			for sIdx := range subPositions {
				if hostPositions[hIdx].Symbol.Value == subPositions[sIdx].Symbol.Value {
					positionFound = true
					if int(hostPositions[hIdx].CurrentQty.Value*ratio) != int(subPositions[sIdx].CurrentQty.Value) {
						subClient.OrderNewMarket(hostPositions[hIdx].Symbol.Value, int(hostPositions[hIdx].CurrentQty.Value*ratio)-int(subPositions[sIdx].CurrentQty.Value))
					}
				}
			}

			if !positionFound {
				subClient.OrderNewMarket(hostPositions[hIdx].Symbol.Value, int(hostPositions[hIdx].CurrentQty.Value*ratio))
			}

		}

		for sIdx := range subPositions {

			positionFound := false

			for hIdx := range hostPositions {
				if subPositions[sIdx].Symbol.Value == hostPositions[hIdx].Symbol.Value {
					positionFound = true
				}
			}

			if !positionFound {
				subClient.OrderNewMarket(subPositions[sIdx].Symbol.Value, -int(subPositions[sIdx].CurrentQty.Value))
			}
		}

	}

	//////////////////////////////////// Calibrate Position Margin ////////////////////////////////////
	//{
	//	mHostPosition := hostClient.ActivePositions()
	//	mSubPosition := subClient.ActivePositions()
	//
	//	for hIdx := range mHostPosition {
	//		for sIdx := range mSubPosition {
	//			if mHostPosition[hIdx].Symbol.Value == mSubPosition[sIdx].Symbol.Value {
	//				if int(mHostPosition[hIdx].CurrentQty.Value*ratio) == int(mSubPosition[sIdx].CurrentQty.Value) {
	//					if (mHostPosition[hIdx].CurrentQty.Value/float32(mHostPosition[hIdx].MarkPrice.Value))/
	//						float32(mHostPosition[hIdx].Leverage.Value) < mHostPosition[hIdx].PosMargin.Value {
	//
	//						marginChange := mHostPosition[hIdx].PosMargin.Value*
	//							(mSubPosition[sIdx].CurrentQty.Value/mHostPosition[hIdx].CurrentQty.Value) -
	//							mSubPosition[sIdx].PosMargin.Value
	//
	//						if marginChange != 0 {
	//							if (mSubPosition[sIdx].CurrentQty.Value/float32(mSubPosition[sIdx].MarkPrice.Value))/
	//								float32(mSubPosition[sIdx].Leverage.Value) < mSubPosition[sIdx].PosMargin.Value+marginChange {
	//								// transfer margin to/from sub
	//								subClient.TransferMargin(mSubPosition[sIdx].Symbol.Value, int(marginChange))
	//							}
	//						}
	//
	//					}
	//				}
	//			}
	//		}
	//	}
	//}

	//////////////////////////////////// Place new orders ////////////////////////////////////
	{
		var symbols []string
		for idx := range ordersToPlace {
			if !isIn(ordersToPlace[idx]["symbol"].(string), symbols) {
				symbols = append(symbols, ordersToPlace[idx]["symbol"].(string))
			}
		}
		for idx := range symbols {
			placeNewOrders := make([]interface{}, 0, 5)
			for oIdx := range ordersToPlace {
				if ordersToPlace[oIdx]["symbol"].(string) == symbols[idx] {
					placeNewOrders = append(placeNewOrders, ordersToPlace[oIdx])
				}
			}
			subClient.OrderNewBulk(&placeNewOrders)

		}
	}

	//////////////////////////////////// Amend amended orders ////////////////////////////////////
	{
		var symbols []string
		for idx := range ordersToAmend {
			if !isIn(ordersToAmend[idx]["symbol"].(string), symbols) {
				symbols = append(symbols, ordersToAmend[idx]["symbol"].(string))
			}
		}
		for idx := range symbols {
			amendOldOrders := make([]interface{}, 0, 5)
			for aIdx := range ordersToAmend {
				if ordersToAmend[aIdx]["symbol"].(string) == symbols[idx] {
					_, ok := ordersToAmend[aIdx]["symbol"]
					if ok {
						delete(ordersToAmend[aIdx], "symbol")
						amendOldOrders = append(amendOldOrders, ordersToAmend[aIdx])
					}
					ordersToAmend[aIdx]["symbol"] = symbols[idx]
				}
			}
			subClient.OrderAmendBulk(&amendOldOrders)
		}
	}

}
