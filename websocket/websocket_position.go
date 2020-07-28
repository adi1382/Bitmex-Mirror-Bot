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

	for _, i := range *inserts {
		for _, p := range *positions {
			if i.Symbol.Value == p.Symbol.Value {
				toUpdate = append(toUpdate, i)
				flag = false
				break
			}
		}
		if flag {
			toInsert = append(toInsert, i)
		}
	}

	positions.PositionUpdate(&toUpdate)
	*inserts = toInsert

	for _, v := range *inserts {
		var insertPosition swagger.Position

		insertPosition.Account.Value = v.Account.Value
		insertPosition.Account.Valid = v.Account.Valid

		insertPosition.Symbol.Value = v.Symbol.Value
		insertPosition.Symbol.Valid = v.Symbol.Valid

		insertPosition.Currency.Value = v.Currency.Value
		insertPosition.Currency.Valid = v.Currency.Valid

		insertPosition.Underlying.Value = v.Underlying.Value
		insertPosition.Underlying.Valid = v.Underlying.Valid

		insertPosition.QuoteCurrency.Value = v.QuoteCurrency.Value
		insertPosition.QuoteCurrency.Valid = v.QuoteCurrency.Valid

		insertPosition.Commission.Value = v.Commission.Value
		insertPosition.Commission.Valid = v.Commission.Valid

		insertPosition.InitMarginReq.Value = v.InitMarginReq.Value
		insertPosition.InitMarginReq.Valid = v.InitMarginReq.Valid

		insertPosition.MaintMarginReq.Value = v.MaintMarginReq.Value
		insertPosition.MaintMarginReq.Valid = v.MaintMarginReq.Valid

		insertPosition.RiskLimit.Value = v.RiskLimit.Value
		insertPosition.RiskLimit.Valid = v.RiskLimit.Valid

		insertPosition.CrossMargin.Value = v.CrossMargin.Value
		insertPosition.CrossMargin.Valid = v.CrossMargin.Valid

		insertPosition.Leverage.Value = v.Leverage.Value
		insertPosition.Leverage.Valid = v.Leverage.Valid

		insertPosition.DeleveragePercentile.Value = v.DeleveragePercentile.Value
		insertPosition.DeleveragePercentile.Valid = v.DeleveragePercentile.Valid

		insertPosition.RebalancedPnl.Value = v.RebalancedPnl.Value
		insertPosition.RebalancedPnl.Valid = v.RebalancedPnl.Valid

		insertPosition.PrevRealisedPnl.Value = v.PrevRealisedPnl.Value
		insertPosition.PrevRealisedPnl.Valid = v.PrevRealisedPnl.Valid

		insertPosition.PrevUnrealisedPnl.Value = v.PrevUnrealisedPnl.Value
		insertPosition.PrevUnrealisedPnl.Valid = v.PrevUnrealisedPnl.Valid

		insertPosition.PrevClosePrice.Value = v.PrevClosePrice.Value
		insertPosition.PrevClosePrice.Valid = v.PrevClosePrice.Valid

		insertPosition.OpeningTimestamp.Value = v.OpeningTimestamp.Value
		insertPosition.OpeningTimestamp.Valid = v.OpeningTimestamp.Valid

		insertPosition.OpeningQty.Value = v.OpeningQty.Value
		insertPosition.OpeningQty.Valid = v.OpeningQty.Valid

		insertPosition.OpeningCost.Value = v.OpeningCost.Value
		insertPosition.OpeningCost.Valid = v.OpeningCost.Valid

		insertPosition.OpeningComm.Value = v.OpeningComm.Value
		insertPosition.OpeningComm.Valid = v.OpeningComm.Valid

		insertPosition.OpenOrderBuyQty.Value = v.OpenOrderBuyQty.Value
		insertPosition.OpenOrderBuyQty.Valid = v.OpenOrderBuyQty.Valid

		insertPosition.OpenOrderBuyCost.Value = v.OpenOrderBuyCost.Value
		insertPosition.OpenOrderBuyCost.Valid = v.OpenOrderBuyCost.Valid

		insertPosition.OpenOrderBuyPremium.Value = v.OpenOrderBuyPremium.Value
		insertPosition.OpenOrderBuyPremium.Valid = v.OpenOrderBuyPremium.Valid

		insertPosition.OpenOrderSellQty.Value = v.OpenOrderSellQty.Value
		insertPosition.OpenOrderSellQty.Valid = v.OpenOrderSellQty.Valid

		insertPosition.OpenOrderSellCost.Value = v.OpenOrderSellCost.Value
		insertPosition.OpenOrderSellCost.Valid = v.OpenOrderSellCost.Valid

		insertPosition.OpenOrderSellPremium.Value = v.OpenOrderSellPremium.Value
		insertPosition.OpenOrderSellPremium.Valid = v.OpenOrderSellPremium.Valid

		insertPosition.ExecBuyQty.Value = v.ExecBuyQty.Value
		insertPosition.ExecBuyQty.Valid = v.ExecBuyQty.Valid

		insertPosition.ExecBuyCost.Value = v.ExecBuyCost.Value
		insertPosition.ExecBuyCost.Valid = v.ExecBuyCost.Valid

		insertPosition.ExecSellQty.Value = v.ExecSellQty.Value
		insertPosition.ExecSellQty.Valid = v.ExecSellQty.Valid

		insertPosition.ExecSellCost.Value = v.ExecSellCost.Value
		insertPosition.ExecSellCost.Valid = v.ExecSellCost.Valid

		insertPosition.ExecQty.Value = v.ExecQty.Value
		insertPosition.ExecQty.Valid = v.ExecQty.Valid

		insertPosition.ExecCost.Value = v.ExecCost.Value
		insertPosition.ExecCost.Valid = v.ExecCost.Valid

		insertPosition.ExecComm.Value = v.ExecComm.Value
		insertPosition.ExecComm.Valid = v.ExecComm.Valid

		insertPosition.CurrentTimestamp.Value = v.CurrentTimestamp.Value
		insertPosition.CurrentTimestamp.Valid = v.CurrentTimestamp.Valid

		insertPosition.CurrentQty.Value = v.CurrentQty.Value
		insertPosition.CurrentQty.Valid = v.CurrentQty.Valid

		insertPosition.CurrentCost.Value = v.CurrentCost.Value
		insertPosition.CurrentCost.Valid = v.CurrentCost.Valid

		insertPosition.CurrentComm.Value = v.CurrentComm.Value
		insertPosition.CurrentComm.Valid = v.CurrentComm.Valid

		insertPosition.RealisedCost.Value = v.RealisedCost.Value
		insertPosition.RealisedCost.Valid = v.RealisedCost.Valid

		insertPosition.UnrealisedCost.Value = v.UnrealisedCost.Value
		insertPosition.UnrealisedCost.Valid = v.UnrealisedCost.Valid

		insertPosition.GrossOpenCost.Value = v.GrossOpenCost.Value
		insertPosition.GrossOpenCost.Valid = v.GrossOpenCost.Valid

		insertPosition.GrossOpenPremium.Value = v.GrossOpenPremium.Value
		insertPosition.GrossOpenPremium.Valid = v.GrossOpenPremium.Valid

		insertPosition.GrossExecCost.Value = v.GrossExecCost.Value
		insertPosition.GrossExecCost.Valid = v.GrossExecCost.Valid

		insertPosition.IsOpen.Value = v.IsOpen.Value
		insertPosition.IsOpen.Valid = v.IsOpen.Valid

		insertPosition.MarkPrice.Value = v.MarkPrice.Value
		insertPosition.MarkPrice.Valid = v.MarkPrice.Valid

		insertPosition.MarkValue.Value = v.MarkValue.Value
		insertPosition.MarkValue.Valid = v.MarkValue.Valid

		insertPosition.RiskValue.Value = v.RiskValue.Value
		insertPosition.RiskValue.Valid = v.RiskValue.Valid

		insertPosition.HomeNotional.Value = v.HomeNotional.Value
		insertPosition.HomeNotional.Valid = v.HomeNotional.Valid

		insertPosition.ForeignNotional.Value = v.ForeignNotional.Value
		insertPosition.ForeignNotional.Valid = v.ForeignNotional.Valid

		insertPosition.PosState.Value = v.PosState.Value
		insertPosition.PosState.Valid = v.PosState.Valid

		insertPosition.PosCost.Value = v.PosCost.Value
		insertPosition.PosCost.Valid = v.PosCost.Valid

		insertPosition.PosCost2.Value = v.PosCost2.Value
		insertPosition.PosCost2.Valid = v.PosCost2.Valid

		insertPosition.PosCross.Value = v.PosCross.Value
		insertPosition.PosCross.Valid = v.PosCross.Valid

		insertPosition.PosInit.Value = v.PosInit.Value
		insertPosition.PosInit.Valid = v.PosInit.Valid

		insertPosition.PosComm.Value = v.PosComm.Value
		insertPosition.PosComm.Valid = v.PosComm.Valid

		insertPosition.PosLoss.Value = v.PosLoss.Value
		insertPosition.PosLoss.Valid = v.PosLoss.Valid

		insertPosition.PosMargin.Value = v.PosMargin.Value
		insertPosition.PosMargin.Valid = v.PosMargin.Valid

		insertPosition.PosMaint.Value = v.PosMaint.Value
		insertPosition.PosMaint.Valid = v.PosMaint.Valid

		insertPosition.PosAllowance.Value = v.PosAllowance.Value
		insertPosition.PosAllowance.Valid = v.PosAllowance.Valid

		insertPosition.TaxableMargin.Value = v.TaxableMargin.Value
		insertPosition.TaxableMargin.Valid = v.TaxableMargin.Valid

		insertPosition.InitMargin.Value = v.InitMargin.Value
		insertPosition.InitMargin.Valid = v.InitMargin.Valid

		insertPosition.MaintMargin.Value = v.MaintMargin.Value
		insertPosition.MaintMargin.Valid = v.MaintMargin.Valid

		insertPosition.SessionMargin.Value = v.SessionMargin.Value
		insertPosition.SessionMargin.Valid = v.SessionMargin.Valid

		insertPosition.TargetExcessMargin.Value = v.TargetExcessMargin.Value
		insertPosition.TargetExcessMargin.Valid = v.TargetExcessMargin.Valid

		insertPosition.VarMargin.Value = v.VarMargin.Value
		insertPosition.VarMargin.Valid = v.VarMargin.Valid

		insertPosition.RealisedGrossPnl.Value = v.RealisedGrossPnl.Value
		insertPosition.RealisedGrossPnl.Valid = v.RealisedGrossPnl.Valid

		insertPosition.RealisedTax.Value = v.RealisedTax.Value
		insertPosition.RealisedTax.Valid = v.RealisedTax.Valid

		insertPosition.RealisedPnl.Value = v.RealisedPnl.Value
		insertPosition.RealisedPnl.Valid = v.RealisedPnl.Valid

		insertPosition.UnrealisedGrossPnl.Value = v.UnrealisedGrossPnl.Value
		insertPosition.UnrealisedGrossPnl.Valid = v.UnrealisedGrossPnl.Valid

		insertPosition.LongBankrupt.Value = v.LongBankrupt.Value
		insertPosition.LongBankrupt.Valid = v.LongBankrupt.Valid

		insertPosition.ShortBankrupt.Value = v.ShortBankrupt.Value
		insertPosition.ShortBankrupt.Valid = v.ShortBankrupt.Valid

		insertPosition.TaxBase.Value = v.TaxBase.Value
		insertPosition.TaxBase.Valid = v.TaxBase.Valid

		insertPosition.IndicativeTaxRate.Value = v.IndicativeTaxRate.Value
		insertPosition.IndicativeTaxRate.Valid = v.IndicativeTaxRate.Valid

		insertPosition.IndicativeTax.Value = v.IndicativeTax.Value
		insertPosition.IndicativeTax.Valid = v.IndicativeTax.Valid

		insertPosition.UnrealisedTax.Value = v.UnrealisedTax.Value
		insertPosition.UnrealisedTax.Valid = v.UnrealisedTax.Valid

		insertPosition.UnrealisedPnl.Value = v.UnrealisedPnl.Value
		insertPosition.UnrealisedPnl.Valid = v.UnrealisedPnl.Valid

		insertPosition.UnrealisedPnlPcnt.Value = v.UnrealisedPnlPcnt.Value
		insertPosition.UnrealisedPnlPcnt.Valid = v.UnrealisedPnlPcnt.Valid

		insertPosition.UnrealisedRoePcnt.Value = v.UnrealisedRoePcnt.Value
		insertPosition.UnrealisedRoePcnt.Valid = v.UnrealisedRoePcnt.Valid

		insertPosition.SimpleQty.Value = v.SimpleQty.Value
		insertPosition.SimpleQty.Valid = v.SimpleQty.Valid

		insertPosition.SimpleCost.Value = v.SimpleCost.Value
		insertPosition.SimpleCost.Valid = v.SimpleCost.Valid

		insertPosition.SimpleValue.Value = v.SimpleValue.Value
		insertPosition.SimpleValue.Valid = v.SimpleValue.Valid

		insertPosition.SimplePnl.Value = v.SimplePnl.Value
		insertPosition.SimplePnl.Valid = v.SimplePnl.Valid

		insertPosition.SimplePnlPcnt.Value = v.SimplePnlPcnt.Value
		insertPosition.SimplePnlPcnt.Valid = v.SimplePnlPcnt.Valid

		insertPosition.AvgCostPrice.Value = v.AvgCostPrice.Value
		insertPosition.AvgCostPrice.Valid = v.AvgCostPrice.Valid

		insertPosition.AvgEntryPrice.Value = v.AvgEntryPrice.Value
		insertPosition.AvgEntryPrice.Valid = v.AvgEntryPrice.Valid

		insertPosition.BreakEvenPrice.Value = v.BreakEvenPrice.Value
		insertPosition.BreakEvenPrice.Valid = v.BreakEvenPrice.Valid

		insertPosition.MarginCallPrice.Value = v.MarginCallPrice.Value
		insertPosition.MarginCallPrice.Valid = v.MarginCallPrice.Valid

		insertPosition.LiquidationPrice.Value = v.LiquidationPrice.Value
		insertPosition.LiquidationPrice.Valid = v.LiquidationPrice.Valid

		insertPosition.BankruptPrice.Value = v.BankruptPrice.Value
		insertPosition.BankruptPrice.Valid = v.BankruptPrice.Valid

		insertPosition.Timestamp.Value = v.Timestamp.Value
		insertPosition.Timestamp.Valid = v.Timestamp.Valid

		insertPosition.LastPrice.Value = v.LastPrice.Value
		insertPosition.LastPrice.Valid = v.LastPrice.Valid

		insertPosition.LastValue.Value = v.LastValue.Value
		insertPosition.LastValue.Valid = v.LastValue.Valid

		*positions = append(*positions, insertPosition)
	}
}

