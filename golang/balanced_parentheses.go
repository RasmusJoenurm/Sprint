package sprint

func BalancedParentheses(input string) bool {
	stack := []rune{}
	matching := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	for _, char := range input {
		switch char {
		case '(', '{', '[':
			stack = append(stack, char)
		case ')', '}', ']':
			if len(stack) == 0 || stack[len(stack)-1] != matching[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
