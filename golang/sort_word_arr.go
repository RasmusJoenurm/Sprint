package sprint

func SortWordArr(a []string) []string {
	quicksort(a, 0, len(a)-1)
	return a
}

func quicksort(a []string, low, high int) {
	if low < high {
		pi := partition(a, low, high)
		quicksort(a, low, pi-1)
		quicksort(a, pi+1, high)
	}
}

func partition(a []string, low, high int) int {
	pivot := a[high]
	i := low - 1
	for j := low; j < high; j++ {
		if a[j] < pivot {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}
	a[i+1], a[high] = a[high], a[i+1]
	return i + 1
}
