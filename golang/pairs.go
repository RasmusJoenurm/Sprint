package sprint

import "fmt"

func Pairs() string {
	result := ""
	for i := 0; i < 100; i++ {
		for j := i + 1; j < 100; j++ {
			if result != "" {
				result += ", "
			}
			result += fmt.Sprintf("%02d %02d", i, j)
		}
	}

	return result
}
