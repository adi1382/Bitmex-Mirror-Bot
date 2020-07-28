package websocket

import "github.com/adi1382/Bitmex-Mirror-Bot/swagger"

type MarginSlice []swagger.Margin

type MarginResponseData struct {
	Account            JSONFloat64 `json:"account"`
	Currency           JSONString  `json:"currency"`
	RiskLimit          JSONFloat64 `json:"riskLimit,omitempty"`
	PrevState          JSONString  `json:"prevState,omitempty"`
	State              JSONString  `json:"state,omitempty"`
	Action             JSONString  `json:"action,omitempty"`
	Amount             JSONFloat64 `json:"amount,omitempty"`
	PendingCredit      JSONFloat64 `json:"pendingCredit,omitempty"`
	PendingDebit       JSONFloat64 `json:"pendingDebit,omitempty"`
	ConfirmedDebit     JSONFloat64 `json:"confirmedDebit,omitempty"`
	PrevRealisedPnl    JSONFloat64 `json:"prevRealisedPnl,omitempty"`
	PrevUnrealisedPnl  JSONFloat64 `json:"prevUnrealisedPnl,omitempty"`
	GrossComm          JSONFloat64 `json:"grossComm,omitempty"`
	GrossOpenCost      JSONFloat64 `json:"grossOpenCost,omitempty"`
	GrossOpenPremium   JSONFloat64 `json:"grossOpenPremium,omitempty"`
	GrossExecCost      JSONFloat64 `json:"grossExecCost,omitempty"`
	GrossMarkValue     JSONFloat64 `json:"grossMarkValue,omitempty"`
	RiskValue          JSONFloat64 `json:"riskValue,omitempty"`
	TaxableMargin      JSONFloat64 `json:"taxableMargin,omitempty"`
	InitMargin         JSONFloat64 `json:"initMargin,omitempty"`
	MaintMargin        JSONFloat64 `json:"maintMargin,omitempty"`
	SessionMargin      JSONFloat64 `json:"sessionMargin,omitempty"`
	TargetExcessMargin JSONFloat64 `json:"targetExcessMargin,omitempty"`
	VarMargin          JSONFloat64 `json:"varMargin,omitempty"`
	RealisedPnl        JSONFloat64 `json:"realisedPnl,omitempty"`
	UnrealisedPnl      JSONFloat64 `json:"unrealisedPnl,omitempty"`
	IndicativeTax      JSONFloat64 `json:"indicativeTax,omitempty"`
	UnrealisedProfit   JSONFloat64 `json:"unrealisedProfit,omitempty"`
	SyntheticMargin    JSONFloat64 `json:"syntheticMargin,omitempty"`
	WalletBalance      JSONFloat64 `json:"walletBalance,omitempty"`
	MarginBalance      JSONFloat64 `json:"marginBalance,omitempty"`
	MarginBalancePcnt  JSONFloat64 `json:"marginBalancePcnt,omitempty"`
	MarginLeverage     JSONFloat64 `json:"marginLeverage,omitempty"`
	MarginUsedPcnt     JSONFloat64 `json:"marginUsedPcnt,omitempty"`
	ExcessMargin       JSONFloat64 `json:"excessMargin,omitempty"`
	ExcessMarginPcnt   JSONFloat64 `json:"excessMarginPcnt,omitempty"`
	AvailableMargin    JSONFloat64 `json:"availableMargin,omitempty"`
	WithdrawableMargin JSONFloat64 `json:"withdrawableMargin,omitempty"`
	Timestamp          JSONTime    `json:"timestamp,omitempty"`
	GrossLastValue     JSONFloat64 `json:"grossLastValue,omitempty"`
	Commission         JSONFloat64 `json:"commission,omitempty"`
}

