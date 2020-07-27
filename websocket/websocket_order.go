package websocket

import (
	"github.com/adi1382/Bitmex-Mirror-Bot/swagger"
)

type OrderSlice []swagger.Order

type OrderResponseData struct {
	OrderID               JSONString  `json:"orderID"`
	ClOrdID               JSONString  `json:"clOrdID,omitempty"`
	ClOrdLinkID           JSONString  `json:"clOrdLinkID,omitempty"`
	Account               JSONFloat64 `json:"account,omitempty"`
	Symbol                JSONString  `json:"symbol,omitempty"`
	Side                  JSONString  `json:"side,omitempty"`
	SimpleOrderQty        JSONFloat64 `json:"simpleOrderQty,omitempty"`
	OrderQty              JSONFloat64 `json:"orderQty,omitempty"`
	Price                 JSONFloat64 `json:"price,omitempty"`
	DisplayQty            JSONFloat64 `json:"displayQty,omitempty"`
	StopPx                JSONFloat64 `json:"stopPx,omitempty"`
	PegOffsetValue        JSONFloat64 `json:"pegOffsetValue,omitempty"`
	PegPriceType          JSONString  `json:"pegPriceType,omitempty"`
	Currency              JSONString  `json:"currency,omitempty"`
	SettlCurrency         JSONString  `json:"settlCurrency,omitempty"`
	OrdType               JSONString  `json:"ordType,omitempty"`
	TimeInForce           JSONString  `json:"timeInForce,omitempty"`
	ExecInst              JSONString  `json:"execInst,omitempty"`
	ContingencyType       JSONString  `json:"contingencyType,omitempty"`
	ExDestination         JSONString  `json:"exDestination,omitempty"`
	OrdStatus             JSONString  `json:"ordStatus,omitempty"`
	Triggered             JSONString  `json:"triggered,omitempty"`
	WorkingIndicator      JSONBool    `json:"workingIndicator,omitempty"`
	OrdRejReason          JSONString  `json:"ordRejReason,omitempty"`
	SimpleLeavesQty       JSONFloat64 `json:"simpleLeavesQty,omitempty"`
	LeavesQty             JSONFloat64 `json:"leavesQty,omitempty"`
	SimpleCumQty          JSONFloat64 `json:"simpleCumQty,omitempty"`
	CumQty                JSONFloat64 `json:"cumQty,omitempty"`
	AvgPx                 JSONFloat64 `json:"avgPx,omitempty"`
	MultiLegReportingType JSONString  `json:"multiLegReportingType,omitempty"`
	Text                  JSONString  `json:"text,omitempty"`
	TransactTime          JSONTime    `json:"transactTime,omitempty"`
	Timestamp             JSONTime    `json:"timestamp,omitempty"`
}

// -1 values signifies a null value

