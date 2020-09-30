package arrays

func mergeArrays(a, b []int) []int {

	r := make([]int, len(a)+len(b))
	ai, bi, ri := 0, 0, 0

	for ri < len(r) {

		if bi == len(b) || (ai < len(a) && a[ai] < b[bi]) {
			r[ri] = a[ai]
			ai++
		} else {
			r[ri] = b[bi]
			bi++
		}

		ri++
	}

	return r
}
