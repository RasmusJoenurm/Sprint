package sprint

func FilterBySum(arr [][]int, limit int) [][]int {
	valid := [][]int{}
	for _, subArray := range arr {
		sum := 0
		for _, number := range subArray {
			sum += number
		}
		if sum >= limit {
			valid = append(valid, subArray)
		}
	}
	return valid
}
