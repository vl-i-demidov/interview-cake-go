package greedyalgs

import (
	"errors"
)

func highestProductOf3(nums []int) (int, error) {

	if len(nums) < 3 {
		return -1, errors.New("nums array must have at least 3 items")
	}

	vals := nums[:3]
	prod := nums[0] * nums[1] * nums[2]

	for i := 3; i < len(nums); i++ {
		pr := prod * nums[i]
		jv := -1
		max := prod
		for j := 0; j < 3; j++ {
			x := pr / vals[j]
			if x > max {
				max = x
				jv = j
			}
		}
		if jv != -1 {
			vals[jv] = nums[i]
			prod = max
		}
	}

	return prod, nil
}
