package websocket

import (
	"github.com/adi1382/Bitmex-Mirror-Bot/swagger"
)

type PositionSlice []swagger.Position

type PositionResponseData struct {
	Account              JSONFloat32 `json:"account"`
	Symbol               JSONString  `json:"symbol"`
	Currency             JSONString  `json:"currency"`
	Underlying           JSONString  `json:"underlying,omitempty"`
	QuoteCurrency        JSONString  `json:"quoteCurrency,omitempty"`
	Commission           JSONFloat64 `json:"commission,omitempty"`
	InitMarginReq        JSONFloat64 `json:"initMarginReq,omitempty"`
	MaintMarginReq       JSONFloat64 `json:"maintMarginReq,omitempty"`
	RiskLimit            JSONFloat32 `json:"riskLimit,omitempty"`
	Leverage             JSONFloat64 `json:"leverage,omitempty"`
	CrossMargin          JSONBool    `json:"crossMargin,omitempty"`
	DeleveragePercentile JSONFloat64 `json:"deleveragePercentile,omitempty"`
	RebalancedPnl        JSONFloat32 `json:"rebalancedPnl,omitempty"`
	PrevRealisedPnl      JSONFloat32 `json:"prevRealisedPnl,omitempty"`
	PrevUnrealisedPnl    JSONFloat32 `json:"prevUnrealisedPnl,omitempty"`
	PrevClosePrice       JSONFloat64 `json:"prevClosePrice,omitempty"`
	OpeningTimestamp     JSONTime    `json:"openingTimestamp,omitempty"`
	OpeningQty           JSONFloat32 `json:"openingQty,omitempty"`
	OpeningCost          JSONFloat32 `json:"openingCost,omitempty"`
	OpeningComm          JSONFloat32 `json:"openingComm,omitempty"`
	OpenOrderBuyQty      JSONFloat32 `json:"openOrderBuyQty,omitempty"`
	OpenOrderBuyCost     JSONFloat32 `json:"openOrderBuyCost,omitempty"`
	OpenOrderBuyPremium  JSONFloat32 `json:"openOrderBuyPremium,omitempty"`
	OpenOrderSellQty     JSONFloat32 `json:"openOrderSellQty,omitempty"`
	OpenOrderSellCost    JSONFloat32 `json:"openOrderSellCost,omitempty"`
	OpenOrderSellPremium JSONFloat32 `json:"openOrderSellPremium,omitempty"`
	ExecBuyQty           JSONFloat32 `json:"execBuyQty,omitempty"`
	ExecBuyCost          JSONFloat32 `json:"execBuyCost,omitempty"`
	ExecSellQty          JSONFloat32 `json:"execSellQty,omitempty"`
	ExecSellCost         JSONFloat32 `json:"execSellCost,omitempty"`
	ExecQty              JSONFloat32 `json:"execQty,omitempty"`
	ExecCost             JSONFloat32 `json:"execCost,omitempty"`
	ExecComm             JSONFloat32 `json:"execComm,omitempty"`
	CurrentTimestamp     JSONTime    `json:"currentTimestamp,omitempty"`
	CurrentQty           JSONFloat64 `json:"currentQty,omitempty"`
	CurrentCost          JSONFloat32 `json:"currentCost,omitempty"`
	CurrentComm          JSONFloat32 `json:"currentComm,omitempty"`
	RealisedCost         JSONFloat32 `json:"realisedCost,omitempty"`
	UnrealisedCost       JSONFloat32 `json:"unrealisedCost,omitempty"`
	GrossOpenCost        JSONFloat32 `json:"grossOpenCost,omitempty"`
	GrossOpenPremium     JSONFloat32 `json:"grossOpenPremium,omitempty"`
	GrossExecCost        JSONFloat32 `json:"grossExecCost,omitempty"`
	IsOpen               JSONBool    `json:"isOpen,omitempty"`
	MarkPrice            JSONFloat64 `json:"markPrice,omitempty"`
	MarkValue            JSONFloat32 `json:"markValue,omitempty"`
	RiskValue            JSONFloat32 `json:"riskValue,omitempty"`
	HomeNotional         JSONFloat64 `json:"homeNotional,omitempty"`
	ForeignNotional      JSONFloat64 `json:"foreignNotional,omitempty"`
	PosState             JSONString  `json:"posState,omitempty"`
	PosCost              JSONFloat32 `json:"posCost,omitempty"`
	PosCost2             JSONFloat32 `json:"posCost2,omitempty"`
	PosCross             JSONFloat32 `json:"posCross,omitempty"`
	PosInit              JSONFloat32 `json:"posInit,omitempty"`
	PosComm              JSONFloat32 `json:"posComm,omitempty"`
	PosLoss              JSONFloat32 `json:"posLoss,omitempty"`
	PosMargin            JSONFloat32 `json:"posMargin,omitempty"`
	PosMaint             JSONFloat32 `json:"posMaint,omitempty"`
	PosAllowance         JSONFloat32 `json:"posAllowance,omitempty"`
	TaxableMargin        JSONFloat32 `json:"taxableMargin,omitempty"`
	InitMargin           JSONFloat32 `json:"initMargin,omitempty"`
	MaintMargin          JSONFloat32 `json:"maintMargin,omitempty"`
	SessionMargin        JSONFloat32 `json:"sessionMargin,omitempty"`
	TargetExcessMargin   JSONFloat32 `json:"targetExcessMargin,omitempty"`
	VarMargin            JSONFloat32 `json:"varMargin,omitempty"`
	RealisedGrossPnl     JSONFloat32 `json:"realisedGrossPnl,omitempty"`
	RealisedTax          JSONFloat32 `json:"realisedTax,omitempty"`
	RealisedPnl          JSONFloat32 `json:"realisedPnl,omitempty"`
	UnrealisedGrossPnl   JSONFloat32 `json:"unrealisedGrossPnl,omitempty"`
	LongBankrupt         JSONFloat32 `json:"longBankrupt,omitempty"`
	ShortBankrupt        JSONFloat32 `json:"shortBankrupt,omitempty"`
	TaxBase              JSONFloat32 `json:"taxBase,omitempty"`
	IndicativeTaxRate    JSONFloat64 `json:"indicativeTaxRate,omitempty"`
	IndicativeTax        JSONFloat32 `json:"indicativeTax,omitempty"`
	UnrealisedTax        JSONFloat32 `json:"unrealisedTax,omitempty"`
	UnrealisedPnl        JSONFloat32 `json:"unrealisedPnl,omitempty"`
	UnrealisedPnlPcnt    JSONFloat64 `json:"unrealisedPnlPcnt,omitempty"`
	UnrealisedRoePcnt    JSONFloat64 `json:"unrealisedRoePcnt,omitempty"`
	SimpleQty            JSONFloat64 `json:"simpleQty,omitempty"`
	SimpleCost           JSONFloat64 `json:"simpleCost,omitempty"`
	SimpleValue          JSONFloat64 `json:"simpleValue,omitempty"`
	SimplePnl            JSONFloat64 `json:"simplePnl,omitempty"`
	SimplePnlPcnt        JSONFloat64 `json:"simplePnlPcnt,omitempty"`
	AvgCostPrice         JSONFloat64 `json:"avgCostPrice,omitempty"`
	AvgEntryPrice        JSONFloat64 `json:"avgEntryPrice,omitempty"`
	BreakEvenPrice       JSONFloat64 `json:"breakEvenPrice,omitempty"`
	MarginCallPrice      JSONFloat64 `json:"marginCallPrice,omitempty"`
	LiquidationPrice     JSONFloat64 `json:"liquidationPrice,omitempty"`
	BankruptPrice        JSONFloat64 `json:"bankruptPrice,omitempty"`
	Timestamp            JSONTime    `json:"timestamp,omitempty"`
	LastPrice            JSONFloat64 `json:"lastPrice,omitempty"`
	LastValue            JSONFloat32 `json:"lastValue,omitempty"`
}

