package sprint

func ToUpperCase(s string) string {
	runes := []rune(s)
	for i, r := range runes {
		if r >= 'a' && r <= 'z' {
			runes[i] = r - ('a' - 'A')
		}
	}
	return string(runes)
}
