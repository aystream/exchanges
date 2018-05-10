package exchanges

import (
	"github.com/CryptoTradingBot/exchanges/nonce"
	"time"
	"log"
	"github.com/CryptoTradingBot/exchanges/common"
	"github.com/CryptoTradingBot/exchanges/request"
	"net/http"
	"github.com/CryptoTradingBot/exchanges/config"
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
	RequestCurrencyPairFormat   config.CurrencyPairFormatConfig
	ConfigCurrencyPairFormat    config.CurrencyPairFormatConfig
	*request.Requester
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

// SetEnabled is a method that sets if the exchange is enabled
func (e *Base) SetEnabled(enabled bool) {
	e.Enabled = enabled
}

// IsEnabled is a method that returns if the current exchange is enabled
func (e *Base) IsEnabled() bool {
	return e.Enabled
}

// SetHTTPClientTimeout sets the timeout value for the exchanges
// HTTP Client
func (e *Base) SetHTTPClientTimeout(t time.Duration) {
	if e.Requester == nil {
		e.Requester = request.New(e.Name, request.NewRateLimit(time.Second, 0), request.NewRateLimit(time.Second, 0), new(http.Client))
	}
	e.Requester.HTTPClient.Timeout = t
}


// SetAutoPairDefaults sets the default values for whether or not the exchange
// supports auto pair updating or not
func (e *Base) SetAutoPairDefaults() error {
	cfg := config.GetConfig()
	exch, err := cfg.GetExchangeConfig(e.Name)
	if err != nil {
		return err
	}

	update := false
	if e.SupportsAutoPairUpdating {
		if !exch.SupportsAutoPairUpdates {
			exch.SupportsAutoPairUpdates = true
			update = true
		}
	} else {
		if exch.PairsLastUpdated == 0 {
			exch.PairsLastUpdated = time.Now().Unix()
			e.PairsLastUpdated = exch.PairsLastUpdated
			update = true
		}
	}

	if update {
		return cfg.UpdateExchangeConfig(exch)
	}
	return nil
}


// SetCurrencyPairFormat checks the exchange request and config currency pair
// formats and sets it to a default setting if it doesn't exist
func (e *Base) SetCurrencyPairFormat() error {
	cfg := config.GetConfig()
	exch, err := cfg.GetExchangeConfig(e.Name)
	if err != nil {
		return err
	}

	update := false
	if exch.RequestCurrencyPairFormat == nil {
		exch.RequestCurrencyPairFormat = &config.CurrencyPairFormatConfig{
			Delimiter: e.RequestCurrencyPairFormat.Delimiter,
			Uppercase: e.RequestCurrencyPairFormat.Uppercase,
			Separator: e.RequestCurrencyPairFormat.Separator,
			Index:     e.RequestCurrencyPairFormat.Index,
		}
		update = true
	} else {
		if CompareCurrencyPairFormats(e.RequestCurrencyPairFormat,
			exch.RequestCurrencyPairFormat) {
			e.RequestCurrencyPairFormat = *exch.RequestCurrencyPairFormat
		} else {
			*exch.RequestCurrencyPairFormat = e.ConfigCurrencyPairFormat
			update = true
		}
	}

	if exch.ConfigCurrencyPairFormat == nil {
		exch.ConfigCurrencyPairFormat = &config.CurrencyPairFormatConfig{
			Delimiter: e.ConfigCurrencyPairFormat.Delimiter,
			Uppercase: e.ConfigCurrencyPairFormat.Uppercase,
			Separator: e.ConfigCurrencyPairFormat.Separator,
			Index:     e.ConfigCurrencyPairFormat.Index,
		}
		update = true
	} else {
		if CompareCurrencyPairFormats(e.ConfigCurrencyPairFormat,
			exch.ConfigCurrencyPairFormat) {
			e.ConfigCurrencyPairFormat = *exch.ConfigCurrencyPairFormat
		} else {
			*exch.ConfigCurrencyPairFormat = e.ConfigCurrencyPairFormat
			update = true
		}
	}

	if update {
		return cfg.UpdateExchangeConfig(exch)
	}
	return nil
}

// CompareCurrencyPairFormats checks and returns whether or not the two supplied
// config currency pairs match
func CompareCurrencyPairFormats(pair1 config.CurrencyPairFormatConfig, pair2 *config.CurrencyPairFormatConfig) bool {
	if pair1.Delimiter != pair2.Delimiter ||
		pair1.Uppercase != pair2.Uppercase ||
		pair1.Separator != pair2.Separator ||
		pair1.Index != pair2.Index {
		return false
	}
	return true
}


// SetAssetTypes checks the exchange asset types (whether it supports SPOT,
// Binary or Futures) and sets it to a default setting if it doesn't exist
func (e *Base) SetAssetTypes() error {
	cfg := config.GetConfig()
	exch, err := cfg.GetExchangeConfig(e.Name)
	if err != nil {
		return err
	}

	update := false
	if exch.AssetTypes == "" {
		exch.AssetTypes = common.JoinStrings(e.AssetTypes, ",")
		update = true
	} else {
		e.AssetTypes = common.SplitStrings(exch.AssetTypes, ",")
	}

	if update {
		return cfg.UpdateExchangeConfig(exch)
	}

	return nil
}