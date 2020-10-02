package arrays

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsFirstComeFirstServed(t *testing.T) {
	var tests = []struct {
		a, b, o  []int
		expected bool
	}{
		// BothRegistersHaveSameNumberOfOrders
		{[]int{1, 4, 5}, []int{2, 3, 6}, []int{1, 2, 3, 4, 5, 6}, true},
		// RegistersHaveDifferentLengths
		{[]int{1, 5}, []int{2, 3, 6}, []int{1, 2, 6, 3, 5}, false},
		// OneRegisterIsEmpty
		{[]int{}, []int{2, 3, 6}, []int{2, 3, 6}, true},
		// ServedOrdersIsMissingOrders
		{[]int{1, 5}, []int{2, 3, 6}, []int{1, 6, 3, 5}, false},
		// ServedOrdersHasExtraOrders
		{[]int{1, 5}, []int{2, 3, 6}, []int{1, 2, 3, 5, 6, 8}, false},
		// OneRegisterHasExtraOrders
		{[]int{1, 9}, []int{7, 8}, []int{1, 7, 8}, false},
		// OneRegisterHasUnservedOrders
		{[]int{55, 9}, []int{7, 8}, []int{1, 7, 8, 9}, false},
		// OrderNumbersAreNotSequential
		{[]int{27, 12, 18}, []int{55, 31, 8}, []int{55, 31, 8, 27, 12, 18}, true},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("%d,%d,%d", tt.a, tt.b, tt.o)
		t.Run(testname, func(t *testing.T) {
			result := isFirstComeFirstServed(tt.a, tt.b, tt.o)
			assert.Equal(t, tt.expected, result)
		})
	}
}
