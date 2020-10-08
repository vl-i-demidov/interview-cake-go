package greedyalgs

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHighestProductOf3(t *testing.T) {
	var tests = []struct {
		nums []int
		prod int
		err  bool
	}{
		{[]int{1, 2, 3, 4}, 24, false},
		{[]int{6, 1, 3, 5, 7, 8, 2}, 336, false},
		{[]int{-5, 4, 8, 2, 3}, 96, false},
		{[]int{-10, 1, 3, 2, -10}, 300, false},
		{[]int{-5, -1, -3, -2}, -6, false},
		{[]int{}, -1, true},
		{[]int{1}, -1, true},
		{[]int{1, 1}, -1, true},
	}
	for _, tt := range tests {

		testname := fmt.Sprintf("%d", tt.nums)
		t.Run(testname, func(t *testing.T) {
			prod, err := highestProductOf3(tt.nums)
			assert.Equal(t, tt.prod, prod)
			if tt.err {
				assert.Error(t, err)
			}
		})
	}
}