func (positions *PositionSlice) PositionPartial(inserts *[]PositionResponseData) {
	*positions = nil

	for i := range *inserts {
		var insertPosition swagger.Position

		insertPosition.Account.Value = (*inserts)[i].Account.Value
		insertPosition.Account.Valid = (*inserts)[i].Account.Valid

		insertPosition.Symbol.Value = (*inserts)[i].Symbol.Value
		insertPosition.Symbol.Valid = (*inserts)[i].Symbol.Valid

		insertPosition.Currency.Value = (*inserts)[i].Currency.Value
		insertPosition.Currency.Valid = (*inserts)[i].Currency.Valid

		insertPosition.Underlying.Value = (*inserts)[i].Underlying.Value
		insertPosition.Underlying.Valid = (*inserts)[i].Underlying.Valid

		insertPosition.QuoteCurrency.Value = (*inserts)[i].QuoteCurrency.Value
		insertPosition.QuoteCurrency.Valid = (*inserts)[i].QuoteCurrency.Valid

		insertPosition.Commission.Value = (*inserts)[i].Commission.Value
		insertPosition.Commission.Valid = (*inserts)[i].Commission.Valid

		insertPosition.InitMarginReq.Value = (*inserts)[i].InitMarginReq.Value
		insertPosition.InitMarginReq.Valid = (*inserts)[i].InitMarginReq.Valid

		insertPosition.MaintMarginReq.Value = (*inserts)[i].MaintMarginReq.Value
		insertPosition.MaintMarginReq.Valid = (*inserts)[i].MaintMarginReq.Valid

		insertPosition.RiskLimit.Value = (*inserts)[i].RiskLimit.Value
		insertPosition.RiskLimit.Valid = (*inserts)[i].RiskLimit.Valid

		insertPosition.CrossMargin.Value = (*inserts)[i].CrossMargin.Value
		insertPosition.CrossMargin.Valid = (*inserts)[i].CrossMargin.Valid

		insertPosition.Leverage.Value = (*inserts)[i].Leverage.Value
		insertPosition.Leverage.Valid = (*inserts)[i].Leverage.Valid

		insertPosition.DeleveragePercentile.Value = (*inserts)[i].DeleveragePercentile.Value
		insertPosition.DeleveragePercentile.Valid = (*inserts)[i].DeleveragePercentile.Valid

		insertPosition.RebalancedPnl.Value = (*inserts)[i].RebalancedPnl.Value
		insertPosition.RebalancedPnl.Valid = (*inserts)[i].RebalancedPnl.Valid

		insertPosition.PrevRealisedPnl.Value = (*inserts)[i].PrevRealisedPnl.Value
		insertPosition.PrevRealisedPnl.Valid = (*inserts)[i].PrevRealisedPnl.Valid

		insertPosition.PrevUnrealisedPnl.Value = (*inserts)[i].PrevUnrealisedPnl.Value
		insertPosition.PrevUnrealisedPnl.Valid = (*inserts)[i].PrevUnrealisedPnl.Valid

		insertPosition.PrevClosePrice.Value = (*inserts)[i].PrevClosePrice.Value
		insertPosition.PrevClosePrice.Valid = (*inserts)[i].PrevClosePrice.Valid

		insertPosition.OpeningTimestamp.Value = (*inserts)[i].OpeningTimestamp.Value
		insertPosition.OpeningTimestamp.Valid = (*inserts)[i].OpeningTimestamp.Valid

		insertPosition.OpeningQty.Value = (*inserts)[i].OpeningQty.Value
		insertPosition.OpeningQty.Valid = (*inserts)[i].OpeningQty.Valid

		insertPosition.OpeningCost.Value = (*inserts)[i].OpeningCost.Value
		insertPosition.OpeningCost.Valid = (*inserts)[i].OpeningCost.Valid

		insertPosition.OpeningComm.Value = (*inserts)[i].OpeningComm.Value
		insertPosition.OpeningComm.Valid = (*inserts)[i].OpeningComm.Valid

		insertPosition.OpenOrderBuyQty.Value = (*inserts)[i].OpenOrderBuyQty.Value
		insertPosition.OpenOrderBuyQty.Valid = (*inserts)[i].OpenOrderBuyQty.Valid

		insertPosition.OpenOrderBuyCost.Value = (*inserts)[i].OpenOrderBuyCost.Value
		insertPosition.OpenOrderBuyCost.Valid = (*inserts)[i].OpenOrderBuyCost.Valid

		insertPosition.OpenOrderBuyPremium.Value = (*inserts)[i].OpenOrderBuyPremium.Value
		insertPosition.OpenOrderBuyPremium.Valid = (*inserts)[i].OpenOrderBuyPremium.Valid

		insertPosition.OpenOrderSellQty.Value = (*inserts)[i].OpenOrderSellQty.Value
		insertPosition.OpenOrderSellQty.Valid = (*inserts)[i].OpenOrderSellQty.Valid

		insertPosition.OpenOrderSellCost.Value = (*inserts)[i].OpenOrderSellCost.Value
		insertPosition.OpenOrderSellCost.Valid = (*inserts)[i].OpenOrderSellCost.Valid

		insertPosition.OpenOrderSellPremium.Value = (*inserts)[i].OpenOrderSellPremium.Value
		insertPosition.OpenOrderSellPremium.Valid = (*inserts)[i].OpenOrderSellPremium.Valid

		insertPosition.ExecBuyQty.Value = (*inserts)[i].ExecBuyQty.Value
		insertPosition.ExecBuyQty.Valid = (*inserts)[i].ExecBuyQty.Valid

		insertPosition.ExecBuyCost.Value = (*inserts)[i].ExecBuyCost.Value
		insertPosition.ExecBuyCost.Valid = (*inserts)[i].ExecBuyCost.Valid

		insertPosition.ExecSellQty.Value = (*inserts)[i].ExecSellQty.Value
		insertPosition.ExecSellQty.Valid = (*inserts)[i].ExecSellQty.Valid

		insertPosition.ExecSellCost.Value = (*inserts)[i].ExecSellCost.Value
		insertPosition.ExecSellCost.Valid = (*inserts)[i].ExecSellCost.Valid

		insertPosition.ExecQty.Value = (*inserts)[i].ExecQty.Value
		insertPosition.ExecQty.Valid = (*inserts)[i].ExecQty.Valid

		insertPosition.ExecCost.Value = (*inserts)[i].ExecCost.Value
		insertPosition.ExecCost.Valid = (*inserts)[i].ExecCost.Valid

		insertPosition.ExecComm.Value = (*inserts)[i].ExecComm.Value
		insertPosition.ExecComm.Valid = (*inserts)[i].ExecComm.Valid

		insertPosition.CurrentTimestamp.Value = (*inserts)[i].CurrentTimestamp.Value
		insertPosition.CurrentTimestamp.Valid = (*inserts)[i].CurrentTimestamp.Valid

		insertPosition.CurrentQty.Value = (*inserts)[i].CurrentQty.Value
		insertPosition.CurrentQty.Valid = (*inserts)[i].CurrentQty.Valid

		insertPosition.CurrentCost.Value = (*inserts)[i].CurrentCost.Value
		insertPosition.CurrentCost.Valid = (*inserts)[i].CurrentCost.Valid

		insertPosition.CurrentComm.Value = (*inserts)[i].CurrentComm.Value
		insertPosition.CurrentComm.Valid = (*inserts)[i].CurrentComm.Valid

		insertPosition.RealisedCost.Value = (*inserts)[i].RealisedCost.Value
		insertPosition.RealisedCost.Valid = (*inserts)[i].RealisedCost.Valid

		insertPosition.UnrealisedCost.Value = (*inserts)[i].UnrealisedCost.Value
		insertPosition.UnrealisedCost.Valid = (*inserts)[i].UnrealisedCost.Valid

		insertPosition.GrossOpenCost.Value = (*inserts)[i].GrossOpenCost.Value
		insertPosition.GrossOpenCost.Valid = (*inserts)[i].GrossOpenCost.Valid

		insertPosition.GrossOpenPremium.Value = (*inserts)[i].GrossOpenPremium.Value
		insertPosition.GrossOpenPremium.Valid = (*inserts)[i].GrossOpenPremium.Valid

		insertPosition.GrossExecCost.Value = (*inserts)[i].GrossExecCost.Value
		insertPosition.GrossExecCost.Valid = (*inserts)[i].GrossExecCost.Valid

		insertPosition.IsOpen.Value = (*inserts)[i].IsOpen.Value
		insertPosition.IsOpen.Valid = (*inserts)[i].IsOpen.Valid

		insertPosition.MarkPrice.Value = (*inserts)[i].MarkPrice.Value
		insertPosition.MarkPrice.Valid = (*inserts)[i].MarkPrice.Valid

		insertPosition.MarkValue.Value = (*inserts)[i].MarkValue.Value
		insertPosition.MarkValue.Valid = (*inserts)[i].MarkValue.Valid

		insertPosition.RiskValue.Value = (*inserts)[i].RiskValue.Value
		insertPosition.RiskValue.Valid = (*inserts)[i].RiskValue.Valid

		insertPosition.HomeNotional.Value = (*inserts)[i].HomeNotional.Value
		insertPosition.HomeNotional.Valid = (*inserts)[i].HomeNotional.Valid

		insertPosition.ForeignNotional.Value = (*inserts)[i].ForeignNotional.Value
		insertPosition.ForeignNotional.Valid = (*inserts)[i].ForeignNotional.Valid

		insertPosition.PosState.Value = (*inserts)[i].PosState.Value
		insertPosition.PosState.Valid = (*inserts)[i].PosState.Valid

		insertPosition.PosCost.Value = (*inserts)[i].PosCost.Value
		insertPosition.PosCost.Valid = (*inserts)[i].PosCost.Valid

		insertPosition.PosCost2.Value = (*inserts)[i].PosCost2.Value
		insertPosition.PosCost2.Valid = (*inserts)[i].PosCost2.Valid

		insertPosition.PosCross.Value = (*inserts)[i].PosCross.Value
		insertPosition.PosCross.Valid = (*inserts)[i].PosCross.Valid

		insertPosition.PosInit.Value = (*inserts)[i].PosInit.Value
		insertPosition.PosInit.Valid = (*inserts)[i].PosInit.Valid

		insertPosition.PosComm.Value = (*inserts)[i].PosComm.Value
		insertPosition.PosComm.Valid = (*inserts)[i].PosComm.Valid

		insertPosition.PosLoss.Value = (*inserts)[i].PosLoss.Value
		insertPosition.PosLoss.Valid = (*inserts)[i].PosLoss.Valid

		insertPosition.PosMargin.Value = (*inserts)[i].PosMargin.Value
		insertPosition.PosMargin.Valid = (*inserts)[i].PosMargin.Valid

		insertPosition.PosMaint.Value = (*inserts)[i].PosMaint.Value
		insertPosition.PosMaint.Valid = (*inserts)[i].PosMaint.Valid

		insertPosition.PosAllowance.Value = (*inserts)[i].PosAllowance.Value
		insertPosition.PosAllowance.Valid = (*inserts)[i].PosAllowance.Valid

		insertPosition.TaxableMargin.Value = (*inserts)[i].TaxableMargin.Value
		insertPosition.TaxableMargin.Valid = (*inserts)[i].TaxableMargin.Valid

		insertPosition.InitMargin.Value = (*inserts)[i].InitMargin.Value
		insertPosition.InitMargin.Valid = (*inserts)[i].InitMargin.Valid

		insertPosition.MaintMargin.Value = (*inserts)[i].MaintMargin.Value
		insertPosition.MaintMargin.Valid = (*inserts)[i].MaintMargin.Valid

		insertPosition.SessionMargin.Value = (*inserts)[i].SessionMargin.Value
		insertPosition.SessionMargin.Valid = (*inserts)[i].SessionMargin.Valid

		insertPosition.TargetExcessMargin.Value = (*inserts)[i].TargetExcessMargin.Value
		insertPosition.TargetExcessMargin.Valid = (*inserts)[i].TargetExcessMargin.Valid

		insertPosition.VarMargin.Value = (*inserts)[i].VarMargin.Value
		insertPosition.VarMargin.Valid = (*inserts)[i].VarMargin.Valid

		insertPosition.RealisedGrossPnl.Value = (*inserts)[i].RealisedGrossPnl.Value
		insertPosition.RealisedGrossPnl.Valid = (*inserts)[i].RealisedGrossPnl.Valid

		insertPosition.RealisedTax.Value = (*inserts)[i].RealisedTax.Value
		insertPosition.RealisedTax.Valid = (*inserts)[i].RealisedTax.Valid

		insertPosition.RealisedPnl.Value = (*inserts)[i].RealisedPnl.Value
		insertPosition.RealisedPnl.Valid = (*inserts)[i].RealisedPnl.Valid

		insertPosition.UnrealisedGrossPnl.Value = (*inserts)[i].UnrealisedGrossPnl.Value
		insertPosition.UnrealisedGrossPnl.Valid = (*inserts)[i].UnrealisedGrossPnl.Valid

		insertPosition.LongBankrupt.Value = (*inserts)[i].LongBankrupt.Value
		insertPosition.LongBankrupt.Valid = (*inserts)[i].LongBankrupt.Valid

		insertPosition.ShortBankrupt.Value = (*inserts)[i].ShortBankrupt.Value
		insertPosition.ShortBankrupt.Valid = (*inserts)[i].ShortBankrupt.Valid

		insertPosition.TaxBase.Value = (*inserts)[i].TaxBase.Value
		insertPosition.TaxBase.Valid = (*inserts)[i].TaxBase.Valid

		insertPosition.IndicativeTaxRate.Value = (*inserts)[i].IndicativeTaxRate.Value
		insertPosition.IndicativeTaxRate.Valid = (*inserts)[i].IndicativeTaxRate.Valid

		insertPosition.IndicativeTax.Value = (*inserts)[i].IndicativeTax.Value
		insertPosition.IndicativeTax.Valid = (*inserts)[i].IndicativeTax.Valid

		insertPosition.UnrealisedTax.Value = (*inserts)[i].UnrealisedTax.Value
		insertPosition.UnrealisedTax.Valid = (*inserts)[i].UnrealisedTax.Valid

		insertPosition.UnrealisedPnl.Value = (*inserts)[i].UnrealisedPnl.Value
		insertPosition.UnrealisedPnl.Valid = (*inserts)[i].UnrealisedPnl.Valid

		insertPosition.UnrealisedPnlPcnt.Value = (*inserts)[i].UnrealisedPnlPcnt.Value
		insertPosition.UnrealisedPnlPcnt.Valid = (*inserts)[i].UnrealisedPnlPcnt.Valid

		insertPosition.UnrealisedRoePcnt.Value = (*inserts)[i].UnrealisedRoePcnt.Value
		insertPosition.UnrealisedRoePcnt.Valid = (*inserts)[i].UnrealisedRoePcnt.Valid

		insertPosition.SimpleQty.Value = (*inserts)[i].SimpleQty.Value
		insertPosition.SimpleQty.Valid = (*inserts)[i].SimpleQty.Valid

		insertPosition.SimpleCost.Value = (*inserts)[i].SimpleCost.Value
		insertPosition.SimpleCost.Valid = (*inserts)[i].SimpleCost.Valid

		insertPosition.SimpleValue.Value = (*inserts)[i].SimpleValue.Value
		insertPosition.SimpleValue.Valid = (*inserts)[i].SimpleValue.Valid

		insertPosition.SimplePnl.Value = (*inserts)[i].SimplePnl.Value
		insertPosition.SimplePnl.Valid = (*inserts)[i].SimplePnl.Valid

		insertPosition.SimplePnlPcnt.Value = (*inserts)[i].SimplePnlPcnt.Value
		insertPosition.SimplePnlPcnt.Valid = (*inserts)[i].SimplePnlPcnt.Valid

		insertPosition.AvgCostPrice.Value = (*inserts)[i].AvgCostPrice.Value
		insertPosition.AvgCostPrice.Valid = (*inserts)[i].AvgCostPrice.Valid

		insertPosition.AvgEntryPrice.Value = (*inserts)[i].AvgEntryPrice.Value
		insertPosition.AvgEntryPrice.Valid = (*inserts)[i].AvgEntryPrice.Valid

		insertPosition.BreakEvenPrice.Value = (*inserts)[i].BreakEvenPrice.Value
		insertPosition.BreakEvenPrice.Valid = (*inserts)[i].BreakEvenPrice.Valid

		insertPosition.MarginCallPrice.Value = (*inserts)[i].MarginCallPrice.Value
		insertPosition.MarginCallPrice.Valid = (*inserts)[i].MarginCallPrice.Valid

		insertPosition.LiquidationPrice.Value = (*inserts)[i].LiquidationPrice.Value
		insertPosition.LiquidationPrice.Valid = (*inserts)[i].LiquidationPrice.Valid

		insertPosition.BankruptPrice.Value = (*inserts)[i].BankruptPrice.Value
		insertPosition.BankruptPrice.Valid = (*inserts)[i].BankruptPrice.Valid

		insertPosition.Timestamp.Value = (*inserts)[i].Timestamp.Value
		insertPosition.Timestamp.Valid = (*inserts)[i].Timestamp.Valid

		insertPosition.LastPrice.Value = (*inserts)[i].LastPrice.Value
		insertPosition.LastPrice.Valid = (*inserts)[i].LastPrice.Valid

		insertPosition.LastValue.Value = (*inserts)[i].LastValue.Value
		insertPosition.LastValue.Valid = (*inserts)[i].LastValue.Valid

		*positions = append(*positions, insertPosition)
	}
}

