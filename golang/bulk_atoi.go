package sprint

func strToInt(s string) int {

	if len(s) == 0 {
		return 0
	}
	sign := 1
	start := 0
	if s[0] == '-' {
		start = 1
		sign = -1
	}
	if s[0] == '+' {
		start = 1
	}
	result := 0
	for i := start; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0
		}
		result = result*10 + int(s[i]-'0')
	}
	return result * sign
}

func BulkAtoi(arr []string) []int {
	valid := []int{}
	for _, element := range arr {
		valid = append(valid, strToInt(element))
	}
	return valid
}
