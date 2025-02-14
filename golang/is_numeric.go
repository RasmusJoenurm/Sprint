package sprint

func IsNumeric2(s string) bool {

	for _, character := range s {
		if character < '0' || character > '9' {
			return false
		}
	}
	return true
}
