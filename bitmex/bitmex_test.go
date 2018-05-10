package bitmex

import (
	"testing"
	"github.com/CryptoTradingBot/exchanges/config"
)

var b Bitmex

var exchangeConfig config.ExchangeConfig

const (
	// fields for test verification bitmex methods
	Key       = "yHCVhvImDYXCyDVJCiJPjUDT"
	KeySecret = "b6LvJGLI3-cZJ5geVhRvccdE_w9y94opWVzaX5mGPReRJHMK"
	SYMBOL    = "XBTUSD"
)

func TestSetDefaults(t *testing.T) {
	b.SetDefaults()
}

func TestBitmex_Setup(t *testing.T) {
	exchangeConfig := config.ExchangeConfig{
		Verbose:true,
		AuthenticatedAPISupport:true,
		APIKey: Key,
		APISecret: KeySecret,
		UseSandbox: true,
	}

	b.Setup(exchangeConfig)
}

func TestBitmex_GetTicker(t *testing.T) {
	_, err := b.GetTicker(SYMBOL)
	if err != nil {
		t.Error("Test faild - Bitmex GetTicker() error")
	}
}

func TestBitmex_GetCandles(t *testing.T) {
	candles, err := b.GetCandles(SYMBOL, "5m", 100)
	if err != nil {
		t.Error("Test faild - Bitmex GetTicker() error")
	}
	t.Logf("%+v\n", candles)
}


func TestBitmex_CreateOrder(t *testing.T) {
	// Get Order book and test method at the first price
	orderBook, err := b.GetOrderBookL2(SYMBOL, 25)
	if err != nil {
		t.Error("Test faild - Bitmex GetOrderBookL2() error")
	}
	// Create limit order!
	order, err := b.CreateOrder(SYMBOL, "Limit", "Buy", orderBook[0].Price - 100, 100,  true)
	if err != nil {
		t.Error("Test faild - Bitmex CreateOrder() error")
	}
	t.Logf("%+v\n", order)
}



func TestBitmex_GetOrders(t *testing.T) {
	orders, err := b.GetOrders(SYMBOL, 100)
	if err != nil {
		t.Error("Test faild - Bitmex GetOrders() error")
	}
	t.Logf("%+v\n", orders)
}

func TestBitmex_GetOpenOrders(t *testing.T) {
	orders, err := b.GetOpenOrders(SYMBOL, 100)
	if err != nil {
		t.Error("Test faild - Bitmex GetOpenOrders() error")
	}
	t.Logf("%+v\n", orders)
}

func TestBitmex_GetOpenPositions(t *testing.T) {
	positions, err := b.GetOpenPositions(SYMBOL)
	if err != nil {
		t.Error("Test faild - Bitmex GetOpenPositions() error")
	}
	t.Logf("%+v\n", positions)
}

func TestBitmex_EditOrderPrice(t *testing.T) {
	// Get Order book and test method at the first price
	orderBook, err := b.GetOrderBookL2(SYMBOL, 25)
	if err != nil {
		t.Error("Test faild - Bitmex GetOrderBookL2() error for EditOrderPrice()")
	}

	// Get Order book and test method at the first price
	currentOrders, err := b.GetOrders(SYMBOL, 100)
	if err != nil {
		t.Error("Test faild - Bitmex GetOrders() error for EditOrderPrice()")
	}

	// Edit limit order!
	order, err := b.EditOrderPrice(currentOrders[0].OrderID, orderBook[0].Price)
	if err != nil {
		t.Error("Test faild - Bitmex EditOrderPrice() error")
	}
	t.Logf("%+v\n", order)
}

func TestBitmex_ClosePosition(t *testing.T) {
	order, err := b.ClosePosition(SYMBOL, 0)
	if err != nil {
		t.Error("Test faild - Bitmex ClosePosition() error")
	}
	t.Logf("%+v\n", order)
}

func TestBitmex_CancelAllOpenOrders(t *testing.T) {
	_,err := b.CancelAllOpenOrders(SYMBOL, "Cancel all position for " + SYMBOL)
	if err != nil {
		t.Error("Test faild - Bitmex CancelAllOpenOrders() error")
	}
}

func TestBitmex_GetWallet(t *testing.T) {
	wallet, err := b.GetWallet()
	if err != nil {
		t.Error("Test faild - Bitmex GetWallet() error")
	}
	t.Logf("%+v\n", wallet)
}

func TestBitmex_GetMarginInfo(t *testing.T) {
	margin, err := b.GetMarginInfo()
	if err != nil {
		t.Error("Test faild - Bitmex GetMarginInfo() error")
	}
	t.Logf("%+v\n", margin)
}

func TestBitmex_GetOrderBookL2(t *testing.T) {
	orderBook, err := b.GetOrderBookL2(SYMBOL, 25)
	if err != nil {
		t.Error("Test faild - Bitmex GetOrderBookL2() error")
	}
	t.Logf("%+v\n", orderBook)
}


func TestBitmex_SetLeverage(t *testing.T) {
	position, err := b.SetLeverage(SYMBOL, 10)
	if err != nil {
		t.Error("Test faild - Bitmex SetLeverage() error")
	}
	t.Logf("%+v\n", position)
}
