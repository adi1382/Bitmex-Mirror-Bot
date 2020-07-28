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
	for v := range *inserts {
		var insertOrder swagger.Order

		insertOrder.OrderID.Value = (*inserts)[v].OrderID.Value
		insertOrder.OrderID.Valid = (*inserts)[v].OrderID.Valid

		insertOrder.ClOrdID.Value = (*inserts)[v].ClOrdID.Value
		insertOrder.ClOrdID.Valid = (*inserts)[v].ClOrdID.Valid

		insertOrder.ClOrdLinkID.Value = (*inserts)[v].ClOrdLinkID.Value
		insertOrder.ClOrdLinkID.Valid = (*inserts)[v].ClOrdLinkID.Valid

		insertOrder.Account.Value = (*inserts)[v].Account.Value
		insertOrder.Account.Valid = (*inserts)[v].Account.Valid

		insertOrder.Symbol.Value = (*inserts)[v].Symbol.Value
		insertOrder.Symbol.Valid = (*inserts)[v].Symbol.Valid

		insertOrder.Side.Value = (*inserts)[v].Side.Value
		insertOrder.Side.Valid = (*inserts)[v].Side.Valid

		insertOrder.SimpleOrderQty.Value = (*inserts)[v].SimpleOrderQty.Value
		insertOrder.SimpleOrderQty.Valid = (*inserts)[v].SimpleOrderQty.Valid

		insertOrder.OrderQty.Value = (*inserts)[v].OrderQty.Value
		insertOrder.OrderQty.Valid = (*inserts)[v].OrderQty.Valid

		insertOrder.Price.Value = (*inserts)[v].Price.Value
		insertOrder.Price.Valid = (*inserts)[v].Price.Valid

		insertOrder.DisplayQty.Value = (*inserts)[v].DisplayQty.Value
		insertOrder.DisplayQty.Valid = (*inserts)[v].DisplayQty.Valid

		insertOrder.StopPx.Value = (*inserts)[v].StopPx.Value
		insertOrder.StopPx.Valid = (*inserts)[v].StopPx.Valid

		insertOrder.PegOffsetValue.Value = (*inserts)[v].PegOffsetValue.Value
		insertOrder.PegOffsetValue.Valid = (*inserts)[v].PegOffsetValue.Valid

		insertOrder.PegPriceType.Value = (*inserts)[v].PegPriceType.Value
		insertOrder.PegPriceType.Valid = (*inserts)[v].PegPriceType.Valid

		insertOrder.Currency.Value = (*inserts)[v].Currency.Value
		insertOrder.Currency.Valid = (*inserts)[v].Currency.Valid

		insertOrder.SettlCurrency.Value = (*inserts)[v].SettlCurrency.Value
		insertOrder.SettlCurrency.Valid = (*inserts)[v].SettlCurrency.Valid

		insertOrder.OrdType.Value = (*inserts)[v].OrdType.Value
		insertOrder.OrdType.Valid = (*inserts)[v].OrdType.Valid

		insertOrder.TimeInForce.Value = (*inserts)[v].TimeInForce.Value
		insertOrder.TimeInForce.Valid = (*inserts)[v].TimeInForce.Valid

		insertOrder.ExecInst.Value = (*inserts)[v].ExecInst.Value
		insertOrder.ExecInst.Valid = (*inserts)[v].ExecInst.Valid

		insertOrder.ContingencyType.Value = (*inserts)[v].ContingencyType.Value
		insertOrder.ContingencyType.Valid = (*inserts)[v].ContingencyType.Valid

		insertOrder.ExDestination.Value = (*inserts)[v].ExDestination.Value
		insertOrder.ExDestination.Valid = (*inserts)[v].ExDestination.Valid

		insertOrder.OrdStatus.Value = (*inserts)[v].OrdStatus.Value
		insertOrder.OrdStatus.Valid = (*inserts)[v].OrdStatus.Valid

		insertOrder.Triggered.Value = (*inserts)[v].Triggered.Value
		insertOrder.Triggered.Valid = (*inserts)[v].Triggered.Valid

		insertOrder.WorkingIndicator.Value = (*inserts)[v].WorkingIndicator.Value
		insertOrder.WorkingIndicator.Valid = (*inserts)[v].WorkingIndicator.Valid

		insertOrder.OrdRejReason.Value = (*inserts)[v].OrdRejReason.Value
		insertOrder.OrdRejReason.Valid = (*inserts)[v].OrdRejReason.Valid

		insertOrder.SimpleLeavesQty.Value = (*inserts)[v].SimpleLeavesQty.Value
		insertOrder.SimpleLeavesQty.Valid = (*inserts)[v].SimpleLeavesQty.Valid

		insertOrder.LeavesQty.Value = (*inserts)[v].LeavesQty.Value
		insertOrder.LeavesQty.Valid = (*inserts)[v].LeavesQty.Valid

		insertOrder.SimpleCumQty.Value = (*inserts)[v].SimpleCumQty.Value
		insertOrder.SimpleCumQty.Valid = (*inserts)[v].SimpleCumQty.Valid

		insertOrder.CumQty.Value = (*inserts)[v].CumQty.Value
		insertOrder.CumQty.Valid = (*inserts)[v].CumQty.Valid

		insertOrder.AvgPx.Value = (*inserts)[v].AvgPx.Value
		insertOrder.AvgPx.Valid = (*inserts)[v].AvgPx.Valid

		insertOrder.MultiLegReportingType.Value = (*inserts)[v].MultiLegReportingType.Value
		insertOrder.MultiLegReportingType.Valid = (*inserts)[v].MultiLegReportingType.Valid

		insertOrder.Text.Value = (*inserts)[v].Text.Value
		insertOrder.Text.Valid = (*inserts)[v].Text.Valid

		insertOrder.TransactTime.Value = (*inserts)[v].TransactTime.Value
		insertOrder.TransactTime.Valid = (*inserts)[v].TransactTime.Valid

		insertOrder.Timestamp.Value = (*inserts)[v].Timestamp.Value
		insertOrder.Timestamp.Valid = (*inserts)[v].Timestamp.Valid

		*orders = append(*orders, insertOrder)
	}
	InfoLogger.Println("Order Partials Processed")
}