func (margin *MarginSlice) MarginPartial(inserts *[]MarginResponseData) {
	InfoLogger.Println("Margin Partials Processing")
	*margin = nil
	for v := range *inserts {
		var insertMargin swagger.Margin

		insertMargin.Account.Value = (*inserts)[v].Account.Value
		insertMargin.Account.Valid = (*inserts)[v].Account.Valid

		insertMargin.Currency.Value = (*inserts)[v].Currency.Value
		insertMargin.Currency.Valid = (*inserts)[v].Currency.Valid

		insertMargin.RiskLimit.Value = (*inserts)[v].RiskLimit.Value
		insertMargin.RiskLimit.Valid = (*inserts)[v].RiskLimit.Valid

		insertMargin.PrevState.Value = (*inserts)[v].PrevState.Value
		insertMargin.PrevState.Valid = (*inserts)[v].PrevState.Valid

		insertMargin.State.Value = (*inserts)[v].State.Value
		insertMargin.State.Valid = (*inserts)[v].State.Valid

		insertMargin.Action.Value = (*inserts)[v].Action.Value
		insertMargin.Action.Valid = (*inserts)[v].Action.Valid

		insertMargin.Amount.Value = (*inserts)[v].Amount.Value
		insertMargin.Amount.Valid = (*inserts)[v].Amount.Valid

		insertMargin.PendingCredit.Value = (*inserts)[v].PendingCredit.Value
		insertMargin.PendingCredit.Valid = (*inserts)[v].PendingCredit.Valid

		insertMargin.PendingDebit.Value = (*inserts)[v].PendingDebit.Value
		insertMargin.PendingDebit.Valid = (*inserts)[v].PendingDebit.Valid

		insertMargin.ConfirmedDebit.Value = (*inserts)[v].ConfirmedDebit.Value
		insertMargin.ConfirmedDebit.Valid = (*inserts)[v].ConfirmedDebit.Valid

		insertMargin.PrevRealisedPnl.Value = (*inserts)[v].PrevRealisedPnl.Value
		insertMargin.PrevRealisedPnl.Valid = (*inserts)[v].PrevRealisedPnl.Valid

		insertMargin.PrevUnrealisedPnl.Value = (*inserts)[v].PrevUnrealisedPnl.Value
		insertMargin.PrevUnrealisedPnl.Valid = (*inserts)[v].PrevUnrealisedPnl.Valid

		insertMargin.GrossComm.Value = (*inserts)[v].GrossComm.Value
		insertMargin.GrossComm.Valid = (*inserts)[v].GrossComm.Valid

		insertMargin.GrossOpenCost.Value = (*inserts)[v].GrossOpenCost.Value
		insertMargin.GrossOpenCost.Valid = (*inserts)[v].GrossOpenCost.Valid

		insertMargin.GrossOpenPremium.Value = (*inserts)[v].GrossOpenPremium.Value
		insertMargin.GrossOpenPremium.Valid = (*inserts)[v].GrossOpenPremium.Valid

		insertMargin.GrossExecCost.Value = (*inserts)[v].GrossExecCost.Value
		insertMargin.GrossExecCost.Valid = (*inserts)[v].GrossExecCost.Valid

		insertMargin.GrossMarkValue.Value = (*inserts)[v].GrossMarkValue.Value
		insertMargin.GrossMarkValue.Valid = (*inserts)[v].GrossMarkValue.Valid

		insertMargin.RiskLimit.Value = (*inserts)[v].RiskLimit.Value
		insertMargin.RiskLimit.Valid = (*inserts)[v].RiskLimit.Valid

		insertMargin.TaxableMargin.Value = (*inserts)[v].TaxableMargin.Value
		insertMargin.TaxableMargin.Valid = (*inserts)[v].TaxableMargin.Valid

		insertMargin.InitMargin.Value = (*inserts)[v].InitMargin.Value
		insertMargin.InitMargin.Valid = (*inserts)[v].InitMargin.Valid

		insertMargin.MaintMargin.Value = (*inserts)[v].MaintMargin.Value
		insertMargin.MaintMargin.Valid = (*inserts)[v].MaintMargin.Valid

		insertMargin.SessionMargin.Value = (*inserts)[v].SessionMargin.Value
		insertMargin.SessionMargin.Valid = (*inserts)[v].SessionMargin.Valid

		insertMargin.TargetExcessMargin.Value = (*inserts)[v].TargetExcessMargin.Value
		insertMargin.TargetExcessMargin.Valid = (*inserts)[v].TargetExcessMargin.Valid

		insertMargin.VarMargin.Value = (*inserts)[v].VarMargin.Value
		insertMargin.VarMargin.Valid = (*inserts)[v].VarMargin.Valid

		insertMargin.RealisedPnl.Value = (*inserts)[v].RealisedPnl.Value
		insertMargin.RealisedPnl.Valid = (*inserts)[v].RealisedPnl.Valid

		insertMargin.UnrealisedPnl.Value = (*inserts)[v].UnrealisedPnl.Value
		insertMargin.UnrealisedPnl.Valid = (*inserts)[v].UnrealisedPnl.Valid

		insertMargin.UnrealisedPnl.Value = (*inserts)[v].UnrealisedPnl.Value
		insertMargin.UnrealisedPnl.Valid = (*inserts)[v].UnrealisedPnl.Valid

		insertMargin.UnrealisedProfit.Value = (*inserts)[v].UnrealisedProfit.Value
		insertMargin.UnrealisedProfit.Valid = (*inserts)[v].UnrealisedProfit.Valid

		insertMargin.SyntheticMargin.Value = (*inserts)[v].SyntheticMargin.Value
		insertMargin.SyntheticMargin.Valid = (*inserts)[v].SyntheticMargin.Valid

		insertMargin.WalletBalance.Value = (*inserts)[v].WalletBalance.Value
		insertMargin.WalletBalance.Valid = (*inserts)[v].WalletBalance.Valid

		insertMargin.MarginBalance.Value = (*inserts)[v].MarginBalance.Value
		insertMargin.MarginBalance.Valid = (*inserts)[v].MarginBalance.Valid

		insertMargin.MarginBalancePcnt.Value = (*inserts)[v].MarginBalancePcnt.Value
		insertMargin.MarginBalancePcnt.Valid = (*inserts)[v].MarginBalancePcnt.Valid

		insertMargin.MarginLeverage.Value = (*inserts)[v].MarginLeverage.Value
		insertMargin.MarginLeverage.Valid = (*inserts)[v].MarginLeverage.Valid

		insertMargin.MarginUsedPcnt.Value = (*inserts)[v].MarginUsedPcnt.Value
		insertMargin.MarginUsedPcnt.Valid = (*inserts)[v].MarginUsedPcnt.Valid

		insertMargin.ExcessMargin.Value = (*inserts)[v].ExcessMargin.Value
		insertMargin.ExcessMargin.Valid = (*inserts)[v].ExcessMargin.Valid

		insertMargin.ExcessMarginPcnt.Value = (*inserts)[v].ExcessMarginPcnt.Value
		insertMargin.ExcessMarginPcnt.Valid = (*inserts)[v].ExcessMarginPcnt.Valid

		insertMargin.AvailableMargin.Value = (*inserts)[v].AvailableMargin.Value
		insertMargin.AvailableMargin.Valid = (*inserts)[v].AvailableMargin.Valid

		insertMargin.WithdrawableMargin.Value = (*inserts)[v].WithdrawableMargin.Value
		insertMargin.WithdrawableMargin.Valid = (*inserts)[v].WithdrawableMargin.Valid

		insertMargin.Timestamp.Value = (*inserts)[v].Timestamp.Value
		insertMargin.Timestamp.Valid = (*inserts)[v].Timestamp.Valid

		insertMargin.GrossLastValue.Value = (*inserts)[v].GrossLastValue.Value
		insertMargin.GrossLastValue.Valid = (*inserts)[v].GrossLastValue.Valid

		*margin = append(*margin, insertMargin)
	}
	InfoLogger.Println("Margin Partials Processed")
}

