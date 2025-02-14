package sprint

func RemoveDuplicates(arr []int) []int {
	result := []int{}
	seen := map[int]bool{}

	for _, value := range arr {
		if !seen[value] {
			result = append(result, value)
			seen[value] = true
		}
	}
	return result
}
