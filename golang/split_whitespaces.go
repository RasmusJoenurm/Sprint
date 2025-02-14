package sprint

func SplitWhitespaces(s string) []string {
	var words []string
	word := ""
	for _, char := range s {
		if char != ' ' && char != '\t' && char != '\n' {
			word += string(char)
		} else {
			if word != "" {
				words = append(words, word)
				word = ""
			}
		}
	}
	if word != "" {
		words = append(words, word)
	}
	return words
}
