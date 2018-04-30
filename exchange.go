package exchanges

import (
	"github.com/CryptoTradingBot/exchanges/nonce"
	"time"
)

const (
	WarningAuthenticatedRequestWithoutCredentialsSet = "WARNING -- Exchange %s authenticated HTTP request called but not supported due to unset/default API keys."
	ErrExchangeNotFound = "Exchange not found in dataset."
)

type Base struct {
	Name                        string
	Enabled                     bool
	Verbose                     bool
	Websocket                   bool
	RESTPollingDelay            time.Duration
	AuthenticatedAPISupport     bool
	APISecret, APIKey, ClientID string
	Nonce                       nonce.Nonce
	TakerFee, MakerFee, Fee     float64
	BaseCurrencies              []string
	AvailablePairs              []string
	EnabledPairs                []string
	AssetTypes                  []string
	PairsLastUpdated            int64
	SupportsAutoPairUpdating    bool
	WebsocketURL                string
	APIUrl                      string
}


// standard function in bot for exchange
type BotExchangeInterface interface {
	GetCandles()
	GetOrder()
	GetOrders()
	GetOpenOrders()
	GetOpenPositions()
	ClosePosition()
	EditOrderPrice()
	CreateOrder()
	CancelAllOpenOrders()
	GetWallet()
	GetMargin()
	GetOrderBook()
	SetLeverage() //?


	/**/
	GetMarkets()
	BuyLimit()
	BuyMarket()
	SellLimit()
	SellMarket()

	GetTicker()
	GetMarketSummaries()
	GetMarketSummary()

}