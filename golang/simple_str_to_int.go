package sprint

func SimpleStrToInt(s string) int {

	if len(s) == 0 {
		return 0
	}
	result := 0
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0
		}
		result = result*10 + int(s[i]-'0')
	}
	return result
}
