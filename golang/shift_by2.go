package sprint

func ShiftBy2(r rune, step int) rune {
	var answer rune
	if r+rune(step)%26 > 122 {
		answer = r + rune(step)%26 - 26
	} else {
		answer = r + rune(step)%26
	}
	return answer
}
