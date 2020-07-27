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
	for _, v := range *inserts {
		var insertMargin swagger.Margin

		insertMargin.Account.Value = v.Account.Value
		insertMargin.Account.Valid = v.Account.Valid

		insertMargin.Currency.Value = v.Currency.Value
		insertMargin.Currency.Valid = v.Currency.Valid

		insertMargin.RiskLimit.Value = v.RiskLimit.Value
		insertMargin.RiskLimit.Valid = v.RiskLimit.Valid

		insertMargin.PrevState.Value = v.PrevState.Value
		insertMargin.PrevState.Valid = v.PrevState.Valid

		insertMargin.State.Value = v.State.Value
		insertMargin.State.Valid = v.State.Valid

		insertMargin.Action.Value = v.Action.Value
		insertMargin.Action.Valid = v.Action.Valid

		insertMargin.Amount.Value = v.Amount.Value
		insertMargin.Amount.Valid = v.Amount.Valid

		insertMargin.PendingCredit.Value = v.PendingCredit.Value
		insertMargin.PendingCredit.Valid = v.PendingCredit.Valid

		insertMargin.PendingDebit.Value = v.PendingDebit.Value
		insertMargin.PendingDebit.Valid = v.PendingDebit.Valid

		insertMargin.ConfirmedDebit.Value = v.ConfirmedDebit.Value
		insertMargin.ConfirmedDebit.Valid = v.ConfirmedDebit.Valid

		insertMargin.PrevRealisedPnl.Value = v.PrevRealisedPnl.Value
		insertMargin.PrevRealisedPnl.Valid = v.PrevRealisedPnl.Valid

		insertMargin.PrevUnrealisedPnl.Value = v.PrevUnrealisedPnl.Value
		insertMargin.PrevUnrealisedPnl.Valid = v.PrevUnrealisedPnl.Valid

		insertMargin.GrossComm.Value = v.GrossComm.Value
		insertMargin.GrossComm.Valid = v.GrossComm.Valid

		insertMargin.GrossOpenCost.Value = v.GrossOpenCost.Value
		insertMargin.GrossOpenCost.Valid = v.GrossOpenCost.Valid

		insertMargin.GrossOpenPremium.Value = v.GrossOpenPremium.Value
		insertMargin.GrossOpenPremium.Valid = v.GrossOpenPremium.Valid

		insertMargin.GrossExecCost.Value = v.GrossExecCost.Value
		insertMargin.GrossExecCost.Valid = v.GrossExecCost.Valid

		insertMargin.GrossMarkValue.Value = v.GrossMarkValue.Value
		insertMargin.GrossMarkValue.Valid = v.GrossMarkValue.Valid

		insertMargin.RiskLimit.Value = v.RiskLimit.Value
		insertMargin.RiskLimit.Valid = v.RiskLimit.Valid

		insertMargin.TaxableMargin.Value = v.TaxableMargin.Value
		insertMargin.TaxableMargin.Valid = v.TaxableMargin.Valid

		insertMargin.InitMargin.Value = v.InitMargin.Value
		insertMargin.InitMargin.Valid = v.InitMargin.Valid

		insertMargin.MaintMargin.Value = v.MaintMargin.Value
		insertMargin.MaintMargin.Valid = v.MaintMargin.Valid

		insertMargin.SessionMargin.Value = v.SessionMargin.Value
		insertMargin.SessionMargin.Valid = v.SessionMargin.Valid

		insertMargin.TargetExcessMargin.Value = v.TargetExcessMargin.Value
		insertMargin.TargetExcessMargin.Valid = v.TargetExcessMargin.Valid

		insertMargin.VarMargin.Value = v.VarMargin.Value
		insertMargin.VarMargin.Valid = v.VarMargin.Valid

		insertMargin.RealisedPnl.Value = v.RealisedPnl.Value
		insertMargin.RealisedPnl.Valid = v.RealisedPnl.Valid

		insertMargin.UnrealisedPnl.Value = v.UnrealisedPnl.Value
		insertMargin.UnrealisedPnl.Valid = v.UnrealisedPnl.Valid

		insertMargin.UnrealisedPnl.Value = v.UnrealisedPnl.Value
		insertMargin.UnrealisedPnl.Valid = v.UnrealisedPnl.Valid

		insertMargin.UnrealisedProfit.Value = v.UnrealisedProfit.Value
		insertMargin.UnrealisedProfit.Valid = v.UnrealisedProfit.Valid

		insertMargin.SyntheticMargin.Value = v.SyntheticMargin.Value
		insertMargin.SyntheticMargin.Valid = v.SyntheticMargin.Valid

		insertMargin.WalletBalance.Value = v.WalletBalance.Value
		insertMargin.WalletBalance.Valid = v.WalletBalance.Valid

		insertMargin.MarginBalance.Value = v.MarginBalance.Value
		insertMargin.MarginBalance.Valid = v.MarginBalance.Valid

		insertMargin.MarginBalancePcnt.Value = v.MarginBalancePcnt.Value
		insertMargin.MarginBalancePcnt.Valid = v.MarginBalancePcnt.Valid

		insertMargin.MarginLeverage.Value = v.MarginLeverage.Value
		insertMargin.MarginLeverage.Valid = v.MarginLeverage.Valid

		insertMargin.MarginUsedPcnt.Value = v.MarginUsedPcnt.Value
		insertMargin.MarginUsedPcnt.Valid = v.MarginUsedPcnt.Valid

		insertMargin.ExcessMargin.Value = v.ExcessMargin.Value
		insertMargin.ExcessMargin.Valid = v.ExcessMargin.Valid

		insertMargin.ExcessMarginPcnt.Value = v.ExcessMarginPcnt.Value
		insertMargin.ExcessMarginPcnt.Valid = v.ExcessMarginPcnt.Valid

		insertMargin.AvailableMargin.Value = v.AvailableMargin.Value
		insertMargin.AvailableMargin.Valid = v.AvailableMargin.Valid

		insertMargin.WithdrawableMargin.Value = v.WithdrawableMargin.Value
		insertMargin.WithdrawableMargin.Valid = v.WithdrawableMargin.Valid

		insertMargin.Timestamp.Value = v.Timestamp.Value
		insertMargin.Timestamp.Valid = v.Timestamp.Valid

		insertMargin.GrossLastValue.Value = v.GrossLastValue.Value
		insertMargin.GrossLastValue.Valid = v.GrossLastValue.Valid

		*margin = append(*margin, insertMargin)
	}
	InfoLogger.Println("Margin Partials Processed")
}

