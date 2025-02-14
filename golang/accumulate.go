package sprint

func Accumulate(n int) int {

	var sum int
	for i := 1; i <= n; i++ {
		sum += i
	}
	if n < 0 {
		return 0
	} else {
		return sum
	}
}
