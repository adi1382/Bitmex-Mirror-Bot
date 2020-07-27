package swagger

type Wallet struct {
	Account          NullInt      `json:"account"`
	Currency         NullString   `json:"currency"`
	PrevDeposited    NullInt      `json:"prevDeposited,omitempty"`
	PrevWithdrawn    NullInt      `json:"prevWithdrawn,omitempty"`
	PrevTransferIn   NullInt      `json:"prevTransferIn,omitempty"`
	PrevTransferOut  NullInt      `json:"prevTransferOut,omitempty"`
	PrevAmount       NullInt      `json:"prevAmount,omitempty"`
	PrevTimestamp    NullTime     `json:"prevTimestamp,omitempty"`
	DeltaDeposited   NullInt      `json:"deltaDeposited,omitempty"`
	DeltaWithdrawn   NullInt      `json:"deltaWithdrawn,omitempty"`
	DeltaTransferIn  NullInt      `json:"deltaTransferIn,omitempty"`
	DeltaTransferOut NullInt      `json:"deltaTransferOut,omitempty"`
	DeltaAmount      NullInt      `json:"deltaAmount,omitempty"`
	Deposited        NullInt      `json:"deposited,omitempty"`
	Withdrawn        NullInt      `json:"withdrawn,omitempty"`
	TransferIn       NullInt      `json:"transferIn,omitempty"`
	TransferOut      NullInt      `json:"transferOut,omitempty"`
	Amount           NullInt      `json:"amount,omitempty"`
	PendingCredit    NullInt      `json:"pendingCredit,omitempty"`
	PendingDebit     NullInt      `json:"pendingDebit,omitempty"`
	ConfirmedDebit   NullInt      `json:"confirmedDebit,omitempty"`
	Timestamp        NullTime     `json:"timestamp,omitempty"`
	Addr             NullString   `json:"addr,omitempty"`
	Script           NullString   `json:"script,omitempty"`
	WithdrawalLock   []NullString `json:"withdrawalLock,omitempty"`
}