func (margin *MarginSlice) MarginInsert(inserts *[]MarginResponseData) {
	InfoLogger.Println("Margin Inserts Processing")
	for v := range *inserts {
		var insertMargin swagger.Margin

		insertMargin.Account.Value = (*inserts)[v].Account.Value
		insertMargin.Account.Valid = (*inserts)[v].Account.Valid

		insertMargin.Currency.Value = (*inserts)[v].Currency.Value
		insertMargin.Currency.Valid = (*inserts)[v].Currency.Valid

		insertMargin.RiskLimit.Value = (*inserts)[v].RiskLimit.Value
		insertMargin.RiskLimit.Valid = (*inserts)[v].RiskLimit.Valid

		insertMargin.PrevState.Value = (*inserts)[v].PrevState.Value
		insertMargin.PrevState.Valid = (*inserts)[v].PrevState.Valid

		insertMargin.State.Value = (*inserts)[v].State.Value
		insertMargin.State.Valid = (*inserts)[v].State.Valid

		insertMargin.Action.Value = (*inserts)[v].Action.Value
		insertMargin.Action.Valid = (*inserts)[v].Action.Valid

		insertMargin.Amount.Value = (*inserts)[v].Amount.Value
		insertMargin.Amount.Valid = (*inserts)[v].Amount.Valid

		insertMargin.PendingCredit.Value = (*inserts)[v].PendingCredit.Value
		insertMargin.PendingCredit.Valid = (*inserts)[v].PendingCredit.Valid

		insertMargin.PendingDebit.Value = (*inserts)[v].PendingDebit.Value
		insertMargin.PendingDebit.Valid = (*inserts)[v].PendingDebit.Valid

		insertMargin.ConfirmedDebit.Value = (*inserts)[v].ConfirmedDebit.Value
		insertMargin.ConfirmedDebit.Valid = (*inserts)[v].ConfirmedDebit.Valid

		insertMargin.PrevRealisedPnl.Value = (*inserts)[v].PrevRealisedPnl.Value
		insertMargin.PrevRealisedPnl.Valid = (*inserts)[v].PrevRealisedPnl.Valid

		insertMargin.PrevUnrealisedPnl.Value = (*inserts)[v].PrevUnrealisedPnl.Value
		insertMargin.PrevUnrealisedPnl.Valid = (*inserts)[v].PrevUnrealisedPnl.Valid

		insertMargin.GrossComm.Value = (*inserts)[v].GrossComm.Value
		insertMargin.GrossComm.Valid = (*inserts)[v].GrossComm.Valid

		insertMargin.GrossOpenCost.Value = (*inserts)[v].GrossOpenCost.Value
		insertMargin.GrossOpenCost.Valid = (*inserts)[v].GrossOpenCost.Valid

		insertMargin.GrossOpenPremium.Value = (*inserts)[v].GrossOpenPremium.Value
		insertMargin.GrossOpenPremium.Valid = (*inserts)[v].GrossOpenPremium.Valid

		insertMargin.GrossExecCost.Value = (*inserts)[v].GrossExecCost.Value
		insertMargin.GrossExecCost.Valid = (*inserts)[v].GrossExecCost.Valid

		insertMargin.GrossMarkValue.Value = (*inserts)[v].GrossMarkValue.Value
		insertMargin.GrossMarkValue.Valid = (*inserts)[v].GrossMarkValue.Valid

		insertMargin.RiskLimit.Value = (*inserts)[v].RiskLimit.Value
		insertMargin.RiskLimit.Valid = (*inserts)[v].RiskLimit.Valid

		insertMargin.TaxableMargin.Value = (*inserts)[v].TaxableMargin.Value
		insertMargin.TaxableMargin.Valid = (*inserts)[v].TaxableMargin.Valid

		insertMargin.InitMargin.Value = (*inserts)[v].InitMargin.Value
		insertMargin.InitMargin.Valid = (*inserts)[v].InitMargin.Valid

		insertMargin.MaintMargin.Value = (*inserts)[v].MaintMargin.Value
		insertMargin.MaintMargin.Valid = (*inserts)[v].MaintMargin.Valid

		insertMargin.SessionMargin.Value = (*inserts)[v].SessionMargin.Value
		insertMargin.SessionMargin.Valid = (*inserts)[v].SessionMargin.Valid

		insertMargin.TargetExcessMargin.Value = (*inserts)[v].TargetExcessMargin.Value
		insertMargin.TargetExcessMargin.Valid = (*inserts)[v].TargetExcessMargin.Valid

		insertMargin.VarMargin.Value = (*inserts)[v].VarMargin.Value
		insertMargin.VarMargin.Valid = (*inserts)[v].VarMargin.Valid

		insertMargin.RealisedPnl.Value = (*inserts)[v].RealisedPnl.Value
		insertMargin.RealisedPnl.Valid = (*inserts)[v].RealisedPnl.Valid

		insertMargin.UnrealisedPnl.Value = (*inserts)[v].UnrealisedPnl.Value
		insertMargin.UnrealisedPnl.Valid = (*inserts)[v].UnrealisedPnl.Valid

		insertMargin.UnrealisedPnl.Value = (*inserts)[v].UnrealisedPnl.Value
		insertMargin.UnrealisedPnl.Valid = (*inserts)[v].UnrealisedPnl.Valid

		insertMargin.UnrealisedProfit.Value = (*inserts)[v].UnrealisedProfit.Value
		insertMargin.UnrealisedProfit.Valid = (*inserts)[v].UnrealisedProfit.Valid

		insertMargin.SyntheticMargin.Value = (*inserts)[v].SyntheticMargin.Value
		insertMargin.SyntheticMargin.Valid = (*inserts)[v].SyntheticMargin.Valid

		insertMargin.WalletBalance.Value = (*inserts)[v].WalletBalance.Value
		insertMargin.WalletBalance.Valid = (*inserts)[v].WalletBalance.Valid

		insertMargin.MarginBalance.Value = (*inserts)[v].MarginBalance.Value
		insertMargin.MarginBalance.Valid = (*inserts)[v].MarginBalance.Valid

		insertMargin.MarginBalancePcnt.Value = (*inserts)[v].MarginBalancePcnt.Value
		insertMargin.MarginBalancePcnt.Valid = (*inserts)[v].MarginBalancePcnt.Valid

		insertMargin.MarginLeverage.Value = (*inserts)[v].MarginLeverage.Value
		insertMargin.MarginLeverage.Valid = (*inserts)[v].MarginLeverage.Valid

		insertMargin.MarginUsedPcnt.Value = (*inserts)[v].MarginUsedPcnt.Value
		insertMargin.MarginUsedPcnt.Valid = (*inserts)[v].MarginUsedPcnt.Valid

		insertMargin.ExcessMargin.Value = (*inserts)[v].ExcessMargin.Value
		insertMargin.ExcessMargin.Valid = (*inserts)[v].ExcessMargin.Valid

		insertMargin.ExcessMarginPcnt.Value = (*inserts)[v].ExcessMarginPcnt.Value
		insertMargin.ExcessMarginPcnt.Valid = (*inserts)[v].ExcessMarginPcnt.Valid

		insertMargin.AvailableMargin.Value = (*inserts)[v].AvailableMargin.Value
		insertMargin.AvailableMargin.Valid = (*inserts)[v].AvailableMargin.Valid

		insertMargin.WithdrawableMargin.Value = (*inserts)[v].WithdrawableMargin.Value
		insertMargin.WithdrawableMargin.Valid = (*inserts)[v].WithdrawableMargin.Valid

		insertMargin.Timestamp.Value = (*inserts)[v].Timestamp.Value
		insertMargin.Timestamp.Valid = (*inserts)[v].Timestamp.Valid

		insertMargin.GrossLastValue.Value = (*inserts)[v].GrossLastValue.Value
		insertMargin.GrossLastValue.Valid = (*inserts)[v].GrossLastValue.Valid

		*margin = append(*margin, insertMargin)
	}
	InfoLogger.Println("Margin Inserts Processed")
}

