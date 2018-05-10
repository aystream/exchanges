package config

import (
	"time"
	"fmt"
	"sync"
)

var (
	ErrExchangeNotFound = "Exchange %s: Not found."
	Cfg                 Config
)

// Config is the overarching object that holds all the information for
// prestart management of portfolio, SMSGlobal, webserver and enabled exchange
type Config struct {
	Name                     string
	EncryptConfig            int
	Cryptocurrencies         string
	CurrencyExchangeProvider string
	CurrencyPairFormat       *CurrencyPairFormatConfig `json:"CurrencyPairFormat"`
	FiatDisplayCurrency      string
	GlobalHTTPTimeout        time.Duration
	Exchanges                []ExchangeConfig          `json:"Exchanges"`
	m                        sync.Mutex
}

// CurrencyPairFormatConfig stores the users preferred currency pair display
type CurrencyPairFormatConfig struct {
	Uppercase bool
	Delimiter string `json:",omitempty"`
	Separator string `json:",omitempty"`
	Index     string `json:",omitempty"`
}

// ExchangeConfig holds all the information needed for each enabled Exchange.
type ExchangeConfig struct {
	Name                      string
	Enabled                   bool
	Verbose                   bool
	Websocket                 bool
	UseSandbox                bool
	RESTPollingDelay          time.Duration
	HTTPTimeout               time.Duration
	AuthenticatedAPISupport   bool
	APIKey                    string
	APISecret                 string
	ClientID                  string                    `json:",omitempty"`
	AvailablePairs            string
	EnabledPairs              string
	BaseCurrencies            string
	AssetTypes                string
	SupportsAutoPairUpdates   bool
	PairsLastUpdated          int64                     `json:",omitempty"`
	ConfigCurrencyPairFormat  *CurrencyPairFormatConfig `json:"ConfigCurrencyPairFormat"`
	RequestCurrencyPairFormat *CurrencyPairFormatConfig `json:"RequestCurrencyPairFormat"`
}

// GetExchangeConfig returns exchange configurations by its indivdual name
func (c *Config) GetExchangeConfig(name string) (ExchangeConfig, error) {
	c.m.Lock()
	defer c.m.Unlock()
	for i := range c.Exchanges {
		if c.Exchanges[i].Name == name {
			return c.Exchanges[i], nil
		}
	}
	return ExchangeConfig{}, fmt.Errorf(ErrExchangeNotFound, name)
}

// UpdateExchangeConfig updates exchange configurations
func (c *Config) UpdateExchangeConfig(e ExchangeConfig) error {
	c.m.Lock()
	defer c.m.Unlock()
	for i := range c.Exchanges {
		if c.Exchanges[i].Name == e.Name {
			c.Exchanges[i] = e
			return nil
		}
	}
	return fmt.Errorf(ErrExchangeNotFound, e.Name)
}

// GetConfig returns a pointer to a configuration object
func GetConfig() *Config {
	return &Cfg
}
