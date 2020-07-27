/*
 * BitMEX API
 *
 * ## REST API for the BitMEX Trading Platform  [View Changelog](/app/apiChangelog)    #### Getting Started   ##### Fetching Data  All REST endpoints are documented below. You can try out any query right from this interface.  Most table queries accept `count`, `start`, and `reverse` params. Set `reverse=true` to get rows newest-first.  Additional documentation regarding filters, timestamps, and authentication is available in [the main API documentation](https://www.bitmex.com/app/restAPI).  *All* table data is available via the [Websocket](/app/wsAPI). We highly recommend using the socket if you want to have the quickest possible data without being subject to ratelimits.  ##### Return Types  By default, all data is returned as JSON. Send `?_format=csv` to get CSV data or `?_format=xml` to get XML data.  ##### Trade Data Queries  *This is only a small subset of what is available, to get you started.*  Fill in the parameters and click the `Try it out!` button to try any of these queries.  * [Pricing Data](#!/Quote/Quote_get)  * [Trade Data](#!/Trade/Trade_get)  * [OrderBook Data](#!/OrderBook/OrderBook_getL2)  * [Settlement Data](#!/Settlement/Settlement_get)  * [Exchange Statistics](#!/Stats/Stats_history)  Every function of the BitMEX.com platform is exposed here and documented. Many more functions are available.  ##### Swagger Specification  [⇩ Download Swagger JSON](swagger.json)    ## All API Endpoints  Click to expand a section.
 *
 * OpenAPI spec version: 1.2.0
 * Contact: support@bitmex.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

type Margin struct {
	Account            NullFloat64 `json:"account"`
	Currency           NullString  `json:"currency"`
	RiskLimit          NullFloat64 `json:"riskLimit,omitempty"`
	PrevState          NullString  `json:"prevState,omitempty"`
	State              NullString  `json:"state,omitempty"`
	Action             NullString  `json:"action,omitempty"`
	Amount             NullFloat64 `json:"amount,omitempty"`
	PendingCredit      NullFloat64 `json:"pendingCredit,omitempty"`
	PendingDebit       NullFloat64 `json:"pendingDebit,omitempty"`
	ConfirmedDebit     NullFloat64 `json:"confirmedDebit,omitempty"`
	PrevRealisedPnl    NullFloat64 `json:"prevRealisedPnl,omitempty"`
	PrevUnrealisedPnl  NullFloat64 `json:"prevUnrealisedPnl,omitempty"`
	GrossComm          NullFloat64 `json:"grossComm,omitempty"`
	GrossOpenCost      NullFloat64 `json:"grossOpenCost,omitempty"`
	GrossOpenPremium   NullFloat64 `json:"grossOpenPremium,omitempty"`
	GrossExecCost      NullFloat64 `json:"grossExecCost,omitempty"`
	GrossMarkValue     NullFloat64 `json:"grossMarkValue,omitempty"`
	RiskValue          NullFloat64 `json:"riskValue,omitempty"`
	TaxableMargin      NullFloat64 `json:"taxableMargin,omitempty"`
	InitMargin         NullFloat64 `json:"initMargin,omitempty"`
	MaintMargin        NullFloat64 `json:"maintMargin,omitempty"`
	SessionMargin      NullFloat64 `json:"sessionMargin,omitempty"`
	TargetExcessMargin NullFloat64 `json:"targetExcessMargin,omitempty"`
	VarMargin          NullFloat64 `json:"varMargin,omitempty"`
	RealisedPnl        NullFloat64 `json:"realisedPnl,omitempty"`
	UnrealisedPnl      NullFloat64 `json:"unrealisedPnl,omitempty"`
	IndicativeTax      NullFloat64 `json:"indicativeTax,omitempty"`
	UnrealisedProfit   NullFloat64 `json:"unrealisedProfit,omitempty"`
	SyntheticMargin    NullFloat64 `json:"syntheticMargin,omitempty"`
	WalletBalance      NullFloat64 `json:"walletBalance,omitempty"`
	MarginBalance      NullFloat64 `json:"marginBalance,omitempty"`
	MarginBalancePcnt  NullFloat64 `json:"marginBalancePcnt,omitempty"`
	MarginLeverage     NullFloat64 `json:"marginLeverage,omitempty"`
	MarginUsedPcnt     NullFloat64 `json:"marginUsedPcnt,omitempty"`
	ExcessMargin       NullFloat64 `json:"excessMargin,omitempty"`
	ExcessMarginPcnt   NullFloat64 `json:"excessMarginPcnt,omitempty"`
	AvailableMargin    NullFloat64 `json:"availableMargin,omitempty"`
	WithdrawableMargin NullFloat64 `json:"withdrawableMargin,omitempty"`
	Timestamp          NullTime    `json:"timestamp,omitempty"`
	GrossLastValue     NullFloat64 `json:"grossLastValue,omitempty"`
	Commission         NullFloat64 `json:"commission,omitempty"`
}
