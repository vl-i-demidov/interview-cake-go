package hashes

func canTwoMoviesFillFlight(movieLengths []int, flightLength int) bool {

	set := make(map[int]bool)

	for _, len := range movieLengths {
		_, prs := set[len]
		if prs {
			return true
		}
		set[flightLength-len] = true
	}
	return false
}
