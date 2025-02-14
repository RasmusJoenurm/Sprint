package sprint

import (
	"math"
	"strings"
)

func ConvertAnyToDec(s string, base string) int {
	if len(base) < 2 || strings.Contains(base, "+") || strings.Contains(base, "-") {
		return 0
	}
	baseMap := make(map[rune]int)
	for i, char := range base {
		baseMap[char] = i
	}
	result := 0
	baseLen := len(base)
	for i, char := range s {
		value, exists := baseMap[char]
		if !exists {
			return 0
		}
		result += value * int(math.Pow(float64(baseLen), float64(len(s)-i-1)))
	}
	return result
}
