package sprint

func CountDivisible(from, to, step, divisor int) int {
	if step <= 0 {
		return 0
	}
	if divisor == 0 {
		return 0
	}
	count := 0
	for i := from; i < to; i += step { //i to the value of from; loop continues until i is less than 10; increments i by step
		if i%divisor == 0 { //checks if i is divisable by divisor
			count++ //count +1
		}
	}
	return count
}
