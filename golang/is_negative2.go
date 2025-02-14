package sprint

func IntVsFloat2(i int, f float32) string {
	if float32(i) < f {
		return "Float"
	} else if float32(i) > f {
		return "Integar"
	} else {
		return "Same"
	}
}
