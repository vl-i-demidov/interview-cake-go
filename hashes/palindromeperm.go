package hashes

func hasPalindromePermutation(s string) bool {
	set := make(map[rune]struct{})
	var es struct{}
	for _, c := range s {
		_, prs := set[c]
		if prs {
			delete(set, c)
		} else {
			set[c] = es
		}
	}

	return len(set) < 2
}
