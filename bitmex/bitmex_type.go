package bitmex

// Ticker ticker data
type Ticker struct {
	Symbol        string  `json:"symbol"`
	Last          float64 `json:"lastPrice"`
	Bid          float64 `json:"bidPrice"`
	Ask          float64 `json:"askPrice"`
	High          float64 `json:"highPrice"`
	low          float64 `json:"lowPrice"`
}