func (orders *OrderSlice) OrderPartial(inserts *[]OrderResponseData) {
	InfoLogger.Println("Order Partials Processing")
	*orders = nil
	for _, v := range *inserts {
		var insertOrder swagger.Order

		insertOrder.OrderID.Value = v.OrderID.Value
		insertOrder.OrderID.Valid = v.OrderID.Valid

		insertOrder.ClOrdID.Value = v.ClOrdID.Value
		insertOrder.ClOrdID.Valid = v.ClOrdID.Valid

		insertOrder.ClOrdLinkID.Value = v.ClOrdLinkID.Value
		insertOrder.ClOrdLinkID.Valid = v.ClOrdLinkID.Valid

		insertOrder.Account.Value = v.Account.Value
		insertOrder.Account.Valid = v.Account.Valid

		insertOrder.Symbol.Value = v.Symbol.Value
		insertOrder.Symbol.Valid = v.Symbol.Valid

		insertOrder.Side.Value = v.Side.Value
		insertOrder.Side.Valid = v.Side.Valid

		insertOrder.SimpleOrderQty.Value = v.SimpleOrderQty.Value
		insertOrder.SimpleOrderQty.Valid = v.SimpleOrderQty.Valid

		insertOrder.OrderQty.Value = v.OrderQty.Value
		insertOrder.OrderQty.Valid = v.OrderQty.Valid

		insertOrder.Price.Value = v.Price.Value
		insertOrder.Price.Valid = v.Price.Valid

		insertOrder.DisplayQty.Value = v.DisplayQty.Value
		insertOrder.DisplayQty.Valid = v.DisplayQty.Valid

		insertOrder.StopPx.Value = v.StopPx.Value
		insertOrder.StopPx.Valid = v.StopPx.Valid

		insertOrder.PegOffsetValue.Value = v.PegOffsetValue.Value
		insertOrder.PegOffsetValue.Valid = v.PegOffsetValue.Valid

		insertOrder.PegPriceType.Value = v.PegPriceType.Value
		insertOrder.PegPriceType.Valid = v.PegPriceType.Valid

		insertOrder.Currency.Value = v.Currency.Value
		insertOrder.Currency.Valid = v.Currency.Valid

		insertOrder.SettlCurrency.Value = v.SettlCurrency.Value
		insertOrder.SettlCurrency.Valid = v.SettlCurrency.Valid

		insertOrder.OrdType.Value = v.OrdType.Value
		insertOrder.OrdType.Valid = v.OrdType.Valid

		insertOrder.TimeInForce.Value = v.TimeInForce.Value
		insertOrder.TimeInForce.Valid = v.TimeInForce.Valid

		insertOrder.ExecInst.Value = v.ExecInst.Value
		insertOrder.ExecInst.Valid = v.ExecInst.Valid

		insertOrder.ContingencyType.Value = v.ContingencyType.Value
		insertOrder.ContingencyType.Valid = v.ContingencyType.Valid

		insertOrder.ExDestination.Value = v.ExDestination.Value
		insertOrder.ExDestination.Valid = v.ExDestination.Valid

		insertOrder.OrdStatus.Value = v.OrdStatus.Value
		insertOrder.OrdStatus.Valid = v.OrdStatus.Valid

		insertOrder.Triggered.Value = v.Triggered.Value
		insertOrder.Triggered.Valid = v.Triggered.Valid

		insertOrder.WorkingIndicator.Value = v.WorkingIndicator.Value
		insertOrder.WorkingIndicator.Valid = v.WorkingIndicator.Valid

		insertOrder.OrdRejReason.Value = v.OrdRejReason.Value
		insertOrder.OrdRejReason.Valid = v.OrdRejReason.Valid

		insertOrder.SimpleLeavesQty.Value = v.SimpleLeavesQty.Value
		insertOrder.SimpleLeavesQty.Valid = v.SimpleLeavesQty.Valid

		insertOrder.LeavesQty.Value = v.LeavesQty.Value
		insertOrder.LeavesQty.Valid = v.LeavesQty.Valid

		insertOrder.SimpleCumQty.Value = v.SimpleCumQty.Value
		insertOrder.SimpleCumQty.Valid = v.SimpleCumQty.Valid

		insertOrder.CumQty.Value = v.CumQty.Value
		insertOrder.CumQty.Valid = v.CumQty.Valid

		insertOrder.AvgPx.Value = v.AvgPx.Value
		insertOrder.AvgPx.Valid = v.AvgPx.Valid

		insertOrder.MultiLegReportingType.Value = v.MultiLegReportingType.Value
		insertOrder.MultiLegReportingType.Valid = v.MultiLegReportingType.Valid

		insertOrder.Text.Value = v.Text.Value
		insertOrder.Text.Valid = v.Text.Valid

		insertOrder.TransactTime.Value = v.TransactTime.Value
		insertOrder.TransactTime.Valid = v.TransactTime.Valid

		insertOrder.Timestamp.Value = v.Timestamp.Value
		insertOrder.Timestamp.Valid = v.Timestamp.Valid

		*orders = append(*orders, insertOrder)
	}
	InfoLogger.Println("Order Partials Processed")
}

