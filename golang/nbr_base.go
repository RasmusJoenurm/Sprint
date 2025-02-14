package sprint

func NbrBase(n int, base string) string {
	if len(base) < 2 || !allCharsUnique(base) || containsInvalidChars(base) {
		return "NV"
	}
	result := ""
	isNegative := n < 0
	if isNegative {
		n = -n
	}
	baseLen := len(base)
	for n >= baseLen {
		result = string(base[n%baseLen]) + result
		n = n / baseLen
	}
	result = string(base[n]) + result

	if isNegative {
		result = "-" + result
	}
	return result
}

func allCharsUnique(base string) bool {
	seen := map[rune]bool{}
	for _, char := range base {
		if seen[char] {
			return false
		}
		seen[char] = true
	}
	return true
}

func containsInvalidChars(base string) bool {
	for i := 0; i < len(base); i++ {
		if base[i] == '+' || base[i] == '-' {
			return true
		}
	}
	return false
}
