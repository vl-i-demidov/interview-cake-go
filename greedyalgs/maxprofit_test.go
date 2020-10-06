package greedyalgs

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortScores(t *testing.T) {
	var tests = []struct {
		prices    []int
		maxProfit int
		err       error
	}{
		// PriceGoesUpThenDownTest
		{[]int{1, 5, 3, 2}, 4, nil},
		// PriceGoesDownThenUpTest
		{[]int{7, 2, 8, 9}, 7, nil},
		// PriceGoesUpAllDayTest
		{[]int{1, 6, 7, 9}, 8, nil},
		// PriceGoesDownAllDayTest
		{[]int{9, 7, 4, 1}, -2, nil},
		// PriceStaysTheSameAllDayTest
		{[]int{1, 1, 1, 1}, 0, nil},
		// ExceptionWithOnePriceTest
		{[]int{5}, -1, errors.New("")},
		// ExceptionWithEmptyPricesTest
		{[]int{}, -1, errors.New("")},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("%d", tt.prices)
		t.Run(testname, func(t *testing.T) {
			maxProfit, err := getMaxProfit(tt.prices)
			assert.Equal(t, tt.maxProfit, maxProfit)
			if tt.err != nil {
				assert.Error(t, err)
			}
		})
	}
}
