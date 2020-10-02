package arrays

func isFirstComeFirstServed(a, b, o []int) bool {

	ai, bi, oi := 0, 0, 0

	for oi < len(o) {

		if ai < len(a) && a[ai] == o[oi] {
			ai++
		} else if bi < len(b) && b[bi] == o[oi] {
			bi++
		} else {
			return false
		}
		oi++
	}

	return ai == len(a) && bi == len(b)
}
