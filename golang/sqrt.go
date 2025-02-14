package sprint

func Sqrt(n int) int {
	if n < 0 || n == 0 {
		return 0
	}
	guess := n / 2
	for i := 0; i < 20; i++ {
		guess = (guess + n/guess) / 2
	}
	if guess*guess == n {
		return guess
	}
	return 0
}
