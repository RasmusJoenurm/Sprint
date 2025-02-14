package sprint

func IsLower2(s string) bool {
	for _, character := range s {
		if 'a' > character || character > 'z' {
			return false
		}
	}
	return true
}
