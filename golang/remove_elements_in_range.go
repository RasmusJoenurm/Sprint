package sprint

func RemoveElementsInRange(arr []float64, from, to int) []float64 {
	if from > to {
		from, to = to, from
	}
	if from < 0 && to > len(arr)-1 {
		return []float64{}
	} else if to > len(arr)-1 {
		return arr[:from]
	} else if from >= 0 {
		return append(arr[:from], arr[to:]...)
	}
	return arr[to:]
}
