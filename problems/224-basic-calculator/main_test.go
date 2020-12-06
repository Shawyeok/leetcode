package leetcode

import "testing"

type testcase struct {
	expression string
	expected   int
}

func Test_calculate(t *testing.T) {
	suit := []testcase{
		{
			"0",
			0,
		},
		{
			"  30",
			30,
		},
		{
			"1",
			1,
		},
		{
			"12345",
			12345,
		},
		{
			"1 + 1",
			2,
		},
		{
			" 2-1 + 2 ",
			3,
		},
		{
			"1000 + 1",
			1001,
		},
		{
			"10 - (20 - 30)",
			20,
		},
		{
			"2-(5-6)",
			3,
		},
		{
			"(1+(4+5+2)-3)+(6+8)",
			23,
		},
		{
			"1+4+5+2-3+6+8",
			23,
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
