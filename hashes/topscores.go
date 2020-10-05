package hashes

// we could use simple array, indexes instead of keys
func sortScores(us []int, max int) []int {
	dict := make(map[int]int)

	for _, score := range us {
		count, prs := dict[score]
		if prs {
			dict[score] = count + 1
		} else {
			dict[score] = 1
		}
	}
	sorted := make([]int, len(us))
	ind := 0
	for score := max; score >= 0; score-- {
		count, prs := dict[score]
		if prs {
			for j := 0; j < count; j++ {
				sorted[ind] = score
				ind++
			}
		}
	}

	return sorted
}
