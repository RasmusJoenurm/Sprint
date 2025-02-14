package sprint

import "fmt"

func Combinations() string {

	result := ""
	for i := 0; i < 8; i++ {
		for j := i + 1; j < 9; j++ {
			for k := j + 1; k < 10; k++ {

				if result != "" {
					result += ", "
				}
				result += fmt.Sprintf("%d%d%d", i, j, k)
			}
		}
	}
	return result
}
