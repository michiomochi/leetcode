func maxProfit(prices []int) int {
	buy := 0
	sell := 0
	score := 0

	for i, p := range prices {
		if i == 0 {
			buy = p
		}

		if p > buy && p > sell && (p - buy) > score {
			sell = p
			score = sell - buy
		}

		if p < buy {
			buy = p
			sell = 0
		}
	}

	return score
}