package main

import "testing"

type testcase struct {
	a        int
	b        []int
	expected int
}

func Test_superPow(t *testing.T) {
	suit := []testcase{
		testcase{
			2,
			[]int{3},
			8,
		},
		testcase{
			2,
			[]int{1, 0},
			1024,
		},
	}

	for _, tc := range suit {
		actual := superPow(tc.a, tc.b)
		if actual != tc.expected {
			t.Errorf("a: %v, b: %v, expected: %v, actual: %v", tc.a, tc.b, tc.expected, actual)
		}
	}
}
