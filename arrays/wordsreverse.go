package arrays

// only ASCII
func reverseWords(a []byte) {

	reverse(a, 0, len(a)-1)

	start, cur := 0, 0

	for cur <= len(a) {
		if cur == len(a) || string(a[cur]) == " " {
			reverse(a, start, cur-1)
			start = cur + 1
		}
		cur++
	}
}

func reverse(a []byte, left, right int) {
	for left < right {
		temp := a[left]
		a[left] = a[right]
		a[right] = temp
		left++
		right--
	}
}
