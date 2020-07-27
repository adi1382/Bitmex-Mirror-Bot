package swagger

// Raw Order and Balance Data
type Execution struct {
	ExecID                NullString  `json:"execID"`
	OrderID               NullString  `json:"orderID,omitempty"`
	ClOrdID               NullString  `json:"clOrdID,omitempty"`
	ClOrdLinkID           NullString  `json:"clOrdLinkID,omitempty"`
	Account               NullInt     `json:"account,omitempty"`
	Symbol                NullString  `json:"symbol,omitempty"`
	Side                  NullString  `json:"side,omitempty"`
	LastQty               NullInt     `json:"lastQty,omitempty"`
	LastPx                NullFloat64 `json:"lastPx,omitempty"`
	UnderlyingLastPx      NullFloat64 `json:"underlyingLastPx,omitempty"`
	LastMkt               NullString  `json:"lastMkt,omitempty"`
	LastLiquidityInd      NullString  `json:"lastLiquidityInd,omitempty"`
	SimpleOrderQty        NullFloat64 `json:"simpleOrderQty,omitempty"`
	OrderQty              NullInt     `json:"orderQty,omitempty"`
	Price                 NullFloat64 `json:"price,omitempty"`
	DisplayQty            NullInt     `json:"displayQty,omitempty"`
	StopPx                NullFloat64 `json:"stopPx,omitempty"`
	PegOffsetValue        NullFloat64 `json:"pegOffsetValue,omitempty"`
	PegPriceType          NullString  `json:"pegPriceType,omitempty"`
	Currency              NullString  `json:"currency,omitempty"`
	SettlCurrency         NullString  `json:"settlCurrency,omitempty"`
	ExecType              NullString  `json:"execType,omitempty"`
	OrdType               NullString  `json:"ordType,omitempty"`
	TimeInForce           NullString  `json:"timeInForce,omitempty"`
	ExecInst              NullString  `json:"execInst,omitempty"`
	ContingencyType       NullString  `json:"contingencyType,omitempty"`
	ExDestination         NullString  `json:"exDestination,omitempty"`
	OrdStatus             NullString  `json:"ordStatus,omitempty"`
	Triggered             NullString  `json:"triggered,omitempty"`
	WorkingIndicator      NullBool    `json:"workingIndicator,omitempty"`
	OrdRejReason          NullString  `json:"ordRejReason,omitempty"`
	SimpleLeavesQty       NullFloat64 `json:"simpleLeavesQty,omitempty"`
	LeavesQty             NullInt     `json:"leavesQty,omitempty"`
	SimpleCumQty          NullFloat64 `json:"simpleCumQty,omitempty"`
	CumQty                NullInt     `json:"cumQty,omitempty"`
	AvgPx                 NullFloat64 `json:"avgPx,omitempty"`
	Commission            NullFloat64 `json:"commission,omitempty"`
	TradePublishIndicator NullString  `json:"tradePublishIndicator,omitempty"`
	MultiLegReportingType NullString  `json:"multiLegReportingType,omitempty"`
	Text                  NullString  `json:"text,omitempty"`
	TrdMatchID            NullString  `json:"trdMatchID,omitempty"`
	ExecCost              NullInt     `json:"execCost,omitempty"`
	ExecComm              NullInt     `json:"execComm,omitempty"`
	HomeNotional          NullFloat64 `json:"homeNotional,omitempty"`
	ForeignNotional       NullFloat64 `json:"foreignNotional,omitempty"`
	TransactTime          NullTime    `json:"transactTime,omitempty"`
	Timestamp             NullTime    `json:"timestamp,omitempty"`
}
