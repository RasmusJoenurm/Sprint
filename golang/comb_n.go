package sprint

import (
	"fmt"
)

func CombN(n int) []string {
	var outcome []string
	generate(n, "", &outcome)
	return outcome
}

func generate(n int, currentNumber string, outcome *[]string) {
	if n == 0 {
		*outcome = append(*outcome, currentNumber)
		return
	}
	lastDigit := 0
	if len(currentNumber) > 0 {
		lastDigit = int(currentNumber[len(currentNumber)-1] - '0')
	}
	for nextDigit := lastDigit; nextDigit <= 9; nextDigit++ {
		nextNumber := currentNumber + fmt.Sprintf("%d", nextDigit)
		if !hasRepeatedDigits(nextNumber) {
			generate(n-1, nextNumber, outcome)
		}
	}
}

func hasRepeatedDigits(number string) bool {
	seen := make(map[rune]bool)
	for _, digit := range number {
		if seen[digit] {
			return true
		}
		seen[digit] = true
	}
	return false
}