func (margin *MarginSlice) MarginUpdate(updates *[]MarginResponseData) {
	InfoLogger.Println("Margin Updates Processing")
	for u := range *updates {
		for i := range *margin {

			if (*updates)[u].Account.Value == (*margin)[i].Account.Value {
				if (*updates)[u].Account.Set {
					(*margin)[i].Account.Value = (*updates)[u].Account.Value
					(*margin)[i].Account.Valid = (*updates)[u].Account.Valid
				}

				if (*updates)[u].Currency.Set {
					(*margin)[i].Currency.Value = (*updates)[u].Currency.Value
					(*margin)[i].Currency.Valid = (*updates)[u].Currency.Valid
				}

				if (*updates)[u].RiskLimit.Set {
					(*margin)[i].RiskLimit.Value = (*updates)[u].RiskLimit.Value
					(*margin)[i].RiskLimit.Valid = (*updates)[u].RiskLimit.Valid
				}

				if (*updates)[u].PrevState.Set {
					(*margin)[i].PrevState.Value = (*updates)[u].PrevState.Value
					(*margin)[i].PrevState.Valid = (*updates)[u].PrevState.Valid
				}

				if (*updates)[u].State.Set {
					(*margin)[i].State.Value = (*updates)[u].State.Value
					(*margin)[i].State.Valid = (*updates)[u].State.Valid
				}

				if (*updates)[u].Action.Set {
					(*margin)[i].Action.Value = (*updates)[u].Action.Value
					(*margin)[i].Action.Valid = (*updates)[u].Action.Valid
				}

				if (*updates)[u].Amount.Set {
					(*margin)[i].Amount.Value = (*updates)[u].Amount.Value
					(*margin)[i].Amount.Valid = (*updates)[u].Amount.Valid
				}

				if (*updates)[u].PendingCredit.Set {
					(*margin)[i].PendingCredit.Value = (*updates)[u].PendingCredit.Value
					(*margin)[i].PendingCredit.Valid = (*updates)[u].PendingCredit.Valid
				}

				if (*updates)[u].PendingDebit.Set {
					(*margin)[i].PendingDebit.Value = (*updates)[u].PendingDebit.Value
					(*margin)[i].PendingDebit.Valid = (*updates)[u].PendingDebit.Valid
				}

				if (*updates)[u].ConfirmedDebit.Set {
					(*margin)[i].ConfirmedDebit.Value = (*updates)[u].ConfirmedDebit.Value
					(*margin)[i].ConfirmedDebit.Valid = (*updates)[u].ConfirmedDebit.Valid
				}

				if (*updates)[u].PrevRealisedPnl.Set {
					(*margin)[i].PrevRealisedPnl.Value = (*updates)[u].PrevRealisedPnl.Value
					(*margin)[i].PrevRealisedPnl.Valid = (*updates)[u].PrevRealisedPnl.Valid
				}

				if (*updates)[u].PrevUnrealisedPnl.Set {
					(*margin)[i].PrevUnrealisedPnl.Value = (*updates)[u].PrevUnrealisedPnl.Value
					(*margin)[i].PrevUnrealisedPnl.Valid = (*updates)[u].PrevUnrealisedPnl.Valid
				}

				if (*updates)[u].GrossComm.Set {
					(*margin)[i].GrossComm.Value = (*updates)[u].GrossComm.Value
					(*margin)[i].GrossComm.Valid = (*updates)[u].GrossComm.Valid
				}

				if (*updates)[u].GrossOpenCost.Set {
					(*margin)[i].GrossOpenCost.Value = (*updates)[u].GrossOpenCost.Value
					(*margin)[i].GrossOpenCost.Valid = (*updates)[u].GrossOpenCost.Valid
				}

				if (*updates)[u].GrossOpenPremium.Set {
					(*margin)[i].GrossOpenPremium.Value = (*updates)[u].GrossOpenPremium.Value
					(*margin)[i].GrossOpenPremium.Valid = (*updates)[u].GrossOpenPremium.Valid
				}

				if (*updates)[u].GrossExecCost.Set {
					(*margin)[i].GrossExecCost.Value = (*updates)[u].GrossExecCost.Value
					(*margin)[i].GrossExecCost.Valid = (*updates)[u].GrossExecCost.Valid
				}

				if (*updates)[u].GrossMarkValue.Set {
					(*margin)[i].GrossMarkValue.Value = (*updates)[u].GrossMarkValue.Value
					(*margin)[i].GrossMarkValue.Valid = (*updates)[u].GrossMarkValue.Valid
				}

				if (*updates)[u].RiskValue.Set {
					(*margin)[i].RiskValue.Value = (*updates)[u].RiskValue.Value
					(*margin)[i].RiskValue.Valid = (*updates)[u].RiskValue.Valid
				}

				if (*updates)[u].TaxableMargin.Set {
					(*margin)[i].TaxableMargin.Value = (*updates)[u].TaxableMargin.Value
					(*margin)[i].TaxableMargin.Valid = (*updates)[u].TaxableMargin.Valid
				}

				if (*updates)[u].InitMargin.Set {
					(*margin)[i].InitMargin.Value = (*updates)[u].InitMargin.Value
					(*margin)[i].InitMargin.Valid = (*updates)[u].InitMargin.Valid
				}

				if (*updates)[u].MaintMargin.Set {
					(*margin)[i].MaintMargin.Value = (*updates)[u].MaintMargin.Value
					(*margin)[i].MaintMargin.Valid = (*updates)[u].MaintMargin.Valid
				}

				if (*updates)[u].SessionMargin.Set {
					(*margin)[i].SessionMargin.Value = (*updates)[u].SessionMargin.Value
					(*margin)[i].SessionMargin.Valid = (*updates)[u].SessionMargin.Valid
				}

				if (*updates)[u].TargetExcessMargin.Set {
					(*margin)[i].TargetExcessMargin.Value = (*updates)[u].TargetExcessMargin.Value
					(*margin)[i].TargetExcessMargin.Valid = (*updates)[u].TargetExcessMargin.Valid
				}

				if (*updates)[u].VarMargin.Set {
					(*margin)[i].VarMargin.Value = (*updates)[u].VarMargin.Value
					(*margin)[i].VarMargin.Valid = (*updates)[u].VarMargin.Valid
				}

				if (*updates)[u].RealisedPnl.Set {
					(*margin)[i].RealisedPnl.Value = (*updates)[u].RealisedPnl.Value
					(*margin)[i].RealisedPnl.Valid = (*updates)[u].RealisedPnl.Valid
				}

				if (*updates)[u].UnrealisedPnl.Set {
					(*margin)[i].UnrealisedPnl.Value = (*updates)[u].UnrealisedPnl.Value
					(*margin)[i].UnrealisedPnl.Valid = (*updates)[u].UnrealisedPnl.Valid
				}

				if (*updates)[u].IndicativeTax.Set {
					(*margin)[i].IndicativeTax.Value = (*updates)[u].IndicativeTax.Value
					(*margin)[i].IndicativeTax.Valid = (*updates)[u].IndicativeTax.Valid
				}

				if (*updates)[u].UnrealisedProfit.Set {
					(*margin)[i].UnrealisedProfit.Value = (*updates)[u].UnrealisedProfit.Value
					(*margin)[i].UnrealisedProfit.Valid = (*updates)[u].UnrealisedProfit.Valid
				}

				if (*updates)[u].SyntheticMargin.Set {
					(*margin)[i].SyntheticMargin.Value = (*updates)[u].SyntheticMargin.Value
					(*margin)[i].SyntheticMargin.Valid = (*updates)[u].SyntheticMargin.Valid
				}

				if (*updates)[u].WalletBalance.Set {
					(*margin)[i].WalletBalance.Value = (*updates)[u].WalletBalance.Value
					(*margin)[i].WalletBalance.Valid = (*updates)[u].WalletBalance.Valid
				}

				if (*updates)[u].MarginBalance.Set {
					(*margin)[i].MarginBalance.Value = (*updates)[u].MarginBalance.Value
					(*margin)[i].MarginBalance.Valid = (*updates)[u].MarginBalance.Valid
				}

				if (*updates)[u].MarginBalancePcnt.Set {
					(*margin)[i].MarginBalancePcnt.Value = (*updates)[u].MarginBalancePcnt.Value
					(*margin)[i].MarginBalancePcnt.Valid = (*updates)[u].MarginBalancePcnt.Valid
				}

				if (*updates)[u].MarginLeverage.Set {
					(*margin)[i].MarginLeverage.Value = (*updates)[u].MarginLeverage.Value
					(*margin)[i].MarginLeverage.Valid = (*updates)[u].MarginLeverage.Valid
				}

				if (*updates)[u].MarginUsedPcnt.Set {
					(*margin)[i].MarginUsedPcnt.Value = (*updates)[u].MarginUsedPcnt.Value
					(*margin)[i].MarginUsedPcnt.Valid = (*updates)[u].MarginUsedPcnt.Valid
				}

				if (*updates)[u].ExcessMargin.Set {
					(*margin)[i].ExcessMargin.Value = (*updates)[u].ExcessMargin.Value
					(*margin)[i].ExcessMargin.Valid = (*updates)[u].ExcessMargin.Valid
				}

				if (*updates)[u].ExcessMarginPcnt.Set {
					(*margin)[i].ExcessMarginPcnt.Value = (*updates)[u].ExcessMarginPcnt.Value
					(*margin)[i].ExcessMarginPcnt.Valid = (*updates)[u].ExcessMarginPcnt.Valid
				}

				if (*updates)[u].AvailableMargin.Set {
					(*margin)[i].AvailableMargin.Value = (*updates)[u].AvailableMargin.Value
					(*margin)[i].AvailableMargin.Valid = (*updates)[u].AvailableMargin.Valid
				}

				if (*updates)[u].WithdrawableMargin.Set {
					(*margin)[i].WithdrawableMargin.Value = (*updates)[u].WithdrawableMargin.Value
					(*margin)[i].WithdrawableMargin.Valid = (*updates)[u].WithdrawableMargin.Valid
				}

				if (*updates)[u].Timestamp.Set {
					(*margin)[i].Timestamp.Value = (*updates)[u].Timestamp.Value
					(*margin)[i].Timestamp.Valid = (*updates)[u].Timestamp.Valid
				}

				if (*updates)[u].GrossLastValue.Set {
					(*margin)[i].GrossLastValue.Value = (*updates)[u].GrossLastValue.Value
					(*margin)[i].GrossLastValue.Valid = (*updates)[u].GrossLastValue.Valid
				}

				if (*updates)[u].Commission.Set {
					(*margin)[i].Commission.Value = (*updates)[u].Commission.Value
					(*margin)[i].Commission.Valid = (*updates)[u].Commission.Valid
				}

				later := (*margin)[i+1:]
				*margin = append((*margin)[:i], (*margin)[i])
				*margin = append(*margin, later...)
			}
		}
	}
	InfoLogger.Println("Margin Updates Processed")
}

func (margin *MarginSlice) MarginDelete(deletes *[]MarginResponseData) {
	InfoLogger.Println("Margin Deletes Processing")
	for u := range *deletes {
		for i := range *margin {
			if (*deletes)[u].Account.Value == (*margin)[i].Account.Value {
				*margin = append((*margin)[:i], (*margin)[i+1:]...)
			}
		}
	}
	InfoLogger.Println("Margin Deletes Processed")
}
