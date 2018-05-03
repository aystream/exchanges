package bitmex

import "time"

// Ticker ticker data
type Ticker struct {
	Symbol string  `json:"symbol"`
	Last   float64 `json:"lastPrice"`
	Bid    float64 `json:"bidPrice"`
	Ask    float64 `json:"askPrice"`
	High   float64 `json:"highPrice"`
	low    float64 `json:"lowPrice"`
}

// Placement, Cancellation, Amending, and History
type Order struct {
	OrderID               string    `json:"orderID"`
	ClOrdID               string    `json:"clOrdID,omitempty"`
	ClOrdLinkID           string    `json:"clOrdLinkID,omitempty"`
	Account               float32   `json:"account,omitempty"`
	Symbol                string    `json:"symbol,omitempty"`
	Side                  string    `json:"side,omitempty"`
	SimpleOrderQty        float64   `json:"simpleOrderQty,omitempty"`
	OrderQty              float32   `json:"orderQty,omitempty"`
	Price                 float64   `json:"price,omitempty"`
	DisplayQty            float32   `json:"displayQty,omitempty"`
	StopPx                float64   `json:"stopPx,omitempty"`
	PegOffsetValue        float64   `json:"pegOffsetValue,omitempty"`
	PegPriceType          string    `json:"pegPriceType,omitempty"`
	Currency              string    `json:"currency,omitempty"`
	SettlCurrency         string    `json:"settlCurrency,omitempty"`
	OrdType               string    `json:"ordType,omitempty"`
	TimeInForce           string    `json:"timeInForce,omitempty"`
	ExecInst              string    `json:"execInst,omitempty"`
	ContingencyType       string    `json:"contingencyType,omitempty"`
	ExDestination         string    `json:"exDestination,omitempty"`
	OrdStatus             string    `json:"ordStatus,omitempty"`
	Triggered             string    `json:"triggered,omitempty"`
	WorkingIndicator      bool      `json:"workingIndicator,omitempty"`
	OrdRejReason          string    `json:"ordRejReason,omitempty"`
	SimpleLeavesQty       float64   `json:"simpleLeavesQty,omitempty"`
	LeavesQty             float32   `json:"leavesQty,omitempty"`
	SimpleCumQty          float64   `json:"simpleCumQty,omitempty"`
	CumQty                float32   `json:"cumQty,omitempty"`
	AvgPx                 float64   `json:"avgPx,omitempty"`
	MultiLegReportingType string    `json:"multiLegReportingType,omitempty"`
	Text                  string    `json:"text,omitempty"`
	TransactTime          time.Time `json:"transactTime,omitempty"`
	Timestamp             time.Time `json:"timestamp,omitempty"`
}

