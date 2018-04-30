package bitmex

import "testing"

var b Bitmex

const SYMBOL = "XBTUSD"


func TestSetDefaults(t *testing.T) {
	b.SetDefaults()
}


func TestGetTicker(t *testing.T) {
	_, err := b.GetTicker(SYMBOL)
	if err != nil {
		t.Error("Test faild - Bitmex GetTicker() error")
	}
}