func (positions *PositionSlice) PositionInsert(inserts *[]PositionResponseData) {

	var toUpdate []PositionResponseData
	var toInsert []PositionResponseData
	flag := true

	for i := range *inserts {
		for p := range *positions {
			if (*inserts)[i].Symbol.Value == (*positions)[p].Symbol.Value {
				toUpdate = append(toUpdate, (*inserts)[i])
				flag = false
				break
			}
		}
		if flag {
			toInsert = append(toInsert, (*inserts)[i])
		}
	}

	positions.PositionUpdate(&toUpdate)
	*inserts = toInsert

	for v := range *inserts {
		var insertPosition swagger.Position

		insertPosition.Account.Value = (*inserts)[v].Account.Value
		insertPosition.Account.Valid = (*inserts)[v].Account.Valid

		insertPosition.Symbol.Value = (*inserts)[v].Symbol.Value
		insertPosition.Symbol.Valid = (*inserts)[v].Symbol.Valid

		insertPosition.Currency.Value = (*inserts)[v].Currency.Value
		insertPosition.Currency.Valid = (*inserts)[v].Currency.Valid

		insertPosition.Underlying.Value = (*inserts)[v].Underlying.Value
		insertPosition.Underlying.Valid = (*inserts)[v].Underlying.Valid

		insertPosition.QuoteCurrency.Value = (*inserts)[v].QuoteCurrency.Value
		insertPosition.QuoteCurrency.Valid = (*inserts)[v].QuoteCurrency.Valid

		insertPosition.Commission.Value = (*inserts)[v].Commission.Value
		insertPosition.Commission.Valid = (*inserts)[v].Commission.Valid

		insertPosition.InitMarginReq.Value = (*inserts)[v].InitMarginReq.Value
		insertPosition.InitMarginReq.Valid = (*inserts)[v].InitMarginReq.Valid

		insertPosition.MaintMarginReq.Value = (*inserts)[v].MaintMarginReq.Value
		insertPosition.MaintMarginReq.Valid = (*inserts)[v].MaintMarginReq.Valid

		insertPosition.RiskLimit.Value = (*inserts)[v].RiskLimit.Value
		insertPosition.RiskLimit.Valid = (*inserts)[v].RiskLimit.Valid

		insertPosition.CrossMargin.Value = (*inserts)[v].CrossMargin.Value
		insertPosition.CrossMargin.Valid = (*inserts)[v].CrossMargin.Valid

		insertPosition.Leverage.Value = (*inserts)[v].Leverage.Value
		insertPosition.Leverage.Valid = (*inserts)[v].Leverage.Valid

		insertPosition.DeleveragePercentile.Value = (*inserts)[v].DeleveragePercentile.Value
		insertPosition.DeleveragePercentile.Valid = (*inserts)[v].DeleveragePercentile.Valid

		insertPosition.RebalancedPnl.Value = (*inserts)[v].RebalancedPnl.Value
		insertPosition.RebalancedPnl.Valid = (*inserts)[v].RebalancedPnl.Valid

		insertPosition.PrevRealisedPnl.Value = (*inserts)[v].PrevRealisedPnl.Value
		insertPosition.PrevRealisedPnl.Valid = (*inserts)[v].PrevRealisedPnl.Valid

		insertPosition.PrevUnrealisedPnl.Value = (*inserts)[v].PrevUnrealisedPnl.Value
		insertPosition.PrevUnrealisedPnl.Valid = (*inserts)[v].PrevUnrealisedPnl.Valid

		insertPosition.PrevClosePrice.Value = (*inserts)[v].PrevClosePrice.Value
		insertPosition.PrevClosePrice.Valid = (*inserts)[v].PrevClosePrice.Valid

		insertPosition.OpeningTimestamp.Value = (*inserts)[v].OpeningTimestamp.Value
		insertPosition.OpeningTimestamp.Valid = (*inserts)[v].OpeningTimestamp.Valid

		insertPosition.OpeningQty.Value = (*inserts)[v].OpeningQty.Value
		insertPosition.OpeningQty.Valid = (*inserts)[v].OpeningQty.Valid

		insertPosition.OpeningCost.Value = (*inserts)[v].OpeningCost.Value
		insertPosition.OpeningCost.Valid = (*inserts)[v].OpeningCost.Valid

		insertPosition.OpeningComm.Value = (*inserts)[v].OpeningComm.Value
		insertPosition.OpeningComm.Valid = (*inserts)[v].OpeningComm.Valid

		insertPosition.OpenOrderBuyQty.Value = (*inserts)[v].OpenOrderBuyQty.Value
		insertPosition.OpenOrderBuyQty.Valid = (*inserts)[v].OpenOrderBuyQty.Valid

		insertPosition.OpenOrderBuyCost.Value = (*inserts)[v].OpenOrderBuyCost.Value
		insertPosition.OpenOrderBuyCost.Valid = (*inserts)[v].OpenOrderBuyCost.Valid

		insertPosition.OpenOrderBuyPremium.Value = (*inserts)[v].OpenOrderBuyPremium.Value
		insertPosition.OpenOrderBuyPremium.Valid = (*inserts)[v].OpenOrderBuyPremium.Valid

		insertPosition.OpenOrderSellQty.Value = (*inserts)[v].OpenOrderSellQty.Value
		insertPosition.OpenOrderSellQty.Valid = (*inserts)[v].OpenOrderSellQty.Valid

		insertPosition.OpenOrderSellCost.Value = (*inserts)[v].OpenOrderSellCost.Value
		insertPosition.OpenOrderSellCost.Valid = (*inserts)[v].OpenOrderSellCost.Valid

		insertPosition.OpenOrderSellPremium.Value = (*inserts)[v].OpenOrderSellPremium.Value
		insertPosition.OpenOrderSellPremium.Valid = (*inserts)[v].OpenOrderSellPremium.Valid

		insertPosition.ExecBuyQty.Value = (*inserts)[v].ExecBuyQty.Value
		insertPosition.ExecBuyQty.Valid = (*inserts)[v].ExecBuyQty.Valid

		insertPosition.ExecBuyCost.Value = (*inserts)[v].ExecBuyCost.Value
		insertPosition.ExecBuyCost.Valid = (*inserts)[v].ExecBuyCost.Valid

		insertPosition.ExecSellQty.Value = (*inserts)[v].ExecSellQty.Value
		insertPosition.ExecSellQty.Valid = (*inserts)[v].ExecSellQty.Valid

		insertPosition.ExecSellCost.Value = (*inserts)[v].ExecSellCost.Value
		insertPosition.ExecSellCost.Valid = (*inserts)[v].ExecSellCost.Valid

		insertPosition.ExecQty.Value = (*inserts)[v].ExecQty.Value
		insertPosition.ExecQty.Valid = (*inserts)[v].ExecQty.Valid

		insertPosition.ExecCost.Value = (*inserts)[v].ExecCost.Value
		insertPosition.ExecCost.Valid = (*inserts)[v].ExecCost.Valid

		insertPosition.ExecComm.Value = (*inserts)[v].ExecComm.Value
		insertPosition.ExecComm.Valid = (*inserts)[v].ExecComm.Valid

		insertPosition.CurrentTimestamp.Value = (*inserts)[v].CurrentTimestamp.Value
		insertPosition.CurrentTimestamp.Valid = (*inserts)[v].CurrentTimestamp.Valid

		insertPosition.CurrentQty.Value = (*inserts)[v].CurrentQty.Value
		insertPosition.CurrentQty.Valid = (*inserts)[v].CurrentQty.Valid

		insertPosition.CurrentCost.Value = (*inserts)[v].CurrentCost.Value
		insertPosition.CurrentCost.Valid = (*inserts)[v].CurrentCost.Valid

		insertPosition.CurrentComm.Value = (*inserts)[v].CurrentComm.Value
		insertPosition.CurrentComm.Valid = (*inserts)[v].CurrentComm.Valid

		insertPosition.RealisedCost.Value = (*inserts)[v].RealisedCost.Value
		insertPosition.RealisedCost.Valid = (*inserts)[v].RealisedCost.Valid

		insertPosition.UnrealisedCost.Value = (*inserts)[v].UnrealisedCost.Value
		insertPosition.UnrealisedCost.Valid = (*inserts)[v].UnrealisedCost.Valid

		insertPosition.GrossOpenCost.Value = (*inserts)[v].GrossOpenCost.Value
		insertPosition.GrossOpenCost.Valid = (*inserts)[v].GrossOpenCost.Valid

		insertPosition.GrossOpenPremium.Value = (*inserts)[v].GrossOpenPremium.Value
		insertPosition.GrossOpenPremium.Valid = (*inserts)[v].GrossOpenPremium.Valid

		insertPosition.GrossExecCost.Value = (*inserts)[v].GrossExecCost.Value
		insertPosition.GrossExecCost.Valid = (*inserts)[v].GrossExecCost.Valid

		insertPosition.IsOpen.Value = (*inserts)[v].IsOpen.Value
		insertPosition.IsOpen.Valid = (*inserts)[v].IsOpen.Valid

		insertPosition.MarkPrice.Value = (*inserts)[v].MarkPrice.Value
		insertPosition.MarkPrice.Valid = (*inserts)[v].MarkPrice.Valid

		insertPosition.MarkValue.Value = (*inserts)[v].MarkValue.Value
		insertPosition.MarkValue.Valid = (*inserts)[v].MarkValue.Valid

		insertPosition.RiskValue.Value = (*inserts)[v].RiskValue.Value
		insertPosition.RiskValue.Valid = (*inserts)[v].RiskValue.Valid

		insertPosition.HomeNotional.Value = (*inserts)[v].HomeNotional.Value
		insertPosition.HomeNotional.Valid = (*inserts)[v].HomeNotional.Valid

		insertPosition.ForeignNotional.Value = (*inserts)[v].ForeignNotional.Value
		insertPosition.ForeignNotional.Valid = (*inserts)[v].ForeignNotional.Valid

		insertPosition.PosState.Value = (*inserts)[v].PosState.Value
		insertPosition.PosState.Valid = (*inserts)[v].PosState.Valid

		insertPosition.PosCost.Value = (*inserts)[v].PosCost.Value
		insertPosition.PosCost.Valid = (*inserts)[v].PosCost.Valid

		insertPosition.PosCost2.Value = (*inserts)[v].PosCost2.Value
		insertPosition.PosCost2.Valid = (*inserts)[v].PosCost2.Valid

		insertPosition.PosCross.Value = (*inserts)[v].PosCross.Value
		insertPosition.PosCross.Valid = (*inserts)[v].PosCross.Valid

		insertPosition.PosInit.Value = (*inserts)[v].PosInit.Value
		insertPosition.PosInit.Valid = (*inserts)[v].PosInit.Valid

		insertPosition.PosComm.Value = (*inserts)[v].PosComm.Value
		insertPosition.PosComm.Valid = (*inserts)[v].PosComm.Valid

		insertPosition.PosLoss.Value = (*inserts)[v].PosLoss.Value
		insertPosition.PosLoss.Valid = (*inserts)[v].PosLoss.Valid

		insertPosition.PosMargin.Value = (*inserts)[v].PosMargin.Value
		insertPosition.PosMargin.Valid = (*inserts)[v].PosMargin.Valid

		insertPosition.PosMaint.Value = (*inserts)[v].PosMaint.Value
		insertPosition.PosMaint.Valid = (*inserts)[v].PosMaint.Valid

		insertPosition.PosAllowance.Value = (*inserts)[v].PosAllowance.Value
		insertPosition.PosAllowance.Valid = (*inserts)[v].PosAllowance.Valid

		insertPosition.TaxableMargin.Value = (*inserts)[v].TaxableMargin.Value
		insertPosition.TaxableMargin.Valid = (*inserts)[v].TaxableMargin.Valid

		insertPosition.InitMargin.Value = (*inserts)[v].InitMargin.Value
		insertPosition.InitMargin.Valid = (*inserts)[v].InitMargin.Valid

		insertPosition.MaintMargin.Value = (*inserts)[v].MaintMargin.Value
		insertPosition.MaintMargin.Valid = (*inserts)[v].MaintMargin.Valid

		insertPosition.SessionMargin.Value = (*inserts)[v].SessionMargin.Value
		insertPosition.SessionMargin.Valid = (*inserts)[v].SessionMargin.Valid

		insertPosition.TargetExcessMargin.Value = (*inserts)[v].TargetExcessMargin.Value
		insertPosition.TargetExcessMargin.Valid = (*inserts)[v].TargetExcessMargin.Valid

		insertPosition.VarMargin.Value = (*inserts)[v].VarMargin.Value
		insertPosition.VarMargin.Valid = (*inserts)[v].VarMargin.Valid

		insertPosition.RealisedGrossPnl.Value = (*inserts)[v].RealisedGrossPnl.Value
		insertPosition.RealisedGrossPnl.Valid = (*inserts)[v].RealisedGrossPnl.Valid

		insertPosition.RealisedTax.Value = (*inserts)[v].RealisedTax.Value
		insertPosition.RealisedTax.Valid = (*inserts)[v].RealisedTax.Valid

		insertPosition.RealisedPnl.Value = (*inserts)[v].RealisedPnl.Value
		insertPosition.RealisedPnl.Valid = (*inserts)[v].RealisedPnl.Valid

		insertPosition.UnrealisedGrossPnl.Value = (*inserts)[v].UnrealisedGrossPnl.Value
		insertPosition.UnrealisedGrossPnl.Valid = (*inserts)[v].UnrealisedGrossPnl.Valid

		insertPosition.LongBankrupt.Value = (*inserts)[v].LongBankrupt.Value
		insertPosition.LongBankrupt.Valid = (*inserts)[v].LongBankrupt.Valid

		insertPosition.ShortBankrupt.Value = (*inserts)[v].ShortBankrupt.Value
		insertPosition.ShortBankrupt.Valid = (*inserts)[v].ShortBankrupt.Valid

		insertPosition.TaxBase.Value = (*inserts)[v].TaxBase.Value
		insertPosition.TaxBase.Valid = (*inserts)[v].TaxBase.Valid

		insertPosition.IndicativeTaxRate.Value = (*inserts)[v].IndicativeTaxRate.Value
		insertPosition.IndicativeTaxRate.Valid = (*inserts)[v].IndicativeTaxRate.Valid

		insertPosition.IndicativeTax.Value = (*inserts)[v].IndicativeTax.Value
		insertPosition.IndicativeTax.Valid = (*inserts)[v].IndicativeTax.Valid

		insertPosition.UnrealisedTax.Value = (*inserts)[v].UnrealisedTax.Value
		insertPosition.UnrealisedTax.Valid = (*inserts)[v].UnrealisedTax.Valid

		insertPosition.UnrealisedPnl.Value = (*inserts)[v].UnrealisedPnl.Value
		insertPosition.UnrealisedPnl.Valid = (*inserts)[v].UnrealisedPnl.Valid

		insertPosition.UnrealisedPnlPcnt.Value = (*inserts)[v].UnrealisedPnlPcnt.Value
		insertPosition.UnrealisedPnlPcnt.Valid = (*inserts)[v].UnrealisedPnlPcnt.Valid

		insertPosition.UnrealisedRoePcnt.Value = (*inserts)[v].UnrealisedRoePcnt.Value
		insertPosition.UnrealisedRoePcnt.Valid = (*inserts)[v].UnrealisedRoePcnt.Valid

		insertPosition.SimpleQty.Value = (*inserts)[v].SimpleQty.Value
		insertPosition.SimpleQty.Valid = (*inserts)[v].SimpleQty.Valid

		insertPosition.SimpleCost.Value = (*inserts)[v].SimpleCost.Value
		insertPosition.SimpleCost.Valid = (*inserts)[v].SimpleCost.Valid

		insertPosition.SimpleValue.Value = (*inserts)[v].SimpleValue.Value
		insertPosition.SimpleValue.Valid = (*inserts)[v].SimpleValue.Valid

		insertPosition.SimplePnl.Value = (*inserts)[v].SimplePnl.Value
		insertPosition.SimplePnl.Valid = (*inserts)[v].SimplePnl.Valid

		insertPosition.SimplePnlPcnt.Value = (*inserts)[v].SimplePnlPcnt.Value
		insertPosition.SimplePnlPcnt.Valid = (*inserts)[v].SimplePnlPcnt.Valid

		insertPosition.AvgCostPrice.Value = (*inserts)[v].AvgCostPrice.Value
		insertPosition.AvgCostPrice.Valid = (*inserts)[v].AvgCostPrice.Valid

		insertPosition.AvgEntryPrice.Value = (*inserts)[v].AvgEntryPrice.Value
		insertPosition.AvgEntryPrice.Valid = (*inserts)[v].AvgEntryPrice.Valid

		insertPosition.BreakEvenPrice.Value = (*inserts)[v].BreakEvenPrice.Value
		insertPosition.BreakEvenPrice.Valid = (*inserts)[v].BreakEvenPrice.Valid

		insertPosition.MarginCallPrice.Value = (*inserts)[v].MarginCallPrice.Value
		insertPosition.MarginCallPrice.Valid = (*inserts)[v].MarginCallPrice.Valid

		insertPosition.LiquidationPrice.Value = (*inserts)[v].LiquidationPrice.Value
		insertPosition.LiquidationPrice.Valid = (*inserts)[v].LiquidationPrice.Valid

		insertPosition.BankruptPrice.Value = (*inserts)[v].BankruptPrice.Value
		insertPosition.BankruptPrice.Valid = (*inserts)[v].BankruptPrice.Valid

		insertPosition.Timestamp.Value = (*inserts)[v].Timestamp.Value
		insertPosition.Timestamp.Valid = (*inserts)[v].Timestamp.Valid

		insertPosition.LastPrice.Value = (*inserts)[v].LastPrice.Value
		insertPosition.LastPrice.Valid = (*inserts)[v].LastPrice.Valid

		insertPosition.LastValue.Value = (*inserts)[v].LastValue.Value
		insertPosition.LastValue.Valid = (*inserts)[v].LastValue.Valid

		*positions = append(*positions, insertPosition)
	}
}

