package sprint

func Countdown(n int) string {
	if n <= 0 || n > 9 {
		return "0!"
	}
	result := ""

	for i := n; i > 0; i -= 2 {
		if result != "" {
			result += ", "
		}
		result += string(rune(i) + 48)
	}
	result += ", 0!"
	return result
}
