package sprint

func FactorialIterative(n int) int {
	if n < 0 {
		return 0
	}
	if n == 0 || n == 1 {
		return 1
	}
	result := 1
	for i := 2; i <= n; i++ {
		if result > (1<<31-1)/i {
			return 0
		}
		result *= i
	}
	return result

}