func (positions *PositionSlice) PositionUpdate(updates *[]PositionResponseData) {
	for u := range *updates {
		for v := range *positions {
			if (*updates)[u].Symbol.Value == (*positions)[v].Symbol.Value {

				if (*updates)[u].Account.Set {
					(*positions)[v].Account.Value = (*updates)[u].Account.Value
					(*positions)[v].Account.Valid = (*updates)[u].Account.Valid
				}

				if (*updates)[u].Symbol.Set {
					(*positions)[v].Symbol.Value = (*updates)[u].Symbol.Value
					(*positions)[v].Symbol.Valid = (*updates)[u].Symbol.Valid
				}

				if (*updates)[u].Currency.Set {
					(*positions)[v].Currency.Value = (*updates)[u].Currency.Value
					(*positions)[v].Currency.Valid = (*updates)[u].Currency.Valid
				}

				if (*updates)[u].Underlying.Set {
					(*positions)[v].Underlying.Value = (*updates)[u].Underlying.Value
					(*positions)[v].Underlying.Valid = (*updates)[u].Underlying.Valid
				}

				if (*updates)[u].QuoteCurrency.Set {
					(*positions)[v].QuoteCurrency.Value = (*updates)[u].QuoteCurrency.Value
					(*positions)[v].QuoteCurrency.Valid = (*updates)[u].QuoteCurrency.Valid
				}

				if (*updates)[u].Commission.Set {
					(*positions)[v].Commission.Value = (*updates)[u].Commission.Value
					(*positions)[v].Commission.Valid = (*updates)[u].Commission.Valid
				}

				if (*updates)[u].InitMarginReq.Set {
					(*positions)[v].InitMarginReq.Value = (*updates)[u].InitMarginReq.Value
					(*positions)[v].InitMarginReq.Valid = (*updates)[u].InitMarginReq.Valid
				}

				if (*updates)[u].MaintMarginReq.Set {
					(*positions)[v].MaintMarginReq.Value = (*updates)[u].MaintMarginReq.Value
					(*positions)[v].MaintMarginReq.Valid = (*updates)[u].MaintMarginReq.Valid
				}

				if (*updates)[u].RiskLimit.Set {
					(*positions)[v].RiskLimit.Value = (*updates)[u].RiskLimit.Value
					(*positions)[v].RiskLimit.Valid = (*updates)[u].RiskLimit.Valid
				}

				if (*updates)[u].Leverage.Set {
					(*positions)[v].Leverage.Value = (*updates)[u].Leverage.Value
					(*positions)[v].Leverage.Valid = (*updates)[u].Leverage.Valid
				}

				if (*updates)[u].CrossMargin.Set {
					(*positions)[v].CrossMargin.Value = (*updates)[u].CrossMargin.Value
					(*positions)[v].CrossMargin.Valid = (*updates)[u].CrossMargin.Valid
				}

				if (*updates)[u].DeleveragePercentile.Set {
					(*positions)[v].DeleveragePercentile.Value = (*updates)[u].DeleveragePercentile.Value
					(*positions)[v].DeleveragePercentile.Valid = (*updates)[u].DeleveragePercentile.Valid
				}

				if (*updates)[u].RebalancedPnl.Set {
					(*positions)[v].RebalancedPnl.Value = (*updates)[u].RebalancedPnl.Value
					(*positions)[v].RebalancedPnl.Valid = (*updates)[u].RebalancedPnl.Valid
				}

				if (*updates)[u].PrevRealisedPnl.Set {
					(*positions)[v].PrevRealisedPnl.Value = (*updates)[u].PrevRealisedPnl.Value
					(*positions)[v].PrevRealisedPnl.Valid = (*updates)[u].PrevRealisedPnl.Valid
				}

				if (*updates)[u].PrevUnrealisedPnl.Set {
					(*positions)[v].PrevUnrealisedPnl.Value = (*updates)[u].PrevUnrealisedPnl.Value
					(*positions)[v].PrevUnrealisedPnl.Valid = (*updates)[u].PrevUnrealisedPnl.Valid
				}

				if (*updates)[u].PrevClosePrice.Set {
					(*positions)[v].PrevClosePrice.Value = (*updates)[u].PrevClosePrice.Value
					(*positions)[v].PrevClosePrice.Valid = (*updates)[u].PrevClosePrice.Valid
				}

				if (*updates)[u].OpeningTimestamp.Set {
					(*positions)[v].OpeningTimestamp.Value = (*updates)[u].OpeningTimestamp.Value
					(*positions)[v].OpeningTimestamp.Valid = (*updates)[u].OpeningTimestamp.Valid
				}

				if (*updates)[u].OpeningQty.Set {
					(*positions)[v].OpeningQty.Value = (*updates)[u].OpeningQty.Value
					(*positions)[v].OpeningQty.Valid = (*updates)[u].OpeningQty.Valid
				}

				if (*updates)[u].OpeningCost.Set {
					(*positions)[v].OpeningCost.Value = (*updates)[u].OpeningCost.Value
					(*positions)[v].OpeningCost.Valid = (*updates)[u].OpeningCost.Valid
				}

				if (*updates)[u].OpeningComm.Set {
					(*positions)[v].OpeningComm.Value = (*updates)[u].OpeningComm.Value
					(*positions)[v].OpeningComm.Valid = (*updates)[u].OpeningComm.Valid
				}

				if (*updates)[u].OpenOrderBuyQty.Set {
					(*positions)[v].OpenOrderBuyQty.Value = (*updates)[u].OpenOrderBuyQty.Value
					(*positions)[v].OpenOrderBuyQty.Valid = (*updates)[u].OpenOrderBuyQty.Valid
				}

				if (*updates)[u].OpenOrderBuyCost.Set {
					(*positions)[v].OpenOrderBuyCost.Value = (*updates)[u].OpenOrderBuyCost.Value
					(*positions)[v].OpenOrderBuyCost.Valid = (*updates)[u].OpenOrderBuyCost.Valid
				}

				if (*updates)[u].OpenOrderBuyPremium.Set {
					(*positions)[v].OpenOrderBuyPremium.Value = (*updates)[u].OpenOrderBuyPremium.Value
					(*positions)[v].OpenOrderBuyPremium.Valid = (*updates)[u].OpenOrderBuyPremium.Valid
				}

				if (*updates)[u].OpenOrderSellQty.Set {
					(*positions)[v].OpenOrderSellQty.Value = (*updates)[u].OpenOrderSellQty.Value
					(*positions)[v].OpenOrderSellQty.Valid = (*updates)[u].OpenOrderSellQty.Valid
				}

				if (*updates)[u].OpenOrderSellCost.Set {
					(*positions)[v].OpenOrderSellCost.Value = (*updates)[u].OpenOrderSellCost.Value
					(*positions)[v].OpenOrderSellCost.Valid = (*updates)[u].OpenOrderSellCost.Valid
				}

				if (*updates)[u].OpenOrderSellPremium.Set {
					(*positions)[v].OpenOrderSellPremium.Value = (*updates)[u].OpenOrderSellPremium.Value
					(*positions)[v].OpenOrderSellPremium.Valid = (*updates)[u].OpenOrderSellPremium.Valid
				}

				if (*updates)[u].ExecBuyQty.Set {
					(*positions)[v].ExecBuyQty.Value = (*updates)[u].ExecBuyQty.Value
					(*positions)[v].ExecBuyQty.Valid = (*updates)[u].ExecBuyQty.Valid
				}

				if (*updates)[u].ExecBuyCost.Set {
					(*positions)[v].ExecBuyCost.Value = (*updates)[u].ExecBuyCost.Value
					(*positions)[v].ExecBuyCost.Valid = (*updates)[u].ExecBuyCost.Valid
				}

				if (*updates)[u].ExecSellQty.Set {
					(*positions)[v].ExecSellQty.Value = (*updates)[u].ExecSellQty.Value
					(*positions)[v].ExecSellQty.Valid = (*updates)[u].ExecSellQty.Valid
				}

				if (*updates)[u].ExecSellCost.Set {
					(*positions)[v].ExecSellCost.Value = (*updates)[u].ExecSellCost.Value
					(*positions)[v].ExecSellCost.Valid = (*updates)[u].ExecSellCost.Valid
				}

				if (*updates)[u].ExecCost.Set {
					(*positions)[v].ExecCost.Value = (*updates)[u].ExecCost.Value
					(*positions)[v].ExecCost.Valid = (*updates)[u].ExecCost.Valid
				}

				if (*updates)[u].ExecComm.Set {
					(*positions)[v].ExecComm.Value = (*updates)[u].ExecComm.Value
					(*positions)[v].ExecComm.Valid = (*updates)[u].ExecComm.Valid
				}

				if (*updates)[u].CurrentTimestamp.Set {
					(*positions)[v].CurrentTimestamp.Value = (*updates)[u].CurrentTimestamp.Value
					(*positions)[v].CurrentTimestamp.Valid = (*updates)[u].CurrentTimestamp.Valid
				}

				if (*updates)[u].CurrentQty.Set {
					(*positions)[v].CurrentQty.Value = (*updates)[u].CurrentQty.Value
					(*positions)[v].CurrentQty.Valid = (*updates)[u].CurrentQty.Valid
				}

				if (*updates)[u].CurrentCost.Set {
					(*positions)[v].CurrentCost.Value = (*updates)[u].CurrentCost.Value
					(*positions)[v].CurrentCost.Valid = (*updates)[u].CurrentCost.Valid
				}

				if (*updates)[u].CurrentComm.Set {
					(*positions)[v].CurrentComm.Value = (*updates)[u].CurrentComm.Value
					(*positions)[v].CurrentComm.Valid = (*updates)[u].CurrentComm.Valid
				}

				if (*updates)[u].RealisedCost.Set {
					(*positions)[v].RealisedCost.Value = (*updates)[u].RealisedCost.Value
					(*positions)[v].RealisedCost.Valid = (*updates)[u].RealisedCost.Valid
				}

				if (*updates)[u].UnrealisedCost.Set {
					(*positions)[v].UnrealisedCost.Value = (*updates)[u].UnrealisedCost.Value
					(*positions)[v].UnrealisedCost.Valid = (*updates)[u].UnrealisedCost.Valid
				}

				if (*updates)[u].GrossOpenCost.Set {
					(*positions)[v].GrossOpenCost.Value = (*updates)[u].GrossOpenCost.Value
					(*positions)[v].GrossOpenCost.Valid = (*updates)[u].GrossOpenCost.Valid
				}

				if (*updates)[u].GrossOpenPremium.Set {
					(*positions)[v].GrossOpenPremium.Value = (*updates)[u].GrossOpenPremium.Value
					(*positions)[v].GrossOpenPremium.Valid = (*updates)[u].GrossOpenPremium.Valid
				}

				if (*updates)[u].GrossExecCost.Set {
					(*positions)[v].GrossExecCost.Value = (*updates)[u].GrossExecCost.Value
					(*positions)[v].GrossExecCost.Valid = (*updates)[u].GrossExecCost.Valid
				}

				if (*updates)[u].IsOpen.Set {
					(*positions)[v].IsOpen.Value = (*updates)[u].IsOpen.Value
					(*positions)[v].IsOpen.Valid = (*updates)[u].IsOpen.Valid
				}

				if (*updates)[u].MarkPrice.Set {
					(*positions)[v].MarkPrice.Value = (*updates)[u].MarkPrice.Value
					(*positions)[v].MarkPrice.Valid = (*updates)[u].MarkPrice.Valid
				}

				if (*updates)[u].MarkValue.Set {
					(*positions)[v].MarkValue.Value = (*updates)[u].MarkValue.Value
					(*positions)[v].MarkValue.Valid = (*updates)[u].MarkValue.Valid
				}

				if (*updates)[u].RiskValue.Set {
					(*positions)[v].RiskValue.Value = (*updates)[u].RiskValue.Value
					(*positions)[v].RiskValue.Valid = (*updates)[u].RiskValue.Valid
				}

				if (*updates)[u].HomeNotional.Set {
					(*positions)[v].HomeNotional.Value = (*updates)[u].HomeNotional.Value
					(*positions)[v].HomeNotional.Valid = (*updates)[u].HomeNotional.Valid
				}

				if (*updates)[u].ForeignNotional.Set {
					(*positions)[v].ForeignNotional.Value = (*updates)[u].ForeignNotional.Value
					(*positions)[v].ForeignNotional.Valid = (*updates)[u].ForeignNotional.Valid
				}

				if (*updates)[u].PosState.Set {
					(*positions)[v].PosState.Value = (*updates)[u].PosState.Value
					(*positions)[v].PosState.Valid = (*updates)[u].PosState.Valid
				}

				if (*updates)[u].PosCost.Set {
					(*positions)[v].PosCost.Value = (*updates)[u].PosCost.Value
					(*positions)[v].PosCost.Valid = (*updates)[u].PosCost.Valid
				}

				if (*updates)[u].PosCost2.Set {
					(*positions)[v].PosCost2.Value = (*updates)[u].PosCost2.Value
					(*positions)[v].PosCost2.Valid = (*updates)[u].PosCost2.Valid
				}

				if (*updates)[u].PosCross.Set {
					(*positions)[v].PosCross.Value = (*updates)[u].PosCross.Value
					(*positions)[v].PosCross.Valid = (*updates)[u].PosCross.Valid
				}

				if (*updates)[u].PosInit.Set {
					(*positions)[v].PosInit.Value = (*updates)[u].PosInit.Value
					(*positions)[v].PosInit.Valid = (*updates)[u].PosInit.Valid
				}

				if (*updates)[u].PosComm.Set {
					(*positions)[v].PosComm.Value = (*updates)[u].PosComm.Value
					(*positions)[v].PosComm.Valid = (*updates)[u].PosComm.Valid
				}

				if (*updates)[u].PosLoss.Set {
					(*positions)[v].PosLoss.Value = (*updates)[u].PosLoss.Value
					(*positions)[v].PosLoss.Valid = (*updates)[u].PosLoss.Valid
				}

				if (*updates)[u].PosMargin.Set {
					(*positions)[v].PosMargin.Value = (*updates)[u].PosMargin.Value
					(*positions)[v].PosMargin.Valid = (*updates)[u].PosMargin.Valid
				}

				if (*updates)[u].PosMaint.Set {
					(*positions)[v].PosMaint.Value = (*updates)[u].PosMaint.Value
					(*positions)[v].PosMaint.Valid = (*updates)[u].PosMaint.Valid
				}

				if (*updates)[u].PosAllowance.Set {
					(*positions)[v].PosAllowance.Value = (*updates)[u].PosAllowance.Value
					(*positions)[v].PosAllowance.Valid = (*updates)[u].PosAllowance.Valid
				}

				if (*updates)[u].TaxableMargin.Set {
					(*positions)[v].TaxableMargin.Value = (*updates)[u].TaxableMargin.Value
					(*positions)[v].TaxableMargin.Valid = (*updates)[u].TaxableMargin.Valid
				}

				if (*updates)[u].InitMargin.Set {
					(*positions)[v].InitMargin.Value = (*updates)[u].InitMargin.Value
					(*positions)[v].InitMargin.Valid = (*updates)[u].InitMargin.Valid
				}

				if (*updates)[u].MaintMargin.Set {
					(*positions)[v].MaintMargin.Value = (*updates)[u].MaintMargin.Value
					(*positions)[v].MaintMargin.Valid = (*updates)[u].MaintMargin.Valid
				}

				if (*updates)[u].SessionMargin.Set {
					(*positions)[v].SessionMargin.Value = (*updates)[u].SessionMargin.Value
					(*positions)[v].SessionMargin.Valid = (*updates)[u].SessionMargin.Valid
				}

				if (*updates)[u].TargetExcessMargin.Set {
					(*positions)[v].TargetExcessMargin.Value = (*updates)[u].TargetExcessMargin.Value
					(*positions)[v].TargetExcessMargin.Valid = (*updates)[u].TargetExcessMargin.Valid
				}

				if (*updates)[u].VarMargin.Set {
					(*positions)[v].VarMargin.Value = (*updates)[u].VarMargin.Value
					(*positions)[v].VarMargin.Valid = (*updates)[u].VarMargin.Valid
				}

				if (*updates)[u].RealisedGrossPnl.Set {
					(*positions)[v].RealisedGrossPnl.Value = (*updates)[u].RealisedGrossPnl.Value
					(*positions)[v].RealisedGrossPnl.Valid = (*updates)[u].RealisedGrossPnl.Valid
				}

				if (*updates)[u].RealisedTax.Set {
					(*positions)[v].RealisedTax.Value = (*updates)[u].RealisedTax.Value
					(*positions)[v].RealisedTax.Valid = (*updates)[u].RealisedTax.Valid
				}

				if (*updates)[u].RealisedPnl.Set {
					(*positions)[v].RealisedPnl.Value = (*updates)[u].RealisedPnl.Value
					(*positions)[v].RealisedPnl.Valid = (*updates)[u].RealisedPnl.Valid
				}

				if (*updates)[u].UnrealisedGrossPnl.Set {
					(*positions)[v].UnrealisedGrossPnl.Value = (*updates)[u].UnrealisedGrossPnl.Value
					(*positions)[v].UnrealisedGrossPnl.Valid = (*updates)[u].UnrealisedGrossPnl.Valid
				}

				if (*updates)[u].LongBankrupt.Set {
					(*positions)[v].LongBankrupt.Value = (*updates)[u].LongBankrupt.Value
					(*positions)[v].LongBankrupt.Valid = (*updates)[u].LongBankrupt.Valid
				}

				if (*updates)[u].ShortBankrupt.Set {
					(*positions)[v].ShortBankrupt.Value = (*updates)[u].ShortBankrupt.Value
					(*positions)[v].ShortBankrupt.Valid = (*updates)[u].ShortBankrupt.Valid
				}

				if (*updates)[u].TaxBase.Set {
					(*positions)[v].TaxBase.Value = (*updates)[u].TaxBase.Value
					(*positions)[v].TaxBase.Valid = (*updates)[u].TaxBase.Valid
				}

				if (*updates)[u].IndicativeTaxRate.Set {
					(*positions)[v].IndicativeTaxRate.Value = (*updates)[u].IndicativeTaxRate.Value
					(*positions)[v].IndicativeTaxRate.Valid = (*updates)[u].IndicativeTaxRate.Valid
				}

				if (*updates)[u].IndicativeTax.Set {
					(*positions)[v].IndicativeTax.Value = (*updates)[u].IndicativeTax.Value
					(*positions)[v].IndicativeTax.Valid = (*updates)[u].IndicativeTax.Valid
				}

				if (*updates)[u].UnrealisedTax.Set {
					(*positions)[v].UnrealisedTax.Value = (*updates)[u].UnrealisedTax.Value
					(*positions)[v].UnrealisedTax.Valid = (*updates)[u].UnrealisedTax.Valid
				}

				if (*updates)[u].UnrealisedPnl.Set {
					(*positions)[v].UnrealisedPnl.Value = (*updates)[u].UnrealisedPnl.Value
					(*positions)[v].UnrealisedPnl.Valid = (*updates)[u].UnrealisedPnl.Valid
				}

				if (*updates)[u].UnrealisedPnlPcnt.Set {
					(*positions)[v].UnrealisedPnlPcnt.Value = (*updates)[u].UnrealisedPnlPcnt.Value
					(*positions)[v].UnrealisedPnlPcnt.Valid = (*updates)[u].UnrealisedPnlPcnt.Valid
				}

				if (*updates)[u].UnrealisedRoePcnt.Set {
					(*positions)[v].UnrealisedRoePcnt.Value = (*updates)[u].UnrealisedRoePcnt.Value
					(*positions)[v].UnrealisedRoePcnt.Valid = (*updates)[u].UnrealisedRoePcnt.Valid
				}

				if (*updates)[u].SimpleQty.Set {
					(*positions)[v].SimpleQty.Value = (*updates)[u].SimpleQty.Value
					(*positions)[v].SimpleQty.Valid = (*updates)[u].SimpleQty.Valid
				}

				if (*updates)[u].SimpleCost.Set {
					(*positions)[v].SimpleCost.Value = (*updates)[u].SimpleCost.Value
					(*positions)[v].SimpleCost.Valid = (*updates)[u].SimpleCost.Valid
				}

				if (*updates)[u].SimpleValue.Set {
					(*positions)[v].SimpleValue.Value = (*updates)[u].SimpleValue.Value
					(*positions)[v].SimpleValue.Valid = (*updates)[u].SimpleValue.Valid
				}

				if (*updates)[u].SimplePnl.Set {
					(*positions)[v].SimplePnl.Value = (*updates)[u].SimplePnl.Value
					(*positions)[v].SimplePnl.Valid = (*updates)[u].SimplePnl.Valid
				}

				if (*updates)[u].SimplePnlPcnt.Set {
					(*positions)[v].SimplePnlPcnt.Value = (*updates)[u].SimplePnlPcnt.Value
					(*positions)[v].SimplePnlPcnt.Valid = (*updates)[u].SimplePnlPcnt.Valid
				}

				if (*updates)[u].AvgCostPrice.Set {
					(*positions)[v].AvgCostPrice.Value = (*updates)[u].AvgCostPrice.Value
					(*positions)[v].AvgCostPrice.Valid = (*updates)[u].AvgCostPrice.Valid
				}

				if (*updates)[u].AvgEntryPrice.Set {
					(*positions)[v].AvgEntryPrice.Value = (*updates)[u].AvgEntryPrice.Value
					(*positions)[v].AvgEntryPrice.Valid = (*updates)[u].AvgEntryPrice.Valid
				}

				if (*updates)[u].BreakEvenPrice.Set {
					(*positions)[v].BreakEvenPrice.Value = (*updates)[u].BreakEvenPrice.Value
					(*positions)[v].BreakEvenPrice.Valid = (*updates)[u].BreakEvenPrice.Valid
				}

				if (*updates)[u].MarginCallPrice.Set {
					(*positions)[v].MarginCallPrice.Value = (*updates)[u].MarginCallPrice.Value
					(*positions)[v].MarginCallPrice.Valid = (*updates)[u].MarginCallPrice.Valid
				}

				if (*updates)[u].LiquidationPrice.Set {
					(*positions)[v].LiquidationPrice.Value = (*updates)[u].LiquidationPrice.Value
					(*positions)[v].LiquidationPrice.Valid = (*updates)[u].LiquidationPrice.Valid
				}

				if (*updates)[u].BankruptPrice.Set {
					(*positions)[v].BankruptPrice.Value = (*updates)[u].BankruptPrice.Value
					(*positions)[v].BankruptPrice.Valid = (*updates)[u].BankruptPrice.Valid
				}

				if (*updates)[u].Timestamp.Set {
					(*positions)[v].Timestamp.Value = (*updates)[u].Timestamp.Value
					(*positions)[v].Timestamp.Valid = (*updates)[u].Timestamp.Valid
				}

				if (*updates)[u].LastPrice.Set {
					(*positions)[v].LastPrice.Value = (*updates)[u].LastPrice.Value
					(*positions)[v].LastPrice.Valid = (*updates)[u].LastPrice.Valid
				}

				if (*updates)[u].LastValue.Set {
					(*positions)[v].LastValue.Value = (*updates)[u].LastValue.Value
					(*positions)[v].LastValue.Valid = (*updates)[u].LastValue.Valid
				}

				later := (*positions)[v+1:]
				*positions = append((*positions)[:v], (*positions)[v])
				*positions = append(*positions, later...)

			}
		}
	}
}

func (positions *PositionSlice) PositionDelete(deletes *[]PositionResponseData) {
	for u := range *deletes {
		for i := range *positions {
			if (*deletes)[u].Symbol.Value == (*positions)[i].Symbol.Value {
				*positions = append((*positions)[:i], (*positions)[i+1:]...)
			}
		}
	}
}
