package swagger

type UserPreferences struct {
	AlertOnLiquidations     NullBool     `json:"alertOnLiquidations,omitempty"`
	AnimationsEnabled       NullBool     `json:"animationsEnabled,omitempty"`
	AnnouncementsLastSeen   NullTime     `json:"announcementsLastSeen,omitempty"`
	ChatChannelID           NullFloat64  `json:"chatChannelID,omitempty"`
	ColorTheme              NullString   `json:"colorTheme,omitempty"`
	Currency                NullString   `json:"currency,omitempty"`
	Debug                   NullBool     `json:"debug,omitempty"`
	DisableEmails           []NullString `json:"disableEmails,omitempty"`
	HideConfirmDialogs      []NullString `json:"hideConfirmDialogs,omitempty"`
	HideConnectionModal     NullBool     `json:"hideConnectionModal,omitempty"`
	HideFromLeaderboard     NullBool     `json:"hideFromLeaderboard,omitempty"`
	HideNameFromLeaderboard NullBool     `json:"hideNameFromLeaderboard,omitempty"`
	HideNotifications       []NullString `json:"hideNotifications,omitempty"`
	Locale                  NullString   `json:"locale,omitempty"`
	MsgsSeen                []NullString `json:"msgsSeen,omitempty"`
	OrderBookBinning        *interface{} `json:"orderBookBinning,omitempty"`
	OrderBookType           NullString   `json:"orderBookType,omitempty"`
	OrderClearImmediate     NullBool     `json:"orderClearImmediate,omitempty"`
	OrderControlsPlusMinus  NullBool     `json:"orderControlsPlusMinus,omitempty"`
	ShowLocaleNumbers       NullBool     `json:"showLocaleNumbers,omitempty"`
	Sounds                  []NullString `json:"sounds,omitempty"`
	StrictIPCheck           NullBool     `json:"strictIPCheck,omitempty"`
	StrictTimeout           NullBool     `json:"strictTimeout,omitempty"`
	TickerGroup             NullString   `json:"tickerGroup,omitempty"`
	TickerPinned            NullBool     `json:"tickerPinned,omitempty"`
	TradeLayout             NullString   `json:"tradeLayout,omitempty"`
}