func (positions *PositionSlice) PositionUpdate(updates *[]PositionResponseData) {
	for _, u := range *updates {
		for i, v := range *positions {
			if u.Symbol.Value == v.Symbol.Value {

				if u.Account.Set {
					v.Account.Value = u.Account.Value
					v.Account.Valid = u.Account.Valid
				}

				if u.Symbol.Set {
					v.Symbol.Value = u.Symbol.Value
					v.Symbol.Valid = u.Symbol.Valid
				}

				if u.Currency.Set {
					v.Currency.Value = u.Currency.Value
					v.Currency.Valid = u.Currency.Valid
				}

				if u.Underlying.Set {
					v.Underlying.Value = u.Underlying.Value
					v.Underlying.Valid = u.Underlying.Valid
				}

				if u.QuoteCurrency.Set {
					v.QuoteCurrency.Value = u.QuoteCurrency.Value
					v.QuoteCurrency.Valid = u.QuoteCurrency.Valid
				}

				if u.Commission.Set {
					v.Commission.Value = u.Commission.Value
					v.Commission.Valid = u.Commission.Valid
				}

				if u.InitMarginReq.Set {
					v.InitMarginReq.Value = u.InitMarginReq.Value
					v.InitMarginReq.Valid = u.InitMarginReq.Valid
				}

				if u.MaintMarginReq.Set {
					v.MaintMarginReq.Value = u.MaintMarginReq.Value
					v.MaintMarginReq.Valid = u.MaintMarginReq.Valid
				}

				if u.RiskLimit.Set {
					v.RiskLimit.Value = u.RiskLimit.Value
					v.RiskLimit.Valid = u.RiskLimit.Valid
				}

				if u.Leverage.Set {
					v.Leverage.Value = u.Leverage.Value
					v.Leverage.Valid = u.Leverage.Valid
				}

				if u.CrossMargin.Set {
					v.CrossMargin.Value = u.CrossMargin.Value
					v.CrossMargin.Valid = u.CrossMargin.Valid
				}

				if u.DeleveragePercentile.Set {
					v.DeleveragePercentile.Value = u.DeleveragePercentile.Value
					v.DeleveragePercentile.Valid = u.DeleveragePercentile.Valid
				}

				if u.RebalancedPnl.Set {
					v.RebalancedPnl.Value = u.RebalancedPnl.Value
					v.RebalancedPnl.Valid = u.RebalancedPnl.Valid
				}

				if u.PrevRealisedPnl.Set {
					v.PrevRealisedPnl.Value = u.PrevRealisedPnl.Value
					v.PrevRealisedPnl.Valid = u.PrevRealisedPnl.Valid
				}

				if u.PrevUnrealisedPnl.Set {
					v.PrevUnrealisedPnl.Value = u.PrevUnrealisedPnl.Value
					v.PrevUnrealisedPnl.Valid = u.PrevUnrealisedPnl.Valid
				}

				if u.PrevClosePrice.Set {
					v.PrevClosePrice.Value = u.PrevClosePrice.Value
					v.PrevClosePrice.Valid = u.PrevClosePrice.Valid
				}

				if u.OpeningTimestamp.Set {
					v.OpeningTimestamp.Value = u.OpeningTimestamp.Value
					v.OpeningTimestamp.Valid = u.OpeningTimestamp.Valid
				}

				if u.OpeningQty.Set {
					v.OpeningQty.Value = u.OpeningQty.Value
					v.OpeningQty.Valid = u.OpeningQty.Valid
				}

				if u.OpeningCost.Set {
					v.OpeningCost.Value = u.OpeningCost.Value
					v.OpeningCost.Valid = u.OpeningCost.Valid
				}

				if u.OpeningComm.Set {
					v.OpeningComm.Value = u.OpeningComm.Value
					v.OpeningComm.Valid = u.OpeningComm.Valid
				}

				if u.OpenOrderBuyQty.Set {
					v.OpenOrderBuyQty.Value = u.OpenOrderBuyQty.Value
					v.OpenOrderBuyQty.Valid = u.OpenOrderBuyQty.Valid
				}

				if u.OpenOrderBuyCost.Set {
					v.OpenOrderBuyCost.Value = u.OpenOrderBuyCost.Value
					v.OpenOrderBuyCost.Valid = u.OpenOrderBuyCost.Valid
				}

				if u.OpenOrderBuyPremium.Set {
					v.OpenOrderBuyPremium.Value = u.OpenOrderBuyPremium.Value
					v.OpenOrderBuyPremium.Valid = u.OpenOrderBuyPremium.Valid
				}

				if u.OpenOrderSellQty.Set {
					v.OpenOrderSellQty.Value = u.OpenOrderSellQty.Value
					v.OpenOrderSellQty.Valid = u.OpenOrderSellQty.Valid
				}

				if u.OpenOrderSellCost.Set {
					v.OpenOrderSellCost.Value = u.OpenOrderSellCost.Value
					v.OpenOrderSellCost.Valid = u.OpenOrderSellCost.Valid
				}

				if u.OpenOrderSellPremium.Set {
					v.OpenOrderSellPremium.Value = u.OpenOrderSellPremium.Value
					v.OpenOrderSellPremium.Valid = u.OpenOrderSellPremium.Valid
				}

				if u.ExecBuyQty.Set {
					v.ExecBuyQty.Value = u.ExecBuyQty.Value
					v.ExecBuyQty.Valid = u.ExecBuyQty.Valid
				}

				if u.ExecBuyCost.Set {
					v.ExecBuyCost.Value = u.ExecBuyCost.Value
					v.ExecBuyCost.Valid = u.ExecBuyCost.Valid
				}

				if u.ExecSellQty.Set {
					v.ExecSellQty.Value = u.ExecSellQty.Value
					v.ExecSellQty.Valid = u.ExecSellQty.Valid
				}

				if u.ExecSellCost.Set {
					v.ExecSellCost.Value = u.ExecSellCost.Value
					v.ExecSellCost.Valid = u.ExecSellCost.Valid
				}

				if u.ExecCost.Set {
					v.ExecCost.Value = u.ExecCost.Value
					v.ExecCost.Valid = u.ExecCost.Valid
				}

				if u.ExecComm.Set {
					v.ExecComm.Value = u.ExecComm.Value
					v.ExecComm.Valid = u.ExecComm.Valid
				}

				if u.CurrentTimestamp.Set {
					v.CurrentTimestamp.Value = u.CurrentTimestamp.Value
					v.CurrentTimestamp.Valid = u.CurrentTimestamp.Valid
				}

				if u.CurrentQty.Set {
					v.CurrentQty.Value = u.CurrentQty.Value
					v.CurrentQty.Valid = u.CurrentQty.Valid
				}

				if u.CurrentCost.Set {
					v.CurrentCost.Value = u.CurrentCost.Value
					v.CurrentCost.Valid = u.CurrentCost.Valid
				}

				if u.CurrentComm.Set {
					v.CurrentComm.Value = u.CurrentComm.Value
					v.CurrentComm.Valid = u.CurrentComm.Valid
				}

				if u.RealisedCost.Set {
					v.RealisedCost.Value = u.RealisedCost.Value
					v.RealisedCost.Valid = u.RealisedCost.Valid
				}

				if u.UnrealisedCost.Set {
					v.UnrealisedCost.Value = u.UnrealisedCost.Value
					v.UnrealisedCost.Valid = u.UnrealisedCost.Valid
				}

				if u.GrossOpenCost.Set {
					v.GrossOpenCost.Value = u.GrossOpenCost.Value
					v.GrossOpenCost.Valid = u.GrossOpenCost.Valid
				}

				if u.GrossOpenPremium.Set {
					v.GrossOpenPremium.Value = u.GrossOpenPremium.Value
					v.GrossOpenPremium.Valid = u.GrossOpenPremium.Valid
				}

				if u.GrossExecCost.Set {
					v.GrossExecCost.Value = u.GrossExecCost.Value
					v.GrossExecCost.Valid = u.GrossExecCost.Valid
				}

				if u.IsOpen.Set {
					v.IsOpen.Value = u.IsOpen.Value
					v.IsOpen.Valid = u.IsOpen.Valid
				}

				if u.MarkPrice.Set {
					v.MarkPrice.Value = u.MarkPrice.Value
					v.MarkPrice.Valid = u.MarkPrice.Valid
				}

				if u.MarkValue.Set {
					v.MarkValue.Value = u.MarkValue.Value
					v.MarkValue.Valid = u.MarkValue.Valid
				}

				if u.RiskValue.Set {
					v.RiskValue.Value = u.RiskValue.Value
					v.RiskValue.Valid = u.RiskValue.Valid
				}

				if u.HomeNotional.Set {
					v.HomeNotional.Value = u.HomeNotional.Value
					v.HomeNotional.Valid = u.HomeNotional.Valid
				}

				if u.ForeignNotional.Set {
					v.ForeignNotional.Value = u.ForeignNotional.Value
					v.ForeignNotional.Valid = u.ForeignNotional.Valid
				}

				if u.PosState.Set {
					v.PosState.Value = u.PosState.Value
					v.PosState.Valid = u.PosState.Valid
				}

				if u.PosCost.Set {
					v.PosCost.Value = u.PosCost.Value
					v.PosCost.Valid = u.PosCost.Valid
				}

				if u.PosCost2.Set {
					v.PosCost2.Value = u.PosCost2.Value
					v.PosCost2.Valid = u.PosCost2.Valid
				}

				if u.PosCross.Set {
					v.PosCross.Value = u.PosCross.Value
					v.PosCross.Valid = u.PosCross.Valid
				}

				if u.PosInit.Set {
					v.PosInit.Value = u.PosInit.Value
					v.PosInit.Valid = u.PosInit.Valid
				}

				if u.PosComm.Set {
					v.PosComm.Value = u.PosComm.Value
					v.PosComm.Valid = u.PosComm.Valid
				}

				if u.PosLoss.Set {
					v.PosLoss.Value = u.PosLoss.Value
					v.PosLoss.Valid = u.PosLoss.Valid
				}

				if u.PosMargin.Set {
					v.PosMargin.Value = u.PosMargin.Value
					v.PosMargin.Valid = u.PosMargin.Valid
				}

				if u.PosMaint.Set {
					v.PosMaint.Value = u.PosMaint.Value
					v.PosMaint.Valid = u.PosMaint.Valid
				}

				if u.PosAllowance.Set {
					v.PosAllowance.Value = u.PosAllowance.Value
					v.PosAllowance.Valid = u.PosAllowance.Valid
				}

				if u.TaxableMargin.Set {
					v.TaxableMargin.Value = u.TaxableMargin.Value
					v.TaxableMargin.Valid = u.TaxableMargin.Valid
				}

				if u.InitMargin.Set {
					v.InitMargin.Value = u.InitMargin.Value
					v.InitMargin.Valid = u.InitMargin.Valid
				}

				if u.MaintMargin.Set {
					v.MaintMargin.Value = u.MaintMargin.Value
					v.MaintMargin.Valid = u.MaintMargin.Valid
				}

				if u.SessionMargin.Set {
					v.SessionMargin.Value = u.SessionMargin.Value
					v.SessionMargin.Valid = u.SessionMargin.Valid
				}

				if u.TargetExcessMargin.Set {
					v.TargetExcessMargin.Value = u.TargetExcessMargin.Value
					v.TargetExcessMargin.Valid = u.TargetExcessMargin.Valid
				}

				if u.VarMargin.Set {
					v.VarMargin.Value = u.VarMargin.Value
					v.VarMargin.Valid = u.VarMargin.Valid
				}

				if u.RealisedGrossPnl.Set {
					v.RealisedGrossPnl.Value = u.RealisedGrossPnl.Value
					v.RealisedGrossPnl.Valid = u.RealisedGrossPnl.Valid
				}

				if u.RealisedTax.Set {
					v.RealisedTax.Value = u.RealisedTax.Value
					v.RealisedTax.Valid = u.RealisedTax.Valid
				}

				if u.RealisedPnl.Set {
					v.RealisedPnl.Value = u.RealisedPnl.Value
					v.RealisedPnl.Valid = u.RealisedPnl.Valid
				}

				if u.UnrealisedGrossPnl.Set {
					v.UnrealisedGrossPnl.Value = u.UnrealisedGrossPnl.Value
					v.UnrealisedGrossPnl.Valid = u.UnrealisedGrossPnl.Valid
				}

				if u.LongBankrupt.Set {
					v.LongBankrupt.Value = u.LongBankrupt.Value
					v.LongBankrupt.Valid = u.LongBankrupt.Valid
				}

				if u.ShortBankrupt.Set {
					v.ShortBankrupt.Value = u.ShortBankrupt.Value
					v.ShortBankrupt.Valid = u.ShortBankrupt.Valid
				}

				if u.TaxBase.Set {
					v.TaxBase.Value = u.TaxBase.Value
					v.TaxBase.Valid = u.TaxBase.Valid
				}

				if u.IndicativeTaxRate.Set {
					v.IndicativeTaxRate.Value = u.IndicativeTaxRate.Value
					v.IndicativeTaxRate.Valid = u.IndicativeTaxRate.Valid
				}

				if u.IndicativeTax.Set {
					v.IndicativeTax.Value = u.IndicativeTax.Value
					v.IndicativeTax.Valid = u.IndicativeTax.Valid
				}

				if u.UnrealisedTax.Set {
					v.UnrealisedTax.Value = u.UnrealisedTax.Value
					v.UnrealisedTax.Valid = u.UnrealisedTax.Valid
				}

				if u.UnrealisedPnl.Set {
					v.UnrealisedPnl.Value = u.UnrealisedPnl.Value
					v.UnrealisedPnl.Valid = u.UnrealisedPnl.Valid
				}

				if u.UnrealisedPnlPcnt.Set {
					v.UnrealisedPnlPcnt.Value = u.UnrealisedPnlPcnt.Value
					v.UnrealisedPnlPcnt.Valid = u.UnrealisedPnlPcnt.Valid
				}

				if u.UnrealisedRoePcnt.Set {
					v.UnrealisedRoePcnt.Value = u.UnrealisedRoePcnt.Value
					v.UnrealisedRoePcnt.Valid = u.UnrealisedRoePcnt.Valid
				}

				if u.SimpleQty.Set {
					v.SimpleQty.Value = u.SimpleQty.Value
					v.SimpleQty.Valid = u.SimpleQty.Valid
				}

				if u.SimpleCost.Set {
					v.SimpleCost.Value = u.SimpleCost.Value
					v.SimpleCost.Valid = u.SimpleCost.Valid
				}

				if u.SimpleValue.Set {
					v.SimpleValue.Value = u.SimpleValue.Value
					v.SimpleValue.Valid = u.SimpleValue.Valid
				}

				if u.SimplePnl.Set {
					v.SimplePnl.Value = u.SimplePnl.Value
					v.SimplePnl.Valid = u.SimplePnl.Valid
				}

				if u.SimplePnlPcnt.Set {
					v.SimplePnlPcnt.Value = u.SimplePnlPcnt.Value
					v.SimplePnlPcnt.Valid = u.SimplePnlPcnt.Valid
				}

				if u.AvgCostPrice.Set {
					v.AvgCostPrice.Value = u.AvgCostPrice.Value
					v.AvgCostPrice.Valid = u.AvgCostPrice.Valid
				}

				if u.AvgEntryPrice.Set {
					v.AvgEntryPrice.Value = u.AvgEntryPrice.Value
					v.AvgEntryPrice.Valid = u.AvgEntryPrice.Valid
				}

				if u.BreakEvenPrice.Set {
					v.BreakEvenPrice.Value = u.BreakEvenPrice.Value
					v.BreakEvenPrice.Valid = u.BreakEvenPrice.Valid
				}

				if u.MarginCallPrice.Set {
					v.MarginCallPrice.Value = u.MarginCallPrice.Value
					v.MarginCallPrice.Valid = u.MarginCallPrice.Valid
				}

				if u.LiquidationPrice.Set {
					v.LiquidationPrice.Value = u.LiquidationPrice.Value
					v.LiquidationPrice.Valid = u.LiquidationPrice.Valid
				}

				if u.BankruptPrice.Set {
					v.BankruptPrice.Value = u.BankruptPrice.Value
					v.BankruptPrice.Valid = u.BankruptPrice.Valid
				}

				if u.Timestamp.Set {
					v.Timestamp.Value = u.Timestamp.Value
					v.Timestamp.Valid = u.Timestamp.Valid
				}

				if u.LastPrice.Set {
					v.LastPrice.Value = u.LastPrice.Value
					v.LastPrice.Valid = u.LastPrice.Valid
				}

				if u.LastValue.Set {
					v.LastValue.Value = u.LastValue.Value
					v.LastValue.Valid = u.LastValue.Valid
				}

				later := (*positions)[i+1:]
				*positions = append((*positions)[:i], v)
				*positions = append(*positions, later...)

			}
		}
	}
}

func (positions *PositionSlice) PositionDelete(deletes *[]PositionResponseData) {
	for _, u := range *deletes {
		for i, v := range *positions {
			if u.Symbol.Value == v.Symbol.Value {
				*positions = append((*positions)[:i], (*positions)[i+1:]...)
			}
		}
	}
}