func (orders *OrderSlice) OrderInsert(inserts *[]OrderResponseData) {

	InfoLogger.Println("Order Inserts Processing")

	for _, v := range *inserts {
		{
			d := false
			if len(*orders) > 0 {
				for _, o := range *orders {
					if v.OrderID.Value == o.OrderID.Value {
						d = true
						break
					}
				}
			}
			if d {
				continue
			}
		}

		var insertOrder swagger.Order

		insertOrder.OrderID.Value = v.OrderID.Value
		insertOrder.OrderID.Valid = v.OrderID.Valid

		insertOrder.ClOrdID.Value = v.ClOrdID.Value
		insertOrder.ClOrdID.Valid = v.ClOrdID.Valid

		insertOrder.ClOrdLinkID.Value = v.ClOrdLinkID.Value
		insertOrder.ClOrdLinkID.Valid = v.ClOrdLinkID.Valid

		insertOrder.Account.Value = v.Account.Value
		insertOrder.Account.Valid = v.Account.Valid

		insertOrder.Symbol.Value = v.Symbol.Value
		insertOrder.Symbol.Valid = v.Symbol.Valid

		insertOrder.Side.Value = v.Side.Value
		insertOrder.Side.Valid = v.Side.Valid

		insertOrder.SimpleOrderQty.Value = v.SimpleOrderQty.Value
		insertOrder.SimpleOrderQty.Valid = v.SimpleOrderQty.Valid

		insertOrder.OrderQty.Value = v.OrderQty.Value
		insertOrder.OrderQty.Valid = v.OrderQty.Valid

		insertOrder.Price.Value = v.Price.Value
		insertOrder.Price.Valid = v.Price.Valid

		insertOrder.DisplayQty.Value = v.DisplayQty.Value
		insertOrder.DisplayQty.Valid = v.DisplayQty.Valid

		insertOrder.StopPx.Value = v.StopPx.Value
		insertOrder.StopPx.Valid = v.StopPx.Valid

		insertOrder.PegOffsetValue.Value = v.PegOffsetValue.Value
		insertOrder.PegOffsetValue.Valid = v.PegOffsetValue.Valid

		insertOrder.PegPriceType.Value = v.PegPriceType.Value
		insertOrder.PegPriceType.Valid = v.PegPriceType.Valid

		insertOrder.Currency.Value = v.Currency.Value
		insertOrder.Currency.Valid = v.Currency.Valid

		insertOrder.SettlCurrency.Value = v.SettlCurrency.Value
		insertOrder.SettlCurrency.Valid = v.SettlCurrency.Valid

		insertOrder.OrdType.Value = v.OrdType.Value
		insertOrder.OrdType.Valid = v.OrdType.Valid

		insertOrder.TimeInForce.Value = v.TimeInForce.Value
		insertOrder.TimeInForce.Valid = v.TimeInForce.Valid

		insertOrder.ExecInst.Value = v.ExecInst.Value
		insertOrder.ExecInst.Valid = v.ExecInst.Valid

		insertOrder.ContingencyType.Value = v.ContingencyType.Value
		insertOrder.ContingencyType.Valid = v.ContingencyType.Valid

		insertOrder.ExDestination.Value = v.ExDestination.Value
		insertOrder.ExDestination.Valid = v.ExDestination.Valid

		insertOrder.OrdStatus.Value = v.OrdStatus.Value
		insertOrder.OrdStatus.Valid = v.OrdStatus.Valid

		insertOrder.Triggered.Value = v.Triggered.Value
		insertOrder.Triggered.Valid = v.Triggered.Valid

		insertOrder.WorkingIndicator.Value = v.WorkingIndicator.Value
		insertOrder.WorkingIndicator.Valid = v.WorkingIndicator.Valid

		insertOrder.OrdRejReason.Value = v.OrdRejReason.Value
		insertOrder.OrdRejReason.Valid = v.OrdRejReason.Valid

		insertOrder.SimpleLeavesQty.Value = v.SimpleLeavesQty.Value
		insertOrder.SimpleLeavesQty.Valid = v.SimpleLeavesQty.Valid

		insertOrder.LeavesQty.Value = v.LeavesQty.Value
		insertOrder.LeavesQty.Valid = v.LeavesQty.Valid

		insertOrder.SimpleCumQty.Value = v.SimpleCumQty.Value
		insertOrder.SimpleCumQty.Valid = v.SimpleCumQty.Valid

		insertOrder.CumQty.Value = v.CumQty.Value
		insertOrder.CumQty.Valid = v.CumQty.Valid

		insertOrder.AvgPx.Value = v.AvgPx.Value
		insertOrder.AvgPx.Valid = v.AvgPx.Valid

		insertOrder.MultiLegReportingType.Value = v.MultiLegReportingType.Value
		insertOrder.MultiLegReportingType.Valid = v.MultiLegReportingType.Valid

		insertOrder.Text.Value = v.Text.Value
		insertOrder.Text.Valid = v.Text.Valid

		insertOrder.TransactTime.Value = v.TransactTime.Value
		insertOrder.TransactTime.Valid = v.TransactTime.Valid

		insertOrder.Timestamp.Value = v.Timestamp.Value
		insertOrder.Timestamp.Valid = v.Timestamp.Valid

		//fmt.Println("New Order Inserted")
		//fmt.Println(insertOrder)
		//fmt.Printf("\n\n")
		if insertOrder.LeavesQty.Value > 0 &&
			insertOrder.OrdStatus.Value != "Filled" && insertOrder.OrdType.Value != "Market" {
			*orders = append(*orders, insertOrder)
		}
	}

	InfoLogger.Println("Order Inserts Processed")
}

