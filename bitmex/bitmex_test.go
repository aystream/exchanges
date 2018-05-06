package bitmex

import (
	"testing"
	"fmt"
)

var b Bitmex

const (
	Login = "caj10958@soioa.com"
	Password = "o9p0[-]="
	Key = "hD6WCCUD9O_dyAoOQNgW8oyz"
	KeySecret = "gzu1KMGGwsNIQFzGXdY9rk0ZGJ8w011BqRYVwJoAFBOKT4HD"
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
	fmt.Print(candles)
}

func TestBitmex_GetOrders(t *testing.T) {
	orders, err := b.GetOrders(SYMBOL, 100)
	if err != nil {
		t.Error("Test faild - Bitmex GetOrders() error")
	}
	fmt.Print(orders)
}

func TestBitmex_GetOpenOrders(t *testing.T) {
	orders, err := b.GetOpenOrders(SYMBOL, 100)
	if err != nil {
		t.Error("Test faild - Bitmex GetOrders() error")
	}
	fmt.Print(orders)
}

func TestBitmex_GetOpenPositions(t *testing.T) {

}

func TestBitmex_GetMarginInfo(t *testing.T) {
	margin, err := b.GetMarginInfo()
	if err != nil {
		t.Error("Test faild - Bitmex GetMarginInfo() error")
	}
	fmt.Print(margin)
}