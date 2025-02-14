package sprint

func ArrCountIf(f func(string) bool, a []string) int {
	count := 0
	for _, v := range a {
		if f(v) {
			count++
		}
	}
	return count
}

func IsLower(s string) bool {
	for _, character := range s {
		if 'a' > character || character > 'z' {
			return false
		}
	}
	return true
}

func IsUpper(s string) bool {
	for _, c := range s {
		if !('A' <= c && c <= 'Z') {
			return false
		}
	}
	return true
}

func IsNumeric(s string) bool {

	for _, character := range s {
		if character < '0' || character > '9' {
			return false
		}
	}
	return true
}

func IsAlphanumeric(s string) bool {
	for _, c := range s {
		if !('A' <= c && c <= 'Z') && !('a' <= c && c <= 'z') && !('0' <= c && c <= '9') {
			return false
		}
	}
	return true
}
