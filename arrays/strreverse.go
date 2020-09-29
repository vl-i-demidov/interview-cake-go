package arrays

func reverseArray(ap *[]byte) *[]byte {
	a := *ap
	left := 0
	right := len(a) - 1

	for left < right {
		temp := a[left]
		a[left] = a[right]
		a[right] = temp
		left++
		right--
	}
	return &a
}

func reverseSlice(a []byte) []byte {
	left := 0
	right := len(a) - 1

	for left < right {
		temp := a[left]
		a[left] = a[right]
		a[right] = temp
		left++
		right--
	}
	return a
}
