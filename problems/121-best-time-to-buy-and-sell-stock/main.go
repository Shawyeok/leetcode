package leetcode

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	best, lowest := 0, prices[0]
	for i := 1; i < len(prices); i++ {
		price := prices[i]
		if price-lowest > best {
			best = price - lowest
		}
		if price < lowest {
			lowest = price
		}
	}
	return best
}
