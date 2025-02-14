package sprint

func IsSorted(f func(a, b string) int, arr []string) bool {
	n := len(arr)
	if n < 2 {
		return true
	}

	ascending := true
	descending := true

	for i := 1; i < n; i++ {
		if f(arr[i-1], arr[i]) > 0 {
			ascending = false
		}
		if f(arr[i-1], arr[i]) < 0 {
			descending = false
		}
	}

	return ascending || descending
}

func StrCompare(a, b string) int {
	aLen := len(a)
	bLen := len(b)
	minLen := aLen
	if bLen < minLen {
		minLen = bLen
	}
	for i := 0; i < minLen; i++ {
		if a[i] < b[i] {
			return -1
		} else if a[i] > b[i] {
			return 1
		}
	}
	if aLen < bLen {
		return -1
	} else if aLen > bLen {
		return 1
	}
	return 0
}
