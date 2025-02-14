package sprint

func manualSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func Overlap(arr1, arr2 []int) []int {
	freq1 := make(map[int]int)
	freq2 := make(map[int]int)

	for _, num := range arr1 {
		freq1[num]++
	}

	for _, num := range arr2 {
		freq2[num]++
	}

	var result []int

	for num, count1 := range freq1 {
		if count2, found := freq2[num]; found {
			minCount := count1
			if count2 < minCount {
				minCount = count2
			}
			for i := 0; i < minCount; i++ {
				result = append(result, num)
			}
		}
	}
	result = manualSort(result)
	if len(result) == 0 {
		return []int{}
	}
	return result
}
