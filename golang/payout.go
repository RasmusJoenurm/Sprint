package sprint

func Payout(amount int, denominations []int) (payout []int) {
	for i := 0; i < len(denominations); i++ {
		for j := i + 1; j < len(denominations); j++ {
			if denominations[i] < denominations[j] {
				denominations[i], denominations[j] = denominations[j], denominations[i]
			}
		}
	}
	for _, denom := range denominations {
		for amount >= denom {
			payout = append(payout, denom)
			amount -= denom
		}
	}
	if amount != 0 {
		return []int{}
	}
	return payout
}
