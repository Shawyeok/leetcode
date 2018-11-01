package main

import (
	"testing"
)

type testcase struct {
	nums     []int
	expected []int
}

func Test_nextPermutation(t *testing.T) {
	suit := []testcase{
		testcase{
			[]int{1, 2, 3},
			[]int{1, 3, 2},
		},
		testcase{
			[]int{3, 2, 1},
			[]int{1, 2, 3},
		},
		testcase{
			[]int{1, 1, 5},
			[]int{1, 5, 1},
		},
		testcase{
			[]int{1, 2, 4, 3},
			[]int{1, 3, 2, 4},
		},
	}

	for _, tc := range suit {
		original := make([]int, len(tc.nums))
		copy(original, tc.nums)
		nextPermutation(tc.nums)
		if !isEqual(tc.nums, tc.expected) {
			t.Errorf("nums: %v, expected: %v, actual: %v", original, tc.expected, tc.nums)
		}
	}
}

func isEqual(nums1, nums2 []int) bool {
	if len(nums1) != len(nums2) {
		return false
	}
	for i := range nums1 {
		if nums1[i] != nums2[i] {
			return false
		}
	}
	return true
}