// Summary of Open and Closed Positions
type Position struct {
	Account              float32   `json:"account"`
	Symbol               string    `json:"symbol"`
	Currency             string    `json:"currency"`
	Underlying           string    `json:"underlying,omitempty"`
	QuoteCurrency        string    `json:"quoteCurrency,omitempty"`
	Commission           float64   `json:"commission,omitempty"`
	InitMarginReq        float64   `json:"initMarginReq,omitempty"`
	MaintMarginReq       float64   `json:"maintMarginReq,omitempty"`
	RiskLimit            float32   `json:"riskLimit,omitempty"`
	Leverage             float64   `json:"leverage,omitempty"`
	CrossMargin          bool      `json:"crossMargin,omitempty"`
	DeleveragePercentile float64   `json:"deleveragePercentile,omitempty"`
	RebalancedPnl        float32   `json:"rebalancedPnl,omitempty"`
	PrevRealisedPnl      float32   `json:"prevRealisedPnl,omitempty"`
	PrevUnrealisedPnl    float32   `json:"prevUnrealisedPnl,omitempty"`
	PrevClosePrice       float64   `json:"prevClosePrice,omitempty"`
	OpeningTimestamp     time.Time `json:"openingTimestamp,omitempty"`
	OpeningQty           float32   `json:"openingQty,omitempty"`
	OpeningCost          float32   `json:"openingCost,omitempty"`
	OpeningComm          float32   `json:"openingComm,omitempty"`
	OpenOrderBuyQty      float32   `json:"openOrderBuyQty,omitempty"`
	OpenOrderBuyCost     float32   `json:"openOrderBuyCost,omitempty"`
	OpenOrderBuyPremium  float32   `json:"openOrderBuyPremium,omitempty"`
	OpenOrderSellQty     float32   `json:"openOrderSellQty,omitempty"`
	OpenOrderSellCost    float32   `json:"openOrderSellCost,omitempty"`
	OpenOrderSellPremium float32   `json:"openOrderSellPremium,omitempty"`
	ExecBuyQty           float32   `json:"execBuyQty,omitempty"`
	ExecBuyCost          float32   `json:"execBuyCost,omitempty"`
	ExecSellQty          float32   `json:"execSellQty,omitempty"`
	ExecSellCost         float32   `json:"execSellCost,omitempty"`
	ExecQty              float32   `json:"execQty,omitempty"`
	ExecCost             float32   `json:"execCost,omitempty"`
	ExecComm             float32   `json:"execComm,omitempty"`
	CurrentTimestamp     time.Time `json:"currentTimestamp,omitempty"`
	CurrentQty           float32   `json:"currentQty,omitempty"`
	CurrentCost          float32   `json:"currentCost,omitempty"`
	CurrentComm          float32   `json:"currentComm,omitempty"`
	RealisedCost         float32   `json:"realisedCost,omitempty"`
	UnrealisedCost       float32   `json:"unrealisedCost,omitempty"`
	GrossOpenCost        float32   `json:"grossOpenCost,omitempty"`
	GrossOpenPremium     float32   `json:"grossOpenPremium,omitempty"`
	GrossExecCost        float32   `json:"grossExecCost,omitempty"`
	IsOpen               bool      `json:"isOpen,omitempty"`
	MarkPrice            float64   `json:"markPrice,omitempty"`
	MarkValue            float32   `json:"markValue,omitempty"`
	RiskValue            float32   `json:"riskValue,omitempty"`
	HomeNotional         float64   `json:"homeNotional,omitempty"`
	ForeignNotional      float64   `json:"foreignNotional,omitempty"`
	PosState             string    `json:"posState,omitempty"`
	PosCost              float32   `json:"posCost,omitempty"`
	PosCost2             float32   `json:"posCost2,omitempty"`
	PosCross             float32   `json:"posCross,omitempty"`
	PosInit              float32   `json:"posInit,omitempty"`
	PosComm              float32   `json:"posComm,omitempty"`
	PosLoss              float32   `json:"posLoss,omitempty"`
	PosMargin            float32   `json:"posMargin,omitempty"`
	PosMaint             float32   `json:"posMaint,omitempty"`
	PosAllowance         float32   `json:"posAllowance,omitempty"`
	TaxableMargin        float32   `json:"taxableMargin,omitempty"`
	InitMargin           float32   `json:"initMargin,omitempty"`
	MaintMargin          float32   `json:"maintMargin,omitempty"`
	SessionMargin        float32   `json:"sessionMargin,omitempty"`
	TargetExcessMargin   float32   `json:"targetExcessMargin,omitempty"`
	VarMargin            float32   `json:"varMargin,omitempty"`
	RealisedGrossPnl     float32   `json:"realisedGrossPnl,omitempty"`
	RealisedTax          float32   `json:"realisedTax,omitempty"`
	RealisedPnl          float32   `json:"realisedPnl,omitempty"`
	UnrealisedGrossPnl   float32   `json:"unrealisedGrossPnl,omitempty"`
	LongBankrupt         float32   `json:"longBankrupt,omitempty"`
	ShortBankrupt        float32   `json:"shortBankrupt,omitempty"`
	TaxBase              float32   `json:"taxBase,omitempty"`
	IndicativeTaxRate    float64   `json:"indicativeTaxRate,omitempty"`
	IndicativeTax        float32   `json:"indicativeTax,omitempty"`
	UnrealisedTax        float32   `json:"unrealisedTax,omitempty"`
	UnrealisedPnl        float32   `json:"unrealisedPnl,omitempty"`
	UnrealisedPnlPcnt    float64   `json:"unrealisedPnlPcnt,omitempty"`
	UnrealisedRoePcnt    float64   `json:"unrealisedRoePcnt,omitempty"`
	SimpleQty            float64   `json:"simpleQty,omitempty"`
	SimpleCost           float64   `json:"simpleCost,omitempty"`
	SimpleValue          float64   `json:"simpleValue,omitempty"`
	SimplePnl            float64   `json:"simplePnl,omitempty"`
	SimplePnlPcnt        float64   `json:"simplePnlPcnt,omitempty"`
	AvgCostPrice         float64   `json:"avgCostPrice,omitempty"`
	AvgEntryPrice        float64   `json:"avgEntryPrice,omitempty"`
	BreakEvenPrice       float64   `json:"breakEvenPrice,omitempty"`
	MarginCallPrice      float64   `json:"marginCallPrice,omitempty"`
	LiquidationPrice     float64   `json:"liquidationPrice,omitempty"`
	BankruptPrice        float64   `json:"bankruptPrice,omitempty"`
	Timestamp            time.Time `json:"timestamp,omitempty"`
	LastPrice            float64   `json:"lastPrice,omitempty"`
	LastValue            float32   `json:"lastValue,omitempty"`
}