func (margin *MarginSlice) MarginInsert(inserts *[]MarginResponseData) {
	InfoLogger.Println("Margin Inserts Processing")
	for _, v := range *inserts {
		var insertMargin swagger.Margin

		insertMargin.Account.Value = v.Account.Value
		insertMargin.Account.Valid = v.Account.Valid

		insertMargin.Currency.Value = v.Currency.Value
		insertMargin.Currency.Valid = v.Currency.Valid

		insertMargin.RiskLimit.Value = v.RiskLimit.Value
		insertMargin.RiskLimit.Valid = v.RiskLimit.Valid

		insertMargin.PrevState.Value = v.PrevState.Value
		insertMargin.PrevState.Valid = v.PrevState.Valid

		insertMargin.State.Value = v.State.Value
		insertMargin.State.Valid = v.State.Valid

		insertMargin.Action.Value = v.Action.Value
		insertMargin.Action.Valid = v.Action.Valid

		insertMargin.Amount.Value = v.Amount.Value
		insertMargin.Amount.Valid = v.Amount.Valid

		insertMargin.PendingCredit.Value = v.PendingCredit.Value
		insertMargin.PendingCredit.Valid = v.PendingCredit.Valid

		insertMargin.PendingDebit.Value = v.PendingDebit.Value
		insertMargin.PendingDebit.Valid = v.PendingDebit.Valid

		insertMargin.ConfirmedDebit.Value = v.ConfirmedDebit.Value
		insertMargin.ConfirmedDebit.Valid = v.ConfirmedDebit.Valid

		insertMargin.PrevRealisedPnl.Value = v.PrevRealisedPnl.Value
		insertMargin.PrevRealisedPnl.Valid = v.PrevRealisedPnl.Valid

		insertMargin.PrevUnrealisedPnl.Value = v.PrevUnrealisedPnl.Value
		insertMargin.PrevUnrealisedPnl.Valid = v.PrevUnrealisedPnl.Valid

		insertMargin.GrossComm.Value = v.GrossComm.Value
		insertMargin.GrossComm.Valid = v.GrossComm.Valid

		insertMargin.GrossOpenCost.Value = v.GrossOpenCost.Value
		insertMargin.GrossOpenCost.Valid = v.GrossOpenCost.Valid

		insertMargin.GrossOpenPremium.Value = v.GrossOpenPremium.Value
		insertMargin.GrossOpenPremium.Valid = v.GrossOpenPremium.Valid

		insertMargin.GrossExecCost.Value = v.GrossExecCost.Value
		insertMargin.GrossExecCost.Valid = v.GrossExecCost.Valid

		insertMargin.GrossMarkValue.Value = v.GrossMarkValue.Value
		insertMargin.GrossMarkValue.Valid = v.GrossMarkValue.Valid

		insertMargin.RiskLimit.Value = v.RiskLimit.Value
		insertMargin.RiskLimit.Valid = v.RiskLimit.Valid

		insertMargin.TaxableMargin.Value = v.TaxableMargin.Value
		insertMargin.TaxableMargin.Valid = v.TaxableMargin.Valid

		insertMargin.InitMargin.Value = v.InitMargin.Value
		insertMargin.InitMargin.Valid = v.InitMargin.Valid

		insertMargin.MaintMargin.Value = v.MaintMargin.Value
		insertMargin.MaintMargin.Valid = v.MaintMargin.Valid

		insertMargin.SessionMargin.Value = v.SessionMargin.Value
		insertMargin.SessionMargin.Valid = v.SessionMargin.Valid

		insertMargin.TargetExcessMargin.Value = v.TargetExcessMargin.Value
		insertMargin.TargetExcessMargin.Valid = v.TargetExcessMargin.Valid

		insertMargin.VarMargin.Value = v.VarMargin.Value
		insertMargin.VarMargin.Valid = v.VarMargin.Valid

		insertMargin.RealisedPnl.Value = v.RealisedPnl.Value
		insertMargin.RealisedPnl.Valid = v.RealisedPnl.Valid

		insertMargin.UnrealisedPnl.Value = v.UnrealisedPnl.Value
		insertMargin.UnrealisedPnl.Valid = v.UnrealisedPnl.Valid

		insertMargin.UnrealisedPnl.Value = v.UnrealisedPnl.Value
		insertMargin.UnrealisedPnl.Valid = v.UnrealisedPnl.Valid

		insertMargin.UnrealisedProfit.Value = v.UnrealisedProfit.Value
		insertMargin.UnrealisedProfit.Valid = v.UnrealisedProfit.Valid

		insertMargin.SyntheticMargin.Value = v.SyntheticMargin.Value
		insertMargin.SyntheticMargin.Valid = v.SyntheticMargin.Valid

		insertMargin.WalletBalance.Value = v.WalletBalance.Value
		insertMargin.WalletBalance.Valid = v.WalletBalance.Valid

		insertMargin.MarginBalance.Value = v.MarginBalance.Value
		insertMargin.MarginBalance.Valid = v.MarginBalance.Valid

		insertMargin.MarginBalancePcnt.Value = v.MarginBalancePcnt.Value
		insertMargin.MarginBalancePcnt.Valid = v.MarginBalancePcnt.Valid

		insertMargin.MarginLeverage.Value = v.MarginLeverage.Value
		insertMargin.MarginLeverage.Valid = v.MarginLeverage.Valid

		insertMargin.MarginUsedPcnt.Value = v.MarginUsedPcnt.Value
		insertMargin.MarginUsedPcnt.Valid = v.MarginUsedPcnt.Valid

		insertMargin.ExcessMargin.Value = v.ExcessMargin.Value
		insertMargin.ExcessMargin.Valid = v.ExcessMargin.Valid

		insertMargin.ExcessMarginPcnt.Value = v.ExcessMarginPcnt.Value
		insertMargin.ExcessMarginPcnt.Valid = v.ExcessMarginPcnt.Valid

		insertMargin.AvailableMargin.Value = v.AvailableMargin.Value
		insertMargin.AvailableMargin.Valid = v.AvailableMargin.Valid

		insertMargin.WithdrawableMargin.Value = v.WithdrawableMargin.Value
		insertMargin.WithdrawableMargin.Valid = v.WithdrawableMargin.Valid

		insertMargin.Timestamp.Value = v.Timestamp.Value
		insertMargin.Timestamp.Valid = v.Timestamp.Valid

		insertMargin.GrossLastValue.Value = v.GrossLastValue.Value
		insertMargin.GrossLastValue.Valid = v.GrossLastValue.Valid

		*margin = append(*margin, insertMargin)
	}
	InfoLogger.Println("Margin Inserts Processed")
}

