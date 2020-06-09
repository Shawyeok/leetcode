package main

import "testing"

type testcase struct {
	matrix   [][]int
	k        int
	expected int
}

func Test_maxSumSubmatrix(t *testing.T) {
	testcases := []testcase{
		{
			[][]int{
				[]int{1, 0, 1},
				[]int{0, -2, 3},
			},
			2,
			2,
		},
	}
	for _, tc := range testcases {
		actual := maxSumSubmatrix(tc.matrix, tc.k)
		if actual != tc.expected {
			t.Errorf("matrix: %v, k: %v, expected: %v, actual: %v", tc.matrix, tc.k, tc.expected, actual)
		}
	}
}
