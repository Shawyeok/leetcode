package main

import "testing"

type testcase struct {
	str      string
	expected int
}

func Test_myAtoi(t *testing.T) {
	suit := []testcase{
		testcase{
			"42",
			42,
		},
		testcase{
			"     -42    ",
			-42,
		},
		testcase{
			"4193 with words",
			4193,
		},
		testcase{
			"-91283472332",
			-2147483648,
		},
		testcase{
			"4193 with words 987",
			4193,
		},
		testcase{
			"-",
			0,
		},
		testcase{
			"-1",
			-1,
		},
		testcase{
			"",
			0,
		},
		testcase{
			"20000000000000000000",
			2147483647,
		},
	}

	for _, tc := range suit {
		actual := myAtoi(tc.str)
		if actual != tc.expected {
			t.Errorf("str: %v, expected: %v, actual: %v", tc.str, tc.expected, actual)
		}
	}
}
