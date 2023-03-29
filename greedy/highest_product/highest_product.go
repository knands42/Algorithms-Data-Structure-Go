package highest_product

func highestProducts(products []int) int {
	highestValue := 0

	for idx, v := range products {
		currentValue := v
		nextValue := products[idx+1]
		nextOfNextValue := products[idx+2]

		multiplication := currentValue * nextValue * nextOfNextValue
		highestValue = max(highestValue, multiplication)

		if idx+2 == len(products)-1 {
			break
		}
	}

	return highestValue
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
