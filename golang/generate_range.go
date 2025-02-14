package sprint

func GenerateRange(min, max int) []int {
	// if max smaller or equal to min = return nil
	if min >= max {
		return nil
	}
	//figure out the length and make a variable length
	length := max - min
	// create a slice with the calculated length
	result := make([]int, length)

	for i := 0; i < length; i++ {
		result[i] = min + i
	}
	return result
}
