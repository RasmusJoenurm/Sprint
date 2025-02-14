package sprint

func StrCompare1(a, b string) int {
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
