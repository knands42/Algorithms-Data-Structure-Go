package bulbs

func lightBulbs(bulbs []int) int {
	cost := 0
	for _, v := range bulbs {
		if cost%2 != 0 {
			v = 1 - v
		}
		if v%2 == 0 {
			cost++
		}
	}

	return cost
}
