package sprint

func ConvertAnyToAny(nbr, baseFrom, baseTo string) string {
	convertToDecimal := func(s, base string) int {
		baseMap := make(map[rune]int)
		for i, char := range base {
			baseMap[char] = i
		}

		result := 0
		baseLen := len(base)
		for _, char := range s {
			value, exists := baseMap[char]
			if !exists {
				return 0
			}
			result = result*baseLen + value
		}
		return result
	}
	convertFromDecimal := func(n int, base string) string {
		if n == 0 {
			return string(base[0])
		}

		baseLen := len(base)
		result := ""
		for n > 0 {
			result = string(base[n%baseLen]) + result
			n /= baseLen
		}
		return result
	}
	decimalValue := convertToDecimal(nbr, baseFrom)
	return convertFromDecimal(decimalValue, baseTo)
}
