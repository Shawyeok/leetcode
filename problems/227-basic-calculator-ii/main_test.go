package leetcode

import "testing"

type testcase struct {
	expression string
	expected   int
}

func Test_calculate(t *testing.T) {
	suit := []testcase{
		{
			"1",
			1,
		},
		{
			"12345",
			12345,
		},
		{
			"1 + 2 * 3 - 4",
			3,
		},
		{
			"3+2*2",
			7,
		},
		{
			" 3/2 ",
			1,
		},
		{
			" 3+5 / 2 ",
			5,
		},
	}
	for _, tc := range suit {
		actual := calculate(tc.expression)
		if actual != tc.expected {
			t.Fail()
			t.Logf("expression: %v, expected: %v, actual: %v", tc.expression, tc.expected, actual)
		}
	}
}
