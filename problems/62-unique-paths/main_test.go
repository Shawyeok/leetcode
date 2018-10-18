package main

import "testing"

type testcase struct {
	m        int
	n        int
	expected int
}

func Test_uniquePaths(t *testing.T) {
	suit := []testcase{
		testcase{
			3,
			2,
			3,
		},
		testcase{
			7,
			3,
			28,
		},
		testcase{
			1,
			1,
			1,
		},
		testcase{
			51,
			9,
			1916797311,
		},
	}

	for _, tc := range suit {
		actual := uniquePaths(tc.m, tc.n)
		if actual != tc.expected {
			t.Errorf("m: %v, n: %v, expected: %v, actual: %v", tc.m, tc.n, tc.expected, actual)
		}
	}
}