// Wallet
type Wallet struct {
	Account float32 `json:"account"`
	Currency string `json:"currency"`
	PrevDeposited float32 `json:"prevDeposited,omitempty"`
	PrevWithdrawn float32 `json:"prevWithdrawn,omitempty"`
	PrevTransferIn float32 `json:"prevTransferIn,omitempty"`
	PrevTransferOut float32 `json:"prevTransferOut,omitempty"`
	PrevAmount float32 `json:"prevAmount,omitempty"`
	PrevTimestamp time.Time `json:"prevTimestamp,omitempty"`
	DeltaDeposited float32 `json:"deltaDeposited,omitempty"`
	DeltaWithdrawn float32 `json:"deltaWithdrawn,omitempty"`
	DeltaTransferIn float32 `json:"deltaTransferIn,omitempty"`
	DeltaTransferOut float32 `json:"deltaTransferOut,omitempty"`
	DeltaAmount float32 `json:"deltaAmount,omitempty"`
	Deposited float32 `json:"deposited,omitempty"`
	Withdrawn float32 `json:"withdrawn,omitempty"`
	TransferIn float32 `json:"transferIn,omitempty"`
	TransferOut float32 `json:"transferOut,omitempty"`
	Amount float32 `json:"amount,omitempty"`
	PendingCredit float32 `json:"pendingCredit,omitempty"`
	PendingDebit float32 `json:"pendingDebit,omitempty"`
	ConfirmedDebit float32 `json:"confirmedDebit,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
	Addr string `json:"addr,omitempty"`
	Script string `json:"script,omitempty"`
	WithdrawalLock []string `json:"withdrawalLock,omitempty"`
}
// Margin info
type Margin struct {
	Account float32 `json:"account"`
	Currency string `json:"currency"`
	RiskLimit float32 `json:"riskLimit,omitempty"`
	PrevState string `json:"prevState,omitempty"`
	State string `json:"state,omitempty"`
	Action string `json:"action,omitempty"`
	Amount float32 `json:"amount,omitempty"`
	PendingCredit float32 `json:"pendingCredit,omitempty"`
	PendingDebit float32 `json:"pendingDebit,omitempty"`
	ConfirmedDebit float32 `json:"confirmedDebit,omitempty"`
	PrevRealisedPnl float32 `json:"prevRealisedPnl,omitempty"`
	PrevUnrealisedPnl float32 `json:"prevUnrealisedPnl,omitempty"`
	GrossComm float32 `json:"grossComm,omitempty"`
	GrossOpenCost float32 `json:"grossOpenCost,omitempty"`
	GrossOpenPremium float32 `json:"grossOpenPremium,omitempty"`
	GrossExecCost float32 `json:"grossExecCost,omitempty"`
	GrossMarkValue float32 `json:"grossMarkValue,omitempty"`
	RiskValue float32 `json:"riskValue,omitempty"`
	TaxableMargin float32 `json:"taxableMargin,omitempty"`
	InitMargin float32 `json:"initMargin,omitempty"`
	MaintMargin float32 `json:"maintMargin,omitempty"`
	SessionMargin float32 `json:"sessionMargin,omitempty"`
	TargetExcessMargin float32 `json:"targetExcessMargin,omitempty"`
	VarMargin float32 `json:"varMargin,omitempty"`
	RealisedPnl float32 `json:"realisedPnl,omitempty"`
	UnrealisedPnl float32 `json:"unrealisedPnl,omitempty"`
	IndicativeTax float32 `json:"indicativeTax,omitempty"`
	UnrealisedProfit float32 `json:"unrealisedProfit,omitempty"`
	SyntheticMargin float32 `json:"syntheticMargin,omitempty"`
	WalletBalance float32 `json:"walletBalance,omitempty"`
	MarginBalance float32 `json:"marginBalance,omitempty"`
	MarginBalancePcnt float64 `json:"marginBalancePcnt,omitempty"`
	MarginLeverage float64 `json:"marginLeverage,omitempty"`
	MarginUsedPcnt float64 `json:"marginUsedPcnt,omitempty"`
	ExcessMargin float32 `json:"excessMargin,omitempty"`
	ExcessMarginPcnt float64 `json:"excessMarginPcnt,omitempty"`
	AvailableMargin float32 `json:"availableMargin,omitempty"`
	WithdrawableMargin float32 `json:"withdrawableMargin,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
	GrossLastValue float32 `json:"grossLastValue,omitempty"`
	Commission float64 `json:"commission,omitempty"`
}

type OrderBookL2 struct {
	Symbol string `json:"symbol"`
	Id float32 `json:"id"`
	Side string `json:"side"`
	Size float32 `json:"size,omitempty"`
	Price float64 `json:"price,omitempty"`
}