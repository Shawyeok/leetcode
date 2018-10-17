package main

import "testing"

type testcase struct {
	A        []int
	expected int
}

func Test_numberOfArithmeticSlices(t *testing.T) {
	suit := []testcase{
		testcase{
			[]int{1, 2, 3, 4},
			3,
		},
		testcase{
			[]int{1, 2, 2, 3, 4, 6},
			1,
		},
		testcase{
			[]int{1, 1, 2, 5, 7},
			0,
		},
		testcase{
			[]int{1, 3, 5, 7, 9},
			6,
		},
		testcase{
			[]int{1, 1, 2, 2},
			0,
		},
	}

	for _, tc := range suit {
		actual := numberOfArithmeticSlices(tc.A)
		if actual != tc.expected {
			t.Errorf("A: %v, expected: %v, actual: %v", tc.A, tc.expected, actual)
		}
	}
}
