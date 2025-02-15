package sprint

func LCM(a, b int) int {
	return (a * b) / GCD(a, b)
}

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