func (margin *MarginSlice) MarginUpdate(updates *[]MarginResponseData) {
	InfoLogger.Println("Margin Updates Processing")
	for _, u := range *updates {
		for i, v := range *margin {

			if u.Account.Value == v.Account.Value {
				if u.Account.Set {
					v.Account.Value = u.Account.Value
					v.Account.Valid = u.Account.Valid
				}

				if u.Currency.Set {
					v.Currency.Value = u.Currency.Value
					v.Currency.Valid = u.Currency.Valid
				}

				if u.RiskLimit.Set {
					v.RiskLimit.Value = u.RiskLimit.Value
					v.RiskLimit.Valid = u.RiskLimit.Valid
				}

				if u.PrevState.Set {
					v.PrevState.Value = u.PrevState.Value
					v.PrevState.Valid = u.PrevState.Valid
				}

				if u.State.Set {
					v.State.Value = u.State.Value
					v.State.Valid = u.State.Valid
				}

				if u.Action.Set {
					v.Action.Value = u.Action.Value
					v.Action.Valid = u.Action.Valid
				}

				if u.Amount.Set {
					v.Amount.Value = u.Amount.Value
					v.Amount.Valid = u.Amount.Valid
				}

				if u.PendingCredit.Set {
					v.PendingCredit.Value = u.PendingCredit.Value
					v.PendingCredit.Valid = u.PendingCredit.Valid
				}

				if u.PendingDebit.Set {
					v.PendingDebit.Value = u.PendingDebit.Value
					v.PendingDebit.Valid = u.PendingDebit.Valid
				}

				if u.ConfirmedDebit.Set {
					v.ConfirmedDebit.Value = u.ConfirmedDebit.Value
					v.ConfirmedDebit.Valid = u.ConfirmedDebit.Valid
				}

				if u.PrevRealisedPnl.Set {
					v.PrevRealisedPnl.Value = u.PrevRealisedPnl.Value
					v.PrevRealisedPnl.Valid = u.PrevRealisedPnl.Valid
				}

				if u.PrevUnrealisedPnl.Set {
					v.PrevUnrealisedPnl.Value = u.PrevUnrealisedPnl.Value
					v.PrevUnrealisedPnl.Valid = u.PrevUnrealisedPnl.Valid
				}

				if u.GrossComm.Set {
					v.GrossComm.Value = u.GrossComm.Value
					v.GrossComm.Valid = u.GrossComm.Valid
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

				if u.GrossMarkValue.Set {
					v.GrossMarkValue.Value = u.GrossMarkValue.Value
					v.GrossMarkValue.Valid = u.GrossMarkValue.Valid
				}

				if u.RiskValue.Set {
					v.RiskValue.Value = u.RiskValue.Value
					v.RiskValue.Valid = u.RiskValue.Valid
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

				if u.RealisedPnl.Set {
					v.RealisedPnl.Value = u.RealisedPnl.Value
					v.RealisedPnl.Valid = u.RealisedPnl.Valid
				}

				if u.UnrealisedPnl.Set {
					v.UnrealisedPnl.Value = u.UnrealisedPnl.Value
					v.UnrealisedPnl.Valid = u.UnrealisedPnl.Valid
				}

				if u.IndicativeTax.Set {
					v.IndicativeTax.Value = u.IndicativeTax.Value
					v.IndicativeTax.Valid = u.IndicativeTax.Valid
				}

				if u.UnrealisedProfit.Set {
					v.UnrealisedProfit.Value = u.UnrealisedProfit.Value
					v.UnrealisedProfit.Valid = u.UnrealisedProfit.Valid
				}

				if u.SyntheticMargin.Set {
					v.SyntheticMargin.Value = u.SyntheticMargin.Value
					v.SyntheticMargin.Valid = u.SyntheticMargin.Valid
				}

				if u.WalletBalance.Set {
					v.WalletBalance.Value = u.WalletBalance.Value
					v.WalletBalance.Valid = u.WalletBalance.Valid
				}

				if u.MarginBalance.Set {
					v.MarginBalance.Value = u.MarginBalance.Value
					v.MarginBalance.Valid = u.MarginBalance.Valid
				}

				if u.MarginBalancePcnt.Set {
					v.MarginBalancePcnt.Value = u.MarginBalancePcnt.Value
					v.MarginBalancePcnt.Valid = u.MarginBalancePcnt.Valid
				}

				if u.MarginLeverage.Set {
					v.MarginLeverage.Value = u.MarginLeverage.Value
					v.MarginLeverage.Valid = u.MarginLeverage.Valid
				}

				if u.MarginUsedPcnt.Set {
					v.MarginUsedPcnt.Value = u.MarginUsedPcnt.Value
					v.MarginUsedPcnt.Valid = u.MarginUsedPcnt.Valid
				}

				if u.ExcessMargin.Set {
					v.ExcessMargin.Value = u.ExcessMargin.Value
					v.ExcessMargin.Valid = u.ExcessMargin.Valid
				}

				if u.ExcessMarginPcnt.Set {
					v.ExcessMarginPcnt.Value = u.ExcessMarginPcnt.Value
					v.ExcessMarginPcnt.Valid = u.ExcessMarginPcnt.Valid
				}

				if u.AvailableMargin.Set {
					v.AvailableMargin.Value = u.AvailableMargin.Value
					v.AvailableMargin.Valid = u.AvailableMargin.Valid
				}

				if u.WithdrawableMargin.Set {
					v.WithdrawableMargin.Value = u.WithdrawableMargin.Value
					v.WithdrawableMargin.Valid = u.WithdrawableMargin.Valid
				}

				if u.Timestamp.Set {
					v.Timestamp.Value = u.Timestamp.Value
					v.Timestamp.Valid = u.Timestamp.Valid
				}

				if u.GrossLastValue.Set {
					v.GrossLastValue.Value = u.GrossLastValue.Value
					v.GrossLastValue.Valid = u.GrossLastValue.Valid
				}

				if u.Commission.Set {
					v.Commission.Value = u.Commission.Value
					v.Commission.Valid = u.Commission.Valid
				}

				later := (*margin)[i+1:]
				*margin = append((*margin)[:i], v)
				*margin = append(*margin, later...)
			}
		}
	}
	InfoLogger.Println("Margin Updates Processed")
}

func (margin *MarginSlice) MarginDelete(deletes *[]MarginResponseData) {
	InfoLogger.Println("Margin Deletes Processing")
	for _, u := range *deletes {
		for i, v := range *margin {
			if u.Account.Value == v.Account.Value {
				*margin = append((*margin)[:i], (*margin)[i+1:]...)
			}
		}
	}
	InfoLogger.Println("Margin Deletes Processed")
}
