package leetcode

func maxProfit(prices []int) int {
	l := len(prices)
	if l <= 1 {
		return 0
	}
	var p, m, t = 0, 0, 0
	for i := 1; i < l; i++ {
		t = p + prices[i] - prices[i-1]
		if t > 0 {
			p, m = t, max(m, t)
		} else {
			p = 0
		}
	}
	return m
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
