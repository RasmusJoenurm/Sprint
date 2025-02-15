package sprint

func IsPrime1(n int) bool {
	if n <= 0 {
		return false
	}

	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}
func NextPrime(n int) int {
	if n <= 1 {
		return 2
	}
	prime := n
	if !IsPrime(prime) {
		prime++
		for !IsPrime(prime) {
			prime++
		}
	}
	return prime
}
