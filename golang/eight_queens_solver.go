package sprint

import (
	"strconv"
	"strings"
)

func EightQueensSolver() string {
	var results []string
	solve([]int{}, &results)
	return strings.Join(results, "\n")
}

func solve(queens []int, results *[]string) {
	n := len(queens)
	if n == 8 {
		*results = append(*results, formatSolution(queens))
		return
	}
	for i := 1; i <= 8; i++ {
		if isValid(queens, i) {
			solve(append(queens, i), results)
		}
	}
}

func isValid(queens []int, pos int) bool {
	n := len(queens)
	for i := 0; i < n; i++ {
		if queens[i] == pos || queens[i]-i == pos-n || queens[i]+i == pos+n {
			return false
		}
	}
	return true
}

func formatSolution(queens []int) string {
	var sb strings.Builder
	for _, q := range queens {
		sb.WriteString(strconv.Itoa(q))
	}
	return sb.String()
}
