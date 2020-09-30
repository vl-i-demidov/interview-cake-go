package arrays

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyString(t *testing.T) {

	var expected []byte
	var actual []byte

	reverseArray(&actual)

	assert.Equal(t, expected, actual)

}

func TestSingleCharacter(t *testing.T) {

	expected := []byte("A")
	actual := []byte("A")

	reverseArray(&actual)

	assert.Equal(t, expected, actual)
}

func TestLongString(t *testing.T) {

	expected := []byte("EDCBA")
	actual := []byte("ABCDE")

	reverseArray(&actual)

	assert.Equal(t, expected, actual)
}

func TestSlice(t *testing.T) {
	var tests = []struct {
		str      []byte
		expected []byte
	}{
		{[]byte(""), []byte("")},
		{[]byte("A"), []byte("A")},
		{[]byte("EDCBA"), []byte("ABCDE")},
		{[]byte("FEDCBA"), []byte("ABCDEF")},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("%d", tt.str)
		t.Run(testname, func(t *testing.T) {
			ans := reverseSlice(tt.str)
			assert.Equal(t, tt.expected, ans)
		})
	}
}
