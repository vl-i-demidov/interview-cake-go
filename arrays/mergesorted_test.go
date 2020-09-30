package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBothArraysHaveSomeNumbers(t *testing.T) {
	a, b := []int{2, 4, 6}, []int{1, 3, 7}
	exp := []int{1, 2, 3, 4, 6, 7}

	actual := mergeArrays(a, b)

	assert.Equal(t, exp, actual)
}

func TestArraysAreDifferentLengths(t *testing.T) {
	a, b := []int{2, 4, 6, 8}, []int{1, 7}
	exp := []int{1, 2, 4, 6, 7, 8}

	actual := mergeArrays(a, b)

	assert.Equal(t, exp, actual)
}

func TestSecondArrayIsEmpty(t *testing.T) {
	a, b := []int{1, 2, 3}, []int{}
	exp := []int{1, 2, 3}

	actual := mergeArrays(a, b)

	assert.Equal(t, exp, actual)
}

func TestFirstArrayIsEmpty(t *testing.T) {
	a, b := []int{}, []int{5, 6, 7}
	exp := []int{5, 6, 7}

	actual := mergeArrays(a, b)

	assert.Equal(t, exp, actual)
}

func TestBothArraysAreEmptyTest(t *testing.T) {
	a, b := []int{}, []int{}
	exp := []int{}

	actual := mergeArrays(a, b)

	assert.Equal(t, exp, actual)
}
