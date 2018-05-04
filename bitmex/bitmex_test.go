package bitmex

import (
	"testing"
	"fmt"
)

var b Bitmex

const SYMBOL = "XBTUSD"

func TestSetDefaults(t *testing.T) {
	b.SetDefaults()
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



