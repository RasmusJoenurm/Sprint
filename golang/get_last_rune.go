package sprint

func GetLastRune(s string) rune {
	lastOne := s[len(s)-1]
	return rune(lastOne)
}
