package arrays

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultipleWordsDifferentLengths(t *testing.T) {
	expected := []byte("chocolate bundt cake is yummy")
	actual := []byte("yummy is cake bundt chocolate")
	reverseWords(actual)
	assert.Equal(t, expected, actual)
}

func TestReverseWords(t *testing.T) {
	var tests = []struct {
		expected []byte
		str      []byte
	}{
		// one word
		{[]byte("vault"), []byte("vault")},
		// two words
		{[]byte("cake thief"), []byte("thief cake")},
		// three words
		{[]byte("get another one"), []byte("one another get")},
		// multiple words same length
		{[]byte("the cat ate the rat"), []byte("rat the ate cat the")},
		// multiple words different length
		{[]byte("chocolate bundt cake is yummy"), []byte("yummy is cake bundt chocolate")},
		// empty string
		{[]byte(""), []byte("")},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("%d", tt.str)
		t.Run(testname, func(t *testing.T) {
			reverseWords(tt.str)
			assert.Equal(t, tt.expected, tt.str)
		})
	}
}
