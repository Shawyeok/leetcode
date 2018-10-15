package main

import "testing"

type testcase struct {
	nums     []int
	expected bool
}

func Test_find132pattern(t *testing.T) {
	suit := []testcase{
		testcase{
			[]int{1, 2, 3, 4},
			false,
		},
		testcase{
			[]int{3, 1, 4, 2},
			true,
		},
		testcase{
			[]int{-1, 3, 2, 0},
			true,
		},
		testcase{
			[]int{-1, -3, -4, -5},
			false,
		},
		testcase{
			[]int{-1, 0, 3, -5, 0},
			true,
		},
	}

	for _, tc := range suit {
		actual := find132pattern(tc.nums)
		if actual != tc.expected {
			t.Errorf("nums: %v, expected: %v, actual: %v", tc.nums, tc.expected, actual)
		}
	}
}
