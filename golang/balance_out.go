package sprint

func BalanceOut(arr []bool) []bool {
	countTrue := 0
	countFalse := 0

	for _, value := range arr {
		if value {
			countTrue++
		} else {
			countFalse++
		}
	}
	difference := countTrue - countFalse
	if difference < 0 {
		for i := 0; i < -difference; i++ {
			arr = append(arr, true)
		}
	} else if difference > 0 {
		for i := 0; i < difference; i++ {
			arr = append(arr, false)
		}
	}
	return arr
}
