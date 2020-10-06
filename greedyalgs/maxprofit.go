package greedyalgs

import (
	"errors"
	"math"
)

// TODO there is a simpler version
// without special case for negative sum
func getMaxProfit(prices []int) (int, error) {

	if len(prices) < 2 {
		return -1, errors.New("Array must have at least 2 items")
	}

	sum := math.MinInt32

	for i := 1; i < len(prices); i++ {
		curProfit := prices[i] - prices[i-1]
		if curProfit > 0 && sum >= 0 {
			sum += curProfit
		} else {
			if curProfit > sum {
				sum = curProfit
			}
		}
	}

	return sum, nil
}
