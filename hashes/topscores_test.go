package hashes

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortScores(t *testing.T) {
	var tests = []struct {
		unsorted []int
		sorted   []int
		max      int
	}{
		// NoScoresTest
		{[]int{}, []int{}, 100},
		// OneScoreTest
		{[]int{55}, []int{55}, 100},
		// TwoScoresTest
		{[]int{30, 60}, []int{60, 30}, 100},
		// ManyScoresTest
		{[]int{37, 89, 41, 65, 91, 53}, []int{91, 89, 65, 53, 41, 37}, 100},
		// RepeatedScoresTest
		{[]int{20, 10, 30, 30, 10, 20}, []int{30, 30, 20, 20, 10, 10}, 100},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("%d", tt.unsorted)
		t.Run(testname, func(t *testing.T) {
			res := sortScores(tt.unsorted, tt.max)
			assert.Equal(t, tt.sorted, res)
		})
	}
}
