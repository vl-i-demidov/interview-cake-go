package hashes

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasPalindromePermutation(t *testing.T) {
	var tests = []struct {
		str      string
		expected bool
	}{
		// PermutationWithOddNumberOfCharsTest
		{"aabcbcd", true},
		// PermutationWithEvenNumberOfCharsTest
		{"aabccbdd", true},
		// NoPermutationWithOddNumberOfChars
		{"aabcd", false},
		// NoPermutationWithEvenNumberOfCharsTest
		{"aabbcd", false},
		// EmptyStringTest
		{"", true},
		// OneCharacterStringTest
		{"a", true},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf(tt.str)
		t.Run(testname, func(t *testing.T) {
			res := hasPalindromePermutation(tt.str)
			assert.Equal(t, tt.expected, res)
		})
	}
}
