package bitmex

import (
	exchange "github.com/CryptoTradingBot/exchanges"
	"net/url"
	"fmt"
	"time"
	"github.com/CryptoTradingBot/exchanges/request"
	"github.com/CryptoTradingBot/exchanges/common"
	"bytes"
	"net/http"
)


const (
	bitmexAPIURL             = "https://www.bitmex.com"
	bitmexAPITradingEndpoint = "/api/v1/"


	// bitme authenticated and unauthenticated limit rates
	bitmexAuthRate   = 1000
	bitmexUnauthRate = 1000
)


// Bitmex is the overacting type across the bitmex methods
type Bitmex struct {
	exchange.Base
	*request.Handler
}


// SetDefaults sets the basic defaults for bitmex
func (b *Bitmex) SetDefaults() {
	b.Name = "Bitfinex"
	b.Enabled = false
	b.Verbose = false
	b.Websocket = false
	b.RESTPollingDelay = 10
	b.SupportsAutoPairUpdating = true
	b.Handler = new(request.Handler)
	b.SetRequestHandler(b.Name, bitmexAuthRate, bitmexUnauthRate, new(http.Client))
}


// SendHTTPRequest sends an unauthenticated HTTP request
func (b *Bitmex) SendHTTPRequest(path string, result interface{}) error {
	return b.SendPayload("GET", path, nil, nil, result, false, b.Verbose)
}

// SendAuthenticatedHTTPRequest sends an authenticated HTTP request
func (p *Bitmex) SendAuthenticatedHTTPRequest(method, endpoint string, values url.Values, result interface{}) error {
	if !p.AuthenticatedAPISupport {
		return fmt.Errorf(exchange.WarningAuthenticatedRequestWithoutCredentialsSet, p.Name)
	}
	headers := make(map[string]string)
	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["Connection"] = "Keep-Alive"
	headers["Keep-Alive"] = "90"

	if p.Nonce.Get() == 0 {
		p.Nonce.Set(time.Now().UnixNano())
	} else {
		p.Nonce.Inc()
	}

	headers["api-key"] = p.APIKey
	headers["api-nonce"] = p.Nonce.String()
	headers["api-signature"] = p.Nonce.String()

	values.Set("nonce", p.Nonce.String())
	values.Set("command", endpoint)

	hmac := common.GetHMAC(common.HashSHA256, []byte(values.Encode()), []byte(p.APISecret))
	headers["api-signature"] = common.HexEncodeToString(hmac)

	path := fmt.Sprintf("%s/%s", bitmexAPIURL, bitmexAPITradingEndpoint)

	return p.SendPayload(method, path, headers, bytes.NewBufferString(values.Encode()), result, true, p.Verbose)
}
