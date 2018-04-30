package models

import (
	"time"
	"log"
)

// General candle when receiving data from exchange
type Candle struct {
	BidPrice  float64   `json:"bidPrice,omitempty"`
	AskPrice  float64   `json:"askPrice,omitempty"`
	Close     float64   `json:"close,omitempty"`
	Open      float64   `json:"open,omitempty"`
	High      float64   `json:"high,omitempty"`
	Low       float64   `json:"low,omitempty"`
	Volume    float64   `json:"volume,omitempty"`
	Timestamp time.Time `json:"timestamp"`
	Symbol    string    `json:"symbol"`
}

/*
	Serves to collect candles in a single data array,
	the strategy list of candles should be from some date to the current date!
*/
func GeneratedCandlesFromTimeframe(candles []Candle, currentTimeframe TimeFrame, timeframe TimeFrame) []Candle {
	var newGeneratedCandles []Candle

	if currentTimeframe > timeframe && timeframe%currentTimeframe != 0 {
		log.Printf("Genearte candles from current timeframe %s to %s timeframe, not work!",
			currentTimeframe, timeframe)
	} else if currentTimeframe == timeframe {
		// if suddenly it happened that the timeframes coincided
		return candles
	}

	// how many times do we have more time and how many candles will it take to collect them
	numberOfTimes := timeframe / currentTimeframe

	// Let's find the starting point from which candle to start counting
	firstIndexCandle := 0
	for index, candle := range candles {
		candle.Timestamp.Minute()
		//candle.Timestamp
		time := candle.Timestamp
		if time.Minute() == 0 {
			// this is zero time
			// with this candle should be the numbering of candles!
			firstIndexCandle = index
			break
		}
	}

	// Number of passes
	var numberPass = 0

	var open float64
	var close float64
	var height float64
	var low float64
	var volume float64
	var timestamp time.Time

	for _, candle := range candles[firstIndexCandle:] {

		if numberPass == 0 {
			// начало свечи или цена открытия будет тут!
			timestamp = candle.Timestamp
			open = candle.Open
			height = candle.High
			low = candle.Low
		}

		// search minimal value candle
		if low > candle.Low {
			low = candle.Low
		}
		// search maximum value candle
		if height < candle.High {
			height = candle.High
		}

		// sum volume
		volume += candle.Volume

		if numberPass == (int(numberOfTimes) - 1) {
			// this is close candle or price close
			close = candle.Close

			newCandle := Candle{AskPrice: close, BidPrice: close, High: height, Low: low, Open: open, Close: close, Symbol: candle.Symbol,
				Timestamp: timestamp, Volume: volume}
			// Add List in new generate candle
			newGeneratedCandles = append(newGeneratedCandles, newCandle)
			volume = 0
			numberPass = 0
		}

		numberPass++
	}

	return newGeneratedCandles
}