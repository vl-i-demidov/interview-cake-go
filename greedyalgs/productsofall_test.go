package greedyalgs

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetProductsOfAll(t *testing.T) {
	tests := []struct {
		nums     []int
		products []int
		err      bool
	}{
		// SmallArrayInputTest
		{[]int{1, 2, 3}, []int{6, 3, 2}, false},
		// LongArrayInputTest
		{[]int{8, 2, 4, 3, 1, 5}, []int{120, 480, 240, 320, 960, 192}, false},
		// InputHasOneZeroTest
		{[]int{6, 2, 0, 3}, []int{0, 0, 36, 0}, false},
		// InputHasTwoZerosTest
		{[]int{4, 0, 9, 1, 0}, []int{0, 0, 0, 0, 0}, false},
		// InputHasOneNegativeNumberTest
		{[]int{-3, 8, 4}, []int{32, -12, -24}, false},
		// AllNegativesInputTest
		{[]int{-7, -1, -4, -2}, []int{-8, -56, -14, -28}, false},
		// ExceptionWithEmptyInputTest
		{[]int{}, nil, true},
		// ExceptionWithOneNumberInputTest
		{[]int{1}, nil, true},
	}
	for _, tt := range tests {

		testname := fmt.Sprintf("%d", tt.nums)
		t.Run(testname, func(t *testing.T) {
			products, err := getProductsOfAll(tt.nums)
			assert.Equal(t, tt.products, products)
			if tt.err {
				assert.Error(t, err)
			}
		})
	}
}
