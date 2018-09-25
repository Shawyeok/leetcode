package main

import (
	"strconv"
)

func main() {
}

func singleNumber(nums []int) int {
	var res int
	for i := uint(0); i < strconv.IntSize; i++ {
		x := 1 << i
		var sum int
		for j := range nums {
			if x&nums[j] != 0 {
				sum++
			}
		}
		res |= (sum % 3) << i
	}

	return res
}
