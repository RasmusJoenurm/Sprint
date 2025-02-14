package sprint

func AlphaNumber(n int) string {

	IsNegative := n < 0
	if IsNegative {
		n = -n
	}
	result := ""
	if n == 0 {
		result = "a"
	} else {
		for n > 0 {
			digit := n % 10
			result = string(rune(digit+97)) + result
			n /= 10
		}
	}
	if IsNegative {
		result = "-" + result
	}
	return result
}
