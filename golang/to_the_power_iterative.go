package sprint

func ToThePowerIterative(n int, power int) int {
	if power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	}
	if n == 1 {
		return 1
	}
	if n == 0 {
		return 0
	}
	result := 1
	for i := 1; i <= power; i++ {
		if result > (1<<31-1)/i {
			return 0
		}
		result *= n
	}
	return result

}
