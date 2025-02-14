package sprint

import (
	"sort"
	"strings"
)

func AreAnagrams(str1, str2 string) bool {

	cleanAndSort := func(s string) string {
		s = strings.ReplaceAll(s, " ", "")
		s = strings.ToLower(s)

		runes := []rune(s)

		sort.Slice(runes, func(i, j int) bool { return runes[i] < runes[j] })

		return string(runes)
	}

	cleanedStr1 := cleanAndSort(str1)
	cleanedStr2 := cleanAndSort(str2)

	return cleanedStr1 == cleanedStr2
}
