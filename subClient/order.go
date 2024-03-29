package subClient

import (
	"encoding/json"
	"fmt"
	"github.com/adi1382/Bitmex-Mirror-Bot/swagger"
	"github.com/adi1382/Bitmex-Mirror-Bot/websocket"
	"net/http"
)

func (c *SubClient) getActiveOrders() websocket.OrderSlice {
	c.ordersLock.Lock()
	defer c.ordersLock.Unlock()
	return c.activeOrders
}

func (c *SubClient) orderNewMarket(symbol string, quantity int) {
	var side string
	if quantity != 0 {
		if quantity > 0 {
			side = "Buy"
		} else {
			quantity = -1 * quantity
			side = "Sell"
		}

		var newOrder swagger.OrderNewOpts
		newOrder.OrdType.Set("Market")
		newOrder.Symbol.Set(symbol)
		newOrder.OrderQty.Set(quantity)
		newOrder.Side.Set(side)
		flag := true

		c.positionsLock.Lock()

		var res swagger.Order
		var response *http.Response
		var err error
	L:
		for {
			res, response, err = c.Rest.OrderApi.OrderNew(&newOrder)
			switch c.SwaggerError(err, response) {
			case 0:
				break L
			case 1:
				continue L
			case 2:
				c.CloseConnection("Rest error")
				break L
			case 3:
				fmt.Println("Restart the bot")
				break L
			}

		}

		for idx := range c.activePositions {
			if c.activePositions[idx].Symbol.Value == res.Symbol.Value {
				if res.Side.Value == "Buy" {
					c.activePositions[idx].CurrentQty.Value += float64(quantity)
				} else {
					c.activePositions[idx].CurrentQty.Value -= float64(quantity)
				}
				flag = false
			}
		}
		if flag {
			var positionAmend swagger.Position
			positionAmend.Symbol = res.Symbol
			positionAmend.CurrentQty = res.OrderQty
			c.activePositions = append(c.activePositions, positionAmend)
			//time.Sleep(1)
		}
		c.positionsLock.Unlock()
	}
}

func (c *SubClient) orderAmendBulk(toAmend *[]interface{}) {
	if len(*toAmend) > 0 {
		message, err := json.Marshal(toAmend)
		c.checkErr(err)

		if c.restartRequired.Load() {
			return
		}

		var bulkAmend swagger.OrderAmendBulkOpts
		bulkAmend.Orders.Set(string(message))

		var res []swagger.Order
		var response *http.Response
		//var err error
	L:
		for {
			res, response, err = c.Rest.OrderApi.OrderAmendBulk(&bulkAmend)
			switch c.SwaggerError(err, response) {
			case 0:
				break L
			case 1:
				continue L
			case 2:
				fmt.Println("Remove the current sub subClient")
				c.CloseConnection("Rest Error")
				return
				//break function
			case 3:
				fmt.Println("Restart the bot")
				return
			}

		}

		c.ordersLock.Lock()
		for resIdx := range res {
			for idx := range c.activeOrders {
				if c.activeOrders[idx].OrderID.Value == res[resIdx].OrderID.Value {
					//fmt.Println("came here")
					if res[resIdx].OrderQty.Valid {
						c.activeOrders[idx].OrderQty.Value = res[resIdx].OrderQty.Value
					}
					if res[resIdx].Price.Valid {
						c.activeOrders[idx].Price.Value = res[resIdx].Price.Value
					}
					if res[resIdx].StopPx.Valid {
						c.activeOrders[idx].StopPx.Value = res[resIdx].StopPx.Value
					}
					if res[resIdx].PegOffsetValue.Valid {
						c.activeOrders[idx].PegOffsetValue.Value = res[resIdx].PegOffsetValue.Value
					}
				}
			}
		}
		c.ordersLock.Unlock()
	}
}

func (c *SubClient) orderCancelBulk(toCancelOrderIDs *[]string) {
	if len(*toCancelOrderIDs) > 0 {
		messageCancel, err := json.Marshal(toCancelOrderIDs)

		if err != nil {
			c.checkErr(err)
			c.restartRequired.Store(true)
		}

		var cancelBulk swagger.OrderCancelOpts
		cancelBulk.OrderID.Set(string(messageCancel))

		var res []swagger.Order
		var response *http.Response
		//var err error
	L:
		for {
			res, response, err = c.Rest.OrderApi.OrderCancel(&cancelBulk)
			switch c.SwaggerError(err, response) {
			case 0:
				break L
			case 1:
				continue L
			case 2:
				c.CloseConnection("Rest Error")
				return
				//break function
			case 3:
				c.logger.Error("Restart Required")
				c.restartRequired.Store(true)
				return
			}

		}

		c.ordersLock.Lock()
		for v := range res {
			if len(c.activeOrders) > 0 {
				for i := range c.activeOrders {
					if res[v].OrderID.Value == c.activeOrders[i].OrderID.Value {
						c.activeOrders = append(c.activeOrders[:i], c.activeOrders[i+1:]...)
						break
					}
				}
			}

		}

		defer c.ordersLock.Unlock()
	}

}

func (c *SubClient) orderNewBulk(toPlace *[]interface{}) {
	if len(*toPlace) > 0 {
		message, err := json.Marshal(toPlace)
		c.checkErr(err)
		var subBulkOrder swagger.OrderNewBulkOpts
		subBulkOrder.Orders.Set(string(message))

		var or []swagger.Order
		var response *http.Response
		//var err error
	L:
		for {
			or, response, err = c.Rest.OrderApi.OrderNewBulk(&subBulkOrder)
			switch c.SwaggerError(err, response) {
			case 0:
				break L
			case 1:
				continue L
			case 2:
				c.CloseConnection("Rest Error")
				return
				//break function
			case 3:
				c.restartRequired.Store(true)
				return
			}

		}

		c.ordersLock.Lock()

		for no := range or {
			{
				d := false
				if len(c.activeOrders) > 0 {
					for idx := range c.activeOrders {
						if or[no].OrderID.Value == c.activeOrders[idx].OrderID.Value {
							d = true
							break
						}
					}
				}
				if d {
					continue
				}
			}
			c.activeOrders = append(c.activeOrders, or[no])
		}
		defer c.ordersLock.Unlock()
	}
}

func (c *SubClient) orderCancelAll() {
	orderCancelStruct := swagger.OrderCancelAllOpts{}
	//var err error
L:
	for {
		_, response, err := c.Rest.OrderApi.OrderCancelAll(&orderCancelStruct)
		switch c.SwaggerError(err, response) {
		case 0:
			break L
		case 1:
			continue L
		case 2:
			fmt.Println("Remove the current sub subClient")
			c.CloseConnection("Rest Error")
			return
			//break function
		case 3:
			fmt.Println("Restart the bot")
			return
		}

	}
}
