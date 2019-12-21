package main

func main() {
}

func longestValidParentheses(s string) int {
	k, p, maxP := 0, 0, 0
	for _, c := range s {
		if c == '(' {
			k++
		} else {
			k--
			if k < 0 {
				k = 0
				maxP = max(maxP, p)
				p = 0
			} else {
				p++
			}
		}
	}
	maxP = max(maxP, p)
	return maxP * 2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
