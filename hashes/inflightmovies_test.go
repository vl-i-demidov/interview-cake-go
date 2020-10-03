package hashes

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanTwoMoviesFillFlight(t *testing.T) {
	var tests = []struct {
		movies   []int
		flight   int
		expected bool
	}{
		// ShortFlightTest
		{[]int{2, 4}, 1, false},
		// LongFlightTest
		{[]int{2, 4}, 6, true},
		// OnlyOneMovieHalfFlightLenghtTest
		{[]int{3, 8}, 6, false},
		// TwoMoviesHalfFlightLengthTest
		{[]int{3, 8, 3}, 6, true},
		// LotsOfPossiblePairsTest
		{[]int{1, 2, 3, 4, 5, 6}, 7, true},
		// NotUsingFirstMovieTest
		{[]int{4, 3, 2}, 5, true},
		// OneMovieTest
		{[]int{6}, 6, false},
		// NoMoviesTest
		{[]int{}, 6, false},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("%d", tt.movies)
		t.Run(testname, func(t *testing.T) {
			res := canTwoMoviesFillFlight(tt.movies, tt.flight)
			assert.Equal(t, tt.expected, res)
		})
	}
}
