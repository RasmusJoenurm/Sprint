package sprint

func LongestClimb(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}
	maxLength := 1
	currentLength := 1
	startIndex := 0
	maxStartIndex := 0

	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[i-1] {
			currentLength++
		} else {
			if currentLength > maxLength {
				maxLength = currentLength
				maxStartIndex = startIndex
			}
			currentLength = 1
			startIndex = i
		}
	}
	if currentLength > maxLength {
		maxLength = currentLength
		maxStartIndex = startIndex
	}
	return arr[maxStartIndex : maxStartIndex+maxLength]
}
