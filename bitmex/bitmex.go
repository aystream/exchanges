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

	bitmexOrderBuy  = "Buy"
	bitmexOrderSell = "Sell"
	bitmexOrderType = "Limit"

	// OrderBook : Level 2 Book Data
	bitmexOrderBookL2 = "orderBook/L2"

	// Position : Summary of Open and Closed Positions
	bitmexPosition               = "position"
	bitmexPositionIsolate        = "position/isolate"
	bitmexPositionLeverage       = "position/leverage"
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
func (b *Bitmex) GetOrder(currencyPair string, orderId string, count int) (Order, error) {
	vals := url.Values{}
	vals.Set("symbol", currencyPair)
	vals.Set("count", strconv.Itoa(count))
	vals.Set("reverse", "true")

	var resp []Order
	var order Order

	err := b.SendAuthenticatedHTTPRequest("GET", bitmexOrder, vals, &resp)
	if err != nil {
		return order, err
	}

	for item := range resp {
		if resp[item].OrderID == orderId {
			order = resp[item]
			return order, nil
		}
	}

	return order, nil
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
 * @return open orders array
 */
func (b *Bitmex) GetOpenOrders(currencyPair string, count int) ([]Order, error) {
	vals := url.Values{}

	vals.Set("symbol", currencyPair)
	vals.Set("count", strconv.Itoa(count))
	vals.Set("reverse", "true")
	var resp []Order

	var openOrders []Order

	err := b.SendAuthenticatedHTTPRequest("GET", bitmexOrder, vals, &resp)

	if err != nil {
		return nil, err
	}

	for i := range resp {
		if resp[i].OrdStatus == "New" || resp[i].OrdStatus == "PartiallyFilled" {
			openOrders = append(openOrders, resp[i])
		}
	}
	return openOrders, err
}

/*
 * Get Open Positions
 *
 * Get all your open positions
 *
 * @return open positions array
 */
func (b *Bitmex) GetOpenPositions(currencyPair string) ([]Position, error) {
	vals := url.Values{}
	vals.Set("symbol", currencyPair)

	var resp []Position
	var openPositions []Position

	err := b.SendAuthenticatedHTTPRequest("GET", bitmexPosition, vals, &resp)

	if err != nil {
		return nil, err
	}

	for i := range resp {
		if resp[i].IsOpen == true {
			openPositions = append(openPositions, resp[i])
		}
	}
	return openPositions, err
}

/*
  * Close Position
  *
  * Close open position by Price.
  * If no price is specified, a market order will be submitted
  * to close the whole of your position
  *
  * @return array TODO checked return type!
  */
func (b *Bitmex) ClosePosition(currencyPair string, price float64) error {
	vals := url.Values{}
	vals.Set("symbol", currencyPair)
	if price != 0 {
		vals.Set("price", strconv.FormatFloat(price, 'E', -1, 64))
	}

	err := b.SendAuthenticatedHTTPRequest("GET", bitmexOrderClosePosition, vals, nil)

	if err != nil {
		return err
	}

	return nil
}

/*
  * Edit Order Price
  *
  * Edit you open order price
  *
  * @param orderID    Order ID
  * @param price      new price
  *
  * @return new order array TODO checked return type!
  */
func (b *Bitmex) EditOrderPrice(currencyPair string, price float64) error {
	vals := url.Values{}
	vals.Set("symbol", currencyPair)

	if price != 0 {
		return fmt.Errorf("Not %s Price for edit currency pair %s!", price, currencyPair)
	}

	vals.Set("price", strconv.FormatFloat(price, 'E', -1, 64))

	err := b.SendAuthenticatedHTTPRequest("PUT", bitmexOrder, vals, nil)

	if err != nil {
		return err
	}

	return nil
}

/*
 * Create Order
 *
 * Create new market order
 *
 * @param type can be "Limit"
 * @param side can be "Buy" or "Sell"
 * @param price BTC price in USD
 * @param quantity should be in USD (number of contracts)
 * @param maker forces platform to complete your order as a 'maker' only
 *
 * @return new order array TODO checked return type!
 */
func (b *Bitmex) CreateOrder(currencyPair string, typeOrder string, side string, price float64, quantity int, maker bool) error {
	vals := url.Values{}
	vals.Set("symbol", currencyPair)
	vals.Set("side", side)
	vals.Set("price", strconv.FormatFloat(price, 'E', -1, 64))
	vals.Set("orderQty", strconv.Itoa(quantity))
	vals.Set("ordType", typeOrder)

	if maker {
		vals.Set("execInst", "ParticipateDoNotInitiate")
	}

	err := b.SendAuthenticatedHTTPRequest("POST", bitmexOrder, vals, nil)

	if err != nil {
		return err
	}

	return nil
}

/*
 * Cancel All Open Orders
 *
 * Cancels all of your open orders
 *
 * @param $text is a note to all closed orders
 *
 * @return all closed orders arrays TODO checked return type!
 */
func (b *Bitmex) CancelAllOpenOrders(currencyPair string, text string) error {
	vals := url.Values{}
	vals.Set("symbol", currencyPair)
	vals.Set("text", text)

	err := b.SendAuthenticatedHTTPRequest("DELETE", bitmexOrderAll, vals, nil)

	if err != nil {
		return err
	}
	return nil
}

/*
  * Get Wallet
  *
  * Get your account wallet
  *
  * @return Wallet
  */
func (b *Bitmex) GetWallet() (Wallet, error) {
	vals := url.Values{}
	vals.Set("currency", "XBt")

	var userWallet Wallet

	err := b.SendAuthenticatedHTTPRequest("GET", bitmexUserWallet, vals, &userWallet)

	if err != nil {
		return userWallet, err
	}
	return userWallet, nil
}

/*
 * Get Margin
 *
 * Get your account margin
 *
 * @return Margin
 */
func (b *Bitmex) GetMarginInfo() (Margin, error) {
	vals := url.Values{}
	vals.Set("currency", "XBt")

	var userMargin Margin

	err := b.SendAuthenticatedHTTPRequest("GET", bitmexUserMargin, vals, &userMargin)

	if err != nil {
		return userMargin, err
	}
	return userMargin, nil
}

/*
 * Get Order Book
 *
 * Get L2 Order Book
 *
 * @return array
 */
func (b *Bitmex) GetOrderBookL2(currencyPair string, depth int) ([]OrderBookL2, error) {
	vals := url.Values{}
	vals.Set("symbol", currencyPair)
	vals.Set("depth", strconv.Itoa(depth))

	var orderBook []OrderBookL2

	err := b.SendAuthenticatedHTTPRequest("GET", bitmexOrderBookL2, vals, &orderBook)

	if err != nil {
		return orderBook, err
	}
	return orderBook, nil
}

/*
 * Set Leverage
 *
 * Set position leverage
 * $leverage = 0 for cross margin
 *
 * @return array
 */
func (b *Bitmex) SetLeverage(currencyPair string, leverage float64) (Position, error) {
	vals := url.Values{}
	vals.Set("symbol", currencyPair)
	vals.Set("leverage", strconv.FormatFloat(leverage, 'E', -1, 64))

	var position Position

	err := b.SendAuthenticatedHTTPRequest("POST", bitmexPositionLeverage, vals, &position)

	if err != nil {
		return position, err
	}
	return position, nil
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
