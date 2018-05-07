package exchanges

import (
	"github.com/CryptoTradingBot/exchanges/nonce"
	"time"
	"log"
	"github.com/CryptoTradingBot/exchanges/common"
	"github.com/CryptoTradingBot/exchanges/request"
)

const (
	warningBase64DecryptSecretKeyFailed              = "WARNING -- Exchange %s unable to base64 decode secret key.. Disabling Authenticated API support."
	WarningAuthenticatedRequestWithoutCredentialsSet = "WARNING -- Exchange %s authenticated HTTP request called but not supported due to unset/default API keys."
	ErrExchangeNotFound                              = "Exchange not found in dataset."
	// DefaultHTTPTimeout is the default HTTP/HTTPS Timeout for exchange requests
	DefaultHTTPTimeout = time.Second * 15
)

// Base stores the individual exchange information
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
	RequestCurrencyPairFormat   CurrencyPairFormatConfig
	ConfigCurrencyPairFormat    CurrencyPairFormatConfig
	*request.Requester
}

// CurrencyPairFormatConfig stores the users preferred currency pair display
type CurrencyPairFormatConfig struct {
	Uppercase bool
	Delimiter string `json:",omitempty"`
	Separator string `json:",omitempty"`
	Index     string `json:",omitempty"`
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

// SetAPIKeys is a method that sets the current API keys for the exchange
func (e *Base) SetAPIKeys(APIKey, APISecret, ClientID string, b64Decode bool) {
	if !e.AuthenticatedAPISupport {
		return
	}

	e.APIKey = APIKey
	e.ClientID = ClientID

	if b64Decode {
		result, err := common.Base64Decode(APISecret)
		if err != nil {
			e.AuthenticatedAPISupport = false
			log.Printf(warningBase64DecryptSecretKeyFailed, e.Name)
		}
		e.APISecret = string(result)
	} else {
		e.APISecret = APISecret
	}
}
