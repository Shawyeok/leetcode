package main

func main() {
}

func numberOfArithmeticSlices(A []int) int {
	if len(A) < 3 {
		return 0
	}
	diffs := make([]int, len(A)-1)
	for i := 0; i < len(A)-1; i++ {
		diffs[i] = A[i] - A[i+1]
	}

	res := 0
	nums := 0
	for i := 0; i < len(diffs)-1; i++ {
		if diffs[i] == diffs[i+1] {
			nums++
			res += nums
		} else {
			nums = 0
		}
	}
	return res
}