func (orders *OrderSlice) OrderInsert(inserts *[]OrderResponseData) {

	InfoLogger.Println("Order Inserts Processing")

	for i := range *inserts {
		{
			d := false
			if len(*orders) > 0 {
				for o := range *orders {
					if (*inserts)[i].OrderID.Value == (*orders)[o].OrderID.Value {
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

		insertOrder.OrderID.Value = (*inserts)[i].OrderID.Value
		insertOrder.OrderID.Valid = (*inserts)[i].OrderID.Valid

		insertOrder.ClOrdID.Value = (*inserts)[i].ClOrdID.Value
		insertOrder.ClOrdID.Valid = (*inserts)[i].ClOrdID.Valid

		insertOrder.ClOrdLinkID.Value = (*inserts)[i].ClOrdLinkID.Value
		insertOrder.ClOrdLinkID.Valid = (*inserts)[i].ClOrdLinkID.Valid

		insertOrder.Account.Value = (*inserts)[i].Account.Value
		insertOrder.Account.Valid = (*inserts)[i].Account.Valid

		insertOrder.Symbol.Value = (*inserts)[i].Symbol.Value
		insertOrder.Symbol.Valid = (*inserts)[i].Symbol.Valid

		insertOrder.Side.Value = (*inserts)[i].Side.Value
		insertOrder.Side.Valid = (*inserts)[i].Side.Valid

		insertOrder.SimpleOrderQty.Value = (*inserts)[i].SimpleOrderQty.Value
		insertOrder.SimpleOrderQty.Valid = (*inserts)[i].SimpleOrderQty.Valid

		insertOrder.OrderQty.Value = (*inserts)[i].OrderQty.Value
		insertOrder.OrderQty.Valid = (*inserts)[i].OrderQty.Valid

		insertOrder.Price.Value = (*inserts)[i].Price.Value
		insertOrder.Price.Valid = (*inserts)[i].Price.Valid

		insertOrder.DisplayQty.Value = (*inserts)[i].DisplayQty.Value
		insertOrder.DisplayQty.Valid = (*inserts)[i].DisplayQty.Valid

		insertOrder.StopPx.Value = (*inserts)[i].StopPx.Value
		insertOrder.StopPx.Valid = (*inserts)[i].StopPx.Valid

		insertOrder.PegOffsetValue.Value = (*inserts)[i].PegOffsetValue.Value
		insertOrder.PegOffsetValue.Valid = (*inserts)[i].PegOffsetValue.Valid

		insertOrder.PegPriceType.Value = (*inserts)[i].PegPriceType.Value
		insertOrder.PegPriceType.Valid = (*inserts)[i].PegPriceType.Valid

		insertOrder.Currency.Value = (*inserts)[i].Currency.Value
		insertOrder.Currency.Valid = (*inserts)[i].Currency.Valid

		insertOrder.SettlCurrency.Value = (*inserts)[i].SettlCurrency.Value
		insertOrder.SettlCurrency.Valid = (*inserts)[i].SettlCurrency.Valid

		insertOrder.OrdType.Value = (*inserts)[i].OrdType.Value
		insertOrder.OrdType.Valid = (*inserts)[i].OrdType.Valid

		insertOrder.TimeInForce.Value = (*inserts)[i].TimeInForce.Value
		insertOrder.TimeInForce.Valid = (*inserts)[i].TimeInForce.Valid

		insertOrder.ExecInst.Value = (*inserts)[i].ExecInst.Value
		insertOrder.ExecInst.Valid = (*inserts)[i].ExecInst.Valid

		insertOrder.ContingencyType.Value = (*inserts)[i].ContingencyType.Value
		insertOrder.ContingencyType.Valid = (*inserts)[i].ContingencyType.Valid

		insertOrder.ExDestination.Value = (*inserts)[i].ExDestination.Value
		insertOrder.ExDestination.Valid = (*inserts)[i].ExDestination.Valid

		insertOrder.OrdStatus.Value = (*inserts)[i].OrdStatus.Value
		insertOrder.OrdStatus.Valid = (*inserts)[i].OrdStatus.Valid

		insertOrder.Triggered.Value = (*inserts)[i].Triggered.Value
		insertOrder.Triggered.Valid = (*inserts)[i].Triggered.Valid

		insertOrder.WorkingIndicator.Value = (*inserts)[i].WorkingIndicator.Value
		insertOrder.WorkingIndicator.Valid = (*inserts)[i].WorkingIndicator.Valid

		insertOrder.OrdRejReason.Value = (*inserts)[i].OrdRejReason.Value
		insertOrder.OrdRejReason.Valid = (*inserts)[i].OrdRejReason.Valid

		insertOrder.SimpleLeavesQty.Value = (*inserts)[i].SimpleLeavesQty.Value
		insertOrder.SimpleLeavesQty.Valid = (*inserts)[i].SimpleLeavesQty.Valid

		insertOrder.LeavesQty.Value = (*inserts)[i].LeavesQty.Value
		insertOrder.LeavesQty.Valid = (*inserts)[i].LeavesQty.Valid

		insertOrder.SimpleCumQty.Value = (*inserts)[i].SimpleCumQty.Value
		insertOrder.SimpleCumQty.Valid = (*inserts)[i].SimpleCumQty.Valid

		insertOrder.CumQty.Value = (*inserts)[i].CumQty.Value
		insertOrder.CumQty.Valid = (*inserts)[i].CumQty.Valid

		insertOrder.AvgPx.Value = (*inserts)[i].AvgPx.Value
		insertOrder.AvgPx.Valid = (*inserts)[i].AvgPx.Valid

		insertOrder.MultiLegReportingType.Value = (*inserts)[i].MultiLegReportingType.Value
		insertOrder.MultiLegReportingType.Valid = (*inserts)[i].MultiLegReportingType.Valid

		insertOrder.Text.Value = (*inserts)[i].Text.Value
		insertOrder.Text.Valid = (*inserts)[i].Text.Valid

		insertOrder.TransactTime.Value = (*inserts)[i].TransactTime.Value
		insertOrder.TransactTime.Valid = (*inserts)[i].TransactTime.Valid

		insertOrder.Timestamp.Value = (*inserts)[i].Timestamp.Value
		insertOrder.Timestamp.Valid = (*inserts)[i].Timestamp.Valid

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

	for u := range *updates {
		for i := range *orders {
			if (*updates)[u].OrderID.Value == (*orders)[i].OrderID.Value {

				if (*updates)[u].OrderID.Set {
					(*orders)[i].OrderID.Value = (*updates)[u].OrderID.Value
					(*orders)[i].OrderID.Valid = (*updates)[u].OrderID.Valid
				}

				if (*updates)[u].ClOrdID.Set {
					(*orders)[i].ClOrdID.Value = (*updates)[u].ClOrdID.Value
					(*orders)[i].ClOrdID.Valid = (*updates)[u].ClOrdID.Valid
				}

				if (*updates)[u].ClOrdLinkID.Set {
					(*orders)[i].ClOrdLinkID.Value = (*updates)[u].ClOrdLinkID.Value
					(*orders)[i].ClOrdLinkID.Valid = (*updates)[u].ClOrdLinkID.Valid
				}

				if (*updates)[u].Account.Set {
					(*orders)[i].Account.Value = (*updates)[u].Account.Value
					(*orders)[i].Account.Valid = (*updates)[u].Account.Valid
				}

				if (*updates)[u].Symbol.Set {
					(*orders)[i].Symbol.Value = (*updates)[u].Symbol.Value
					(*orders)[i].Symbol.Valid = (*updates)[u].Symbol.Valid
				}

				if (*updates)[u].Side.Set {
					(*orders)[i].Side.Value = (*updates)[u].Side.Value
					(*orders)[i].Side.Valid = (*updates)[u].Side.Valid
				}

				if (*updates)[u].SimpleOrderQty.Set {
					(*orders)[i].SimpleOrderQty.Value = (*updates)[u].SimpleOrderQty.Value
					(*orders)[i].SimpleOrderQty.Valid = (*updates)[u].SimpleOrderQty.Valid
				}

				if (*updates)[u].OrderQty.Set {
					(*orders)[i].OrderQty.Value = (*updates)[u].OrderQty.Value
					(*orders)[i].OrderQty.Valid = (*updates)[u].OrderQty.Valid
				}

				if (*updates)[u].Price.Set {
					(*orders)[i].Price.Value = (*updates)[u].Price.Value
					(*orders)[i].Price.Valid = (*updates)[u].Price.Valid
				}

				if (*updates)[u].DisplayQty.Set {
					(*orders)[i].DisplayQty.Value = (*updates)[u].DisplayQty.Value
					(*orders)[i].DisplayQty.Valid = (*updates)[u].DisplayQty.Valid
				}

				if (*updates)[u].StopPx.Set {
					(*orders)[i].StopPx.Value = (*updates)[u].StopPx.Value
					(*orders)[i].StopPx.Valid = (*updates)[u].StopPx.Valid
				}

				if (*updates)[u].PegOffsetValue.Set {
					(*orders)[i].PegOffsetValue.Value = (*updates)[u].PegOffsetValue.Value
					(*orders)[i].PegOffsetValue.Valid = (*updates)[u].PegOffsetValue.Valid
				}

				if (*updates)[u].PegPriceType.Set {
					(*orders)[i].PegPriceType.Value = (*updates)[u].PegPriceType.Value
					(*orders)[i].PegPriceType.Valid = (*updates)[u].PegPriceType.Valid
				}

				if (*updates)[u].Currency.Set {
					(*orders)[i].Currency.Value = (*updates)[u].Currency.Value
					(*orders)[i].Currency.Valid = (*updates)[u].Currency.Valid
				}

				if (*updates)[u].SettlCurrency.Set {
					(*orders)[i].SettlCurrency.Value = (*updates)[u].SettlCurrency.Value
					(*orders)[i].SettlCurrency.Valid = (*updates)[u].SettlCurrency.Valid
				}

				if (*updates)[u].OrdType.Set {
					(*orders)[i].OrdType.Value = (*updates)[u].OrdType.Value
					(*orders)[i].OrdType.Valid = (*updates)[u].OrdType.Valid
				}

				if (*updates)[u].TimeInForce.Set {
					(*orders)[i].TimeInForce.Value = (*updates)[u].TimeInForce.Value
					(*orders)[i].TimeInForce.Valid = (*updates)[u].TimeInForce.Valid
				}

				if (*updates)[u].ExecInst.Set {
					(*orders)[i].ExecInst.Value = (*updates)[u].ExecInst.Value
					(*orders)[i].ExecInst.Valid = (*updates)[u].ExecInst.Valid
				}

				if (*updates)[u].ContingencyType.Set {
					(*orders)[i].ContingencyType.Value = (*updates)[u].ContingencyType.Value
					(*orders)[i].ContingencyType.Valid = (*updates)[u].ContingencyType.Valid
				}

				if (*updates)[u].ExDestination.Set {
					(*orders)[i].ExDestination.Value = (*updates)[u].ExDestination.Value
					(*orders)[i].ExDestination.Valid = (*updates)[u].ExDestination.Valid
				}

				if (*updates)[u].OrdStatus.Set {
					(*orders)[i].OrdStatus.Value = (*updates)[u].OrdStatus.Value
					(*orders)[i].OrdStatus.Valid = (*updates)[u].OrdStatus.Valid
				}

				if (*updates)[u].Triggered.Set {
					(*orders)[i].Triggered.Value = (*updates)[u].Triggered.Value
					(*orders)[i].Triggered.Valid = (*updates)[u].Triggered.Valid
				}

				if (*updates)[u].WorkingIndicator.Set {
					(*orders)[i].WorkingIndicator.Value = (*updates)[u].WorkingIndicator.Value
					(*orders)[i].WorkingIndicator.Valid = (*updates)[u].WorkingIndicator.Valid
				}

				if (*updates)[u].OrdRejReason.Set {
					(*orders)[i].OrdRejReason.Value = (*updates)[u].OrdRejReason.Value
					(*orders)[i].OrdRejReason.Valid = (*updates)[u].OrdRejReason.Valid
				}

				if (*updates)[u].SimpleLeavesQty.Set {
					(*orders)[i].SimpleLeavesQty.Value = (*updates)[u].SimpleLeavesQty.Value
					(*orders)[i].SimpleLeavesQty.Valid = (*updates)[u].SimpleLeavesQty.Valid
				}

				if (*updates)[u].LeavesQty.Set {
					(*orders)[i].LeavesQty.Value = (*updates)[u].LeavesQty.Value
					(*orders)[i].LeavesQty.Valid = (*updates)[u].LeavesQty.Valid
				}

				if (*updates)[u].SimpleCumQty.Set {
					(*orders)[i].SimpleCumQty.Value = (*updates)[u].SimpleCumQty.Value
					(*orders)[i].SimpleCumQty.Valid = (*updates)[u].SimpleCumQty.Valid
				}

				if (*updates)[u].CumQty.Set {
					(*orders)[i].CumQty.Value = (*updates)[u].CumQty.Value
					(*orders)[i].CumQty.Valid = (*updates)[u].CumQty.Valid
				}

				if (*updates)[u].AvgPx.Set {
					(*orders)[i].AvgPx.Value = (*updates)[u].AvgPx.Value
					(*orders)[i].AvgPx.Valid = (*updates)[u].AvgPx.Valid
				}

				if (*updates)[u].MultiLegReportingType.Set {
					(*orders)[i].MultiLegReportingType.Value = (*updates)[u].MultiLegReportingType.Value
					(*orders)[i].MultiLegReportingType.Valid = (*updates)[u].MultiLegReportingType.Valid
				}

				if (*updates)[u].Text.Set {
					(*orders)[i].Text.Value = (*updates)[u].Text.Value
					(*orders)[i].Text.Valid = (*updates)[u].Text.Valid
				}

				if (*updates)[u].TransactTime.Set {
					(*orders)[i].TransactTime.Value = (*updates)[u].TransactTime.Value
					(*orders)[i].TransactTime.Valid = (*updates)[u].TransactTime.Valid
				}

				if (*orders)[i].LeavesQty.Value > 0 &&
					(*orders)[i].OrdStatus.Value != "Filled" {
					later := (*orders)[i+1:]
					*orders = append((*orders)[:i], (*orders)[i])
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
	for u := range *deletes {
		for i := range *orders {
			if (*deletes)[u].OrderID.Value == (*orders)[i].OrderID.Value {
				*orders = append((*orders)[:i], (*orders)[i+1:]...)
			}
		}
	}
	InfoLogger.Println("Order Deletes Processed")
}
