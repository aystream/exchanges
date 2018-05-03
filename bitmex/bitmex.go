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
	"strconv"
	"github.com/CryptoTradingBot/exchanges/models"
)

const (
	bitmexAPIURL     = "https://www.bitmex.com/api/v1"
	bitmexAPIVersion = "v1"

	bitmexMaxOpenOrders = 200
	bitmexMaxOrderStop  = 10
	bitmexMaxOrderLimit = 10

	// APIKey : Persistent API Keys for Developers
	bimexAPIKey         = "APIKey"
	bitmexAPIKeyEnable  = "APIKey/enable"
	bitmexAPIKeyDisable = "APIKey/disable"

	// Chat : Trollbox Data
	bitmexChat          = "chat"
	bitmexChatChannels  = "chat/channels"
	bitmexChatConnected = "chat/connected"

	// Execution : Raw Order and Balance Data
	bitmexExecution    = "execution"
	bitmexTradeHistory = "execution/tradeHistory"

	// Funding : Swap Funding History
	bitmexFunding = "funding"

	// Instrument : Tradeable Contracts, Indices, and History
	bitmexInstrument                 = "instrument"
	bitmexInstrumentActive           = "instrument/active"
	bitmexInstrumentActiveAndIndices = "instrument/activeAndIndices"
	bitmexInstrumentActiveIntervals  = "instrument/activeIntervals"
	bitmexInstrumentCompositeIndex   = "instrument/compositeIndex"
	bitmexInstrumentIndices          = "instrument/indices"

	// Insurance : Insurance Fund Data
	bitmexInsurance = "insurance"

	// Leaderboard : Information on Top Users
	bitmexLeaderboard     = "leaderboard"
	bitmexLeaderboardName = "leaderboard/name"

	// Liquidation : Active Liquidations
	bitmexLiquidation = "liquidation"

	// Notification : Account Notifications
	bitmexNotification = "notification"

	// Order : Placement, Cancellation, Amending, and History
	bitmexOrder               = "order"
	bitmexOrderAll            = "order/all"
	bitmexOrderBulk           = "order/bulk"
	bitmexOrderCancelAllAfter = "order/cancelAllAfter"
	bitmexOrderClosePosition  = "order/closePosition"

	// OrderBook : Level 2 Book Data
	bitmexOrderBookL2 = "orderBook/L2"

	// Position : Summary of Open and Closed Positions
	bittmexPosition              = "position"
	bitmexPositionIsolate        = "position/isolate"
	bitmexPositionIeverage       = "position/ieverage"
	bitmexPositionRiskLimit      = "position/riskLimit"
	bitmexPositionTransferMargin = "position/transferMargin"

	// Quote : Best Bid/Offer Snapshots & Historical Bins
	bitmexQuote         = "quote"
	bitmexQuoteBucketed = "quote/bucketed"

	// Schema : Dynamic Schemata for Developers
	bitmexSchema              = "schema"
	bitmexSchemaWebsocketHelp = "schema/websocketHelp"

	// Settlement : Historical Settlement Data
	bitmexSettlement = "settlement"

	// Stats : Exchange Statistics
	bitmexStats           = "stats"
	bitmexStatsHistory    = "stats/history"
	bitmexStatsHistoryUSD = "stats/historyUSD"

	// Trade : Individual & Bucketed Trades
	bitmexTrade         = "trade"
	bitmexTradeBucketed = "trade/bucketed"

	// User : Account Operations
	bitmexUser                  = "user"
	bitmexUserAffiliateStatus   = "user/affiliateStatus"
	bitmexUserCancelWithdrawal  = "user/cancelWithdrawal"
	bitmexUserCheckReferralCode = "user/checkReferralCode"
	bitmexUserCommission        = "user/commission"
	bitmexUserConfirmEmail      = "user/confirmEmail"
	bitmexUserConfirmEnableTFA  = "user/confirmEnableTFA"
	bitmexUserConfirmWithdrawal = "user/confirmWithdrawal"
	bitmexUserDepositAddress    = "user/depositAddress"
	bitmexUserDisableTFA        = "user/disableTFA"
	bitmexUserLogout            = "user/logout"
	bitmexUserLogoutAll         = "user/logoutAll"
	bitmexUserMargin            = "user/margin"
	bitmexUserMinWithdrawalFee  = "user/minWithdrawalFee"
	bitmexUserPreferences       = "user/preferences"
	bitmexUserRequestEnableTFA  = "user/requestEnableTFA"
	bitmexUserRequestWithdrawal = "user/requestWithdrawal"
	bitmexUserWallet            = "user/wallet"
	bitmexUserWalletHistory     = "user/walletHistory"
	bitmexUserWalletSummary     = "user/walletSummary"

	// bitmex authenticated and unauthenticated limit rates
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
	b.Name = "Bitmex"
	b.Enabled = false
	b.Verbose = false
	b.Fee = 0
	b.Websocket = false
	b.RESTPollingDelay = 10
	b.SupportsAutoPairUpdating = true
	b.Handler = new(request.Handler)
	b.SetRequestHandler(b.Name, bitmexAuthRate, bitmexUnauthRate, new(http.Client))
}

