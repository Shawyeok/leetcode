package leetcode

import "testing"

type testcase struct {
	prices   []int
	expected int
}

func Test_maxProfit(t *testing.T) {
	suit := []testcase{
		{
			[]int{1, 2},
			1,
		},
		{
			[]int{7, 1, 5, 3, 6, 4},
			5,
		},
		{
			[]int{7, 1, 5, 1, 2, 1},
			4,
		},
		{
			[]int{7, 6, 4, 3, 1},
			0,
		},
	}
	for _, tc := range suit {
		actual := maxProfit(tc.prices)
		if actual != tc.expected {
			t.Fail()
			t.Logf("prices: %v, expected: %v, actual: %v", tc.prices, tc.expected, actual)
		}
	}
}
