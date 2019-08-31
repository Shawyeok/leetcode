package main

func main() {
}

func minCostClimbingStairs(cost []int) int {
	f1, f2 := 0, 0
	for i := len(cost) - 1; i >= 0; i-- {
		f0 := cost[i] + min(f1, f2)
		f2 = f1
		f1 = f0
	}
	return min(f1, f2)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
