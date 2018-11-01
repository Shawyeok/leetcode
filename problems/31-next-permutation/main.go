package main

import (
	"math"
)

func main() {
}

func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}
	for i := len(nums) - 2; i >= 0; i-- {
		k := 0
		min := math.MinInt32
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > nums[i] {
				if min == math.MinInt32 {
					min = nums[j]
					k = j
				} else if nums[j] < min {
					min = nums[j]
					k = j
				}
			}
		}
		if min != math.MinInt32 {
			swap(nums, i, k)
			quicksort(nums, i+1, len(nums)-1)
			return
		}
	}
	for i := 0; i < len(nums); i++ {
		j := len(nums) - 1 - i
		if i < j {
			swap(nums, i, j)
		}
	}
}

func quicksort(arr []int, low, high int) {
	if low < high {
		p := partition(arr, low, high)
		quicksort(arr, low, p-1)
		quicksort(arr, p+1, high)
	}
}

func partition(arr []int, low, high int) int {
	comp := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < comp {
			if i != j {
				swap(arr, i, j)
			}
			i++
		}
	}
	swap(arr, i, high)
	return i
}

func swap(nums []int, i, j int) {
	temp := nums[j]
	nums[j] = nums[i]
	nums[i] = temp
}