func (orders *OrderSlice) OrderUpdate(updates *[]OrderResponseData) {

	InfoLogger.Println("Order Updates Processing")

	for _, u := range *updates {
		for i, v := range *orders {
			if u.OrderID.Value == v.OrderID.Value {

				if u.OrderID.Set {
					v.OrderID.Value = u.OrderID.Value
					v.OrderID.Valid = u.OrderID.Valid
				}

				if u.ClOrdID.Set {
					v.ClOrdID.Value = u.ClOrdID.Value
					v.ClOrdID.Valid = u.ClOrdID.Valid
				}

				if u.ClOrdLinkID.Set {
					v.ClOrdLinkID.Value = u.ClOrdLinkID.Value
					v.ClOrdLinkID.Valid = u.ClOrdLinkID.Valid
				}

				if u.Account.Set {
					v.Account.Value = u.Account.Value
					v.Account.Valid = u.Account.Valid
				}

				if u.Symbol.Set {
					v.Symbol.Value = u.Symbol.Value
					v.Symbol.Valid = u.Symbol.Valid
				}

				if u.Side.Set {
					v.Side.Value = u.Side.Value
					v.Side.Valid = u.Side.Valid
				}

				if u.SimpleOrderQty.Set {
					v.SimpleOrderQty.Value = u.SimpleOrderQty.Value
					v.SimpleOrderQty.Valid = u.SimpleOrderQty.Valid
				}

				if u.OrderQty.Set {
					v.OrderQty.Value = u.OrderQty.Value
					v.OrderQty.Valid = u.OrderQty.Valid
				}

				if u.Price.Set {
					v.Price.Value = u.Price.Value
					v.Price.Valid = u.Price.Valid
				}

				if u.DisplayQty.Set {
					v.DisplayQty.Value = u.DisplayQty.Value
					v.DisplayQty.Valid = u.DisplayQty.Valid
				}

				if u.StopPx.Set {
					v.StopPx.Value = u.StopPx.Value
					v.StopPx.Valid = u.StopPx.Valid
				}

				if u.PegOffsetValue.Set {
					v.PegOffsetValue.Value = u.PegOffsetValue.Value
					v.PegOffsetValue.Valid = u.PegOffsetValue.Valid
				}

				if u.PegPriceType.Set {
					v.PegPriceType.Value = u.PegPriceType.Value
					v.PegPriceType.Valid = u.PegPriceType.Valid
				}

				if u.Currency.Set {
					v.Currency.Value = u.Currency.Value
					v.Currency.Valid = u.Currency.Valid
				}

				if u.SettlCurrency.Set {
					v.SettlCurrency.Value = u.SettlCurrency.Value
					v.SettlCurrency.Valid = u.SettlCurrency.Valid
				}

				if u.OrdType.Set {
					v.OrdType.Value = u.OrdType.Value
					v.OrdType.Valid = u.OrdType.Valid
				}

				if u.TimeInForce.Set {
					v.TimeInForce.Value = u.TimeInForce.Value
					v.TimeInForce.Valid = u.TimeInForce.Valid
				}

				if u.ExecInst.Set {
					v.ExecInst.Value = u.ExecInst.Value
					v.ExecInst.Valid = u.ExecInst.Valid
				}

				if u.ContingencyType.Set {
					v.ContingencyType.Value = u.ContingencyType.Value
					v.ContingencyType.Valid = u.ContingencyType.Valid
				}

				if u.ExDestination.Set {
					v.ExDestination.Value = u.ExDestination.Value
					v.ExDestination.Valid = u.ExDestination.Valid
				}

				if u.OrdStatus.Set {
					v.OrdStatus.Value = u.OrdStatus.Value
					v.OrdStatus.Valid = u.OrdStatus.Valid
				}

				if u.Triggered.Set {
					v.Triggered.Value = u.Triggered.Value
					v.Triggered.Valid = u.Triggered.Valid
				}

				if u.WorkingIndicator.Set {
					v.WorkingIndicator.Value = u.WorkingIndicator.Value
					v.WorkingIndicator.Valid = u.WorkingIndicator.Valid
				}

				if u.OrdRejReason.Set {
					v.OrdRejReason.Value = u.OrdRejReason.Value
					v.OrdRejReason.Valid = u.OrdRejReason.Valid
				}

				if u.SimpleLeavesQty.Set {
					v.SimpleLeavesQty.Value = u.SimpleLeavesQty.Value
					v.SimpleLeavesQty.Valid = u.SimpleLeavesQty.Valid
				}

				if u.LeavesQty.Set {
					v.LeavesQty.Value = u.LeavesQty.Value
					v.LeavesQty.Valid = u.LeavesQty.Valid
				}

				if u.SimpleCumQty.Set {
					v.SimpleCumQty.Value = u.SimpleCumQty.Value
					v.SimpleCumQty.Valid = u.SimpleCumQty.Valid
				}

				if u.CumQty.Set {
					v.CumQty.Value = u.CumQty.Value
					v.CumQty.Valid = u.CumQty.Valid
				}

				if u.AvgPx.Set {
					v.AvgPx.Value = u.AvgPx.Value
					v.AvgPx.Valid = u.AvgPx.Valid
				}

				if u.MultiLegReportingType.Set {
					v.MultiLegReportingType.Value = u.MultiLegReportingType.Value
					v.MultiLegReportingType.Valid = u.MultiLegReportingType.Valid
				}

				if u.Text.Set {
					v.Text.Value = u.Text.Value
					v.Text.Valid = u.Text.Valid
				}

				if u.TransactTime.Set {
					v.TransactTime.Value = u.TransactTime.Value
					v.TransactTime.Valid = u.TransactTime.Valid
				}

				if v.LeavesQty.Value > 0 &&
					v.OrdStatus.Value != "Filled" {
					later := (*orders)[i+1:]
					*orders = append((*orders)[:i], v)
					*orders = append(*orders, later...)
				} else {
					*orders = append((*orders)[:i], (*orders)[i+1:]...)
				}
				break

			}
		}
	}

	InfoLogger.Println("Order Updates Processed")
}

func (orders *OrderSlice) OrderDelete(deletes *[]OrderResponseData) {

	InfoLogger.Println("Order Deletes Processing")
	for _, u := range *deletes {
		for i, v := range *orders {
			if u.OrderID.Value == v.OrderID.Value {
				*orders = append((*orders)[:i], (*orders)[i+1:]...)
			}
		}
	}
	InfoLogger.Println("Order Deletes Processed")
}
