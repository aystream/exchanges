package bitmex

import (
	"testing"
)

var b Bitmex

const (
	Login = "caj10958@soioa.com"
	Password = "o9p0[-]="
	Key = "yHCVhvImDYXCyDVJCiJPjUDT"
	KeySecret = "b6LvJGLI3-cZJ5geVhRvccdE_w9y94opWVzaX5mGPReRJHMK"
	SYMBOL = "XBTUSD"
)

func TestSetDefaults(t *testing.T) {
	b.SetDefaults()
}

func TestBitmex_Setup(t *testing.T) {
	b.Setup(Key, KeySecret,true)
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
	t.Logf("%+v\n",candles)
}

func TestBitmex_GetOrders(t *testing.T) {
	orders, err := b.GetOrders(SYMBOL, 100)
	if err != nil {
		t.Error("Test faild - Bitmex GetOrders() error")
	}
	t.Logf("%+v\n",orders)
}

func TestBitmex_GetOpenOrders(t *testing.T) {
	orders, err := b.GetOpenOrders(SYMBOL, 100)
	if err != nil {
		t.Error("Test faild - Bitmex GetOrders() error")
	}
	t.Logf("%+v\n",orders)
}

func TestBitmex_GetOpenPositions(t *testing.T) {

}

func TestBitmex_GetMarginInfo(t *testing.T) {
	margin, err := b.GetMarginInfo()
	if err != nil {
		t.Error("Test faild - Bitmex GetMarginInfo() error")
	}
	t.Logf("%+v\n",margin)
}

func TestBitmex_SetLeverage(t *testing.T) {
	position, err := b.SetLeverage(SYMBOL, 10)
	if err != nil {
		t.Error("Test faild - Bitmex SetLeverage() error")
	}
	t.Logf("%+v\n", position)
}