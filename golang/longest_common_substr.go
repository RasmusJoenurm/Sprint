package sprint

func LongestCommonSubstr(str1, str2 string) string {
	m, n := len(str1), len(str2)
	longest := make([][]int, m+1)
	for i := range longest {
		longest[i] = make([]int, n+1)
	}

	length := 0
	endIndex := 0

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if str1[i-1] == str2[j-1] {
				longest[i][j] = longest[i-1][j-1] + 1
				if longest[i][j] > length {
					length = longest[i][j]
					endIndex = i - 1
				}
			} else {
				longest[i][j] = 0
			}
		}
	}
	if length == 0 {
		return ""
	}
	return str1[endIndex-length+1 : endIndex+1]
}
