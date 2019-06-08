package candles

import "errors"

func BirthdayCakeCandles(candles []int) (int, error) {
	if len(candles) < 1 {
		return 0, errors.New("list of candles must not be empty")
	}

	tallestCandle := max(candles)
	tallesCandleCount := countCandles(candles, tallestCandle)
	return tallesCandleCount, nil
}

func max(candles []int) int {
	var max int

	for _, candle := range candles {
		if candle > max {
			max = candle
		}
	}

	return max
}

func countCandles(candles []int, tallest int) int {
	var candleCountsByHeight int

	for _, candleHeight := range candles {
		if candleHeight == tallest {
			candleCountsByHeight++
		}
	}

	return candleCountsByHeight
}
