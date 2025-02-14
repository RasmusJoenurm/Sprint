package sprint

func ArrAny(f func(string) bool, a []string) bool {
	for _, v := range a {
		if f(v) {
			return true
		}
	}
	return false
}

// IsLower returns true if the string contains only lowercase characters, and false otherwise.
func IsLower1(s string) bool {
	for _, character := range s {
		if 'a' > character || character > 'z' {
			return false
		}
	}
	return true
}

// IsNumeric returns true if the string contains only numeric characters, and false otherwise.
func IsNumeric1(s string) bool {
	for _, character := range s {
		if character < '0' || character > '9' {
			return false
		}
	}
	return true
}

// IsUpper returns true if the string contains only uppercase characters, and false otherwise.
func IsUpper1(s string) bool {
	for _, c := range s {
		if !('A' <= c && c <= 'Z') {
			return false
		}
	}
	return true
}

// IsAlphanumeric returns true if the string contains only alphanumeric characters, and false otherwise.
func IsAlphanumeric1(s string) bool {
	for _, c := range s {
		if !('A' <= c && c <= 'Z') && !('a' <= c && c <= 'z') && !('0' <= c && c <= '9') {
			return false
		}
	}
	return true
}
