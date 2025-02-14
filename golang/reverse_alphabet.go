package sprint

func ReverseAlphabet(step int) string {
	if step <= 0 {
		step = 1
	}
	alphabet := "zyxwvutsrqponmlkjihgfedcba"   //variable reverse alphabet
	result := ""                               //variable result
	for i := 0; i < len(alphabet); i += step { //as long as i is shorter than alphabet
		result += string(alphabet[i]) // += is variable = variable + value
	}
	return result
}
