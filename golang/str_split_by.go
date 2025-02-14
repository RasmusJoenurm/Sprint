package sprint

func StrSplitBy(s, sep string) []string {

	var outcome []string
	temp := ""
	sepLen := len(sep)

	if s == "" {
		return []string{}
	}

	for i := 0; i < len(s); i++ {
		if i+sepLen <= len(s) && s[i:i+sepLen] == sep {
			outcome = append(outcome, temp)
			temp = ""
			i += sepLen - 1
		} else {
			temp += string(s[i])
		}
	}
	if temp != "" {
		outcome = append(outcome, temp)
	}
	return outcome
}
