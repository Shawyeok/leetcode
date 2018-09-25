package main

import "testing"

type testcase struct {
	nums     []int
	expected int
}

func Test_singleNumber(t *testing.T) {
	suit := []testcase{
		testcase{
			[]int{2, 2, 1, 1, 3, 1, 3, 3, 4, 2},
			4,
		},
		testcase{
			[]int{-2, -2, 1, 1, -3, 1, -3, -3, -4, -2},
			-4,
		},
		testcase{
			[]int{2, 2, 3, 2},
			3,
		},
	}

	for _, tc := range suit {
		actual := singleNumber(tc.nums)
		if actual != tc.expected {
			t.Errorf("nums: %v, expected: %v, actual: %v", tc.nums, tc.expected, actual)
		}
	}
}