/*
  * Get Ticker
  *
  * @return ticker array
  */
func (b *Bitmex) GetTicker(currencyPair string) ([]Ticker, error) {
	vals := url.Values{}

	if currencyPair != "" {
		vals.Set("symbol", currencyPair)
	}

	var resp []Ticker
	path := fmt.Sprintf("%s/%s?%s", bitmexAPIURL, bitmexInstrument, vals.Encode())

	err := b.SendHTTPRequest(path, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Get Candles history by current pair and timeFrame (can be 1m 5m 1h)
func (b *Bitmex) GetCandles(currencyPair string, timeframe string, count int) ([]models.Candle, error) {
	vals := url.Values{}

	vals.Set("symbol", currencyPair)
	vals.Set("count", strconv.Itoa(count))
	vals.Set("binSize", timeframe)
	vals.Set("partial", "false")
	vals.Set("reverse", "true")

	var resp []models.Candle
	path := fmt.Sprintf("%s/%s?%s", bitmexAPIURL, bitmexTradeBucketed, vals.Encode())

	err := b.SendHTTPRequest(path, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

/*
 * Get Order
 *
 * Get order by order ID
 *
 * @return array
 */
func (b *Bitmex) GetOrder(currencyPair string, orderId string, count int) ([]Order, error) {
	vals := url.Values{}

	vals.Set("symbol", currencyPair)
	vals.Set("count", strconv.Itoa(count))
	vals.Set("reverse", "true")
	if orderId != "" {
		vals.Set("filter", "{'orderID':'"+orderId+"'}")
	}

	var resp []Order

	err := b.SendAuthenticatedHTTPRequest("GET", bitmexOrder, vals, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

/*
 * Get Orders
 * @return array
 */
func (b *Bitmex) GetOrders(currencyPair string, count int) ([]Order, error) {
	vals := url.Values{}

	vals.Set("symbol", currencyPair)
	vals.Set("count", strconv.Itoa(count))
	vals.Set("reverse", "true")

	var resp []Order

	err := b.SendAuthenticatedHTTPRequest("GET", bitmexOrder, vals, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

/*
 * Get Open Orders
 *
 * Get open orders from the last 100 orders
 *
 * @return open orders array
 */
func GetOpenOrders(currencyPair string) () {
	vals := url.Values{}

	vals.Set("symbol", currencyPair)
	vals.Set("reverse", "true")

}

// SendHTTPRequest sends an unauthenticated HTTP request
func (b *Bitmex) SendHTTPRequest(path string, result interface{}) error {
	headers := make(map[string]string)
	headers["Connection"] = "Keep-Alive"
	headers["Keep-Alive"] = "90"
	return b.SendPayload("GET", path, headers, nil, result, false, b.Verbose)
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
	path := fmt.Sprintf("%s/%s", bitmexAPIURL, endpoint)

	return p.SendPayload(method, path, headers, bytes.NewBufferString(values.Encode()), result, true, p.Verbose)
}
