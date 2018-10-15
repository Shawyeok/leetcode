package main

func main() {
}

func find132pattern(nums []int) bool {
	for i := len(nums) - 2; i > 0; i-- {
		s := nums[i]
		for j := i - 1; j >= 0; j-- {
			s = min(s, nums[j])
		}
		if s == nums[i] {
			continue
		}
		for k := i + 1; k < len(nums); k++ {
			if nums[k] < nums[i] && nums[k] > s {
				return true
			}
		}
	}

	return false
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
