package main

import "testing"

type testcase struct {
	dividend int
	divisor  int
	expected int
}

func Test_divide(t *testing.T) {
	suit := []testcase{
		testcase{
			10,
			3,
			3,
		},
		testcase{
			7,
			-3,
			-2,
		},
		testcase{
			7,
			7,
			1,
		},
		testcase{
			1,
			1,
			1,
		},
		testcase{
			1,
			-1,
			-1,
		},
		testcase{
			0,
			-3,
			0,
		},
		testcase{
			2129871,
			2,
			1064935,
		},
		testcase{
			-2147483648,
			1,
			-2147483648,
		},
		testcase{
			-2147483648,
			-1,
			2147483647,
		},
	}

	for _, tc := range suit {
		actual := divide(tc.dividend, tc.divisor)
		if actual != tc.expected {
			t.Errorf("dividend: %v, divisor: %v, expected: %v, actual: %v", tc.dividend, tc.divisor, tc.expected, actual)
		}
	}
}
