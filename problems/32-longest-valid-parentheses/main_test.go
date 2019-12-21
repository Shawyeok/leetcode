package main

import "testing"

type testcase struct {
	s        string
	expected int
}

func Test_longestValidParentheses(t *testing.T) {
	suit := []testcase{
		testcase{
			"())(()",
			2,
		},
		testcase{
			"()(()",
			2,
		},
		testcase{
			"()))((()",
			2,
		},
		testcase{
			")))(((",
			0,
		},
		testcase{
			"((()))",
			6,
		},
		testcase{
			"(()(()",
			2,
		},
		testcase{
			"()()()(",
			6,
		},
	}
	for _, tc := range suit {
		actual := longestValidParentheses(tc.s)
		if actual != tc.expected {
			t.Errorf("s: %v, expected: %v, actual: %v", tc.s, tc.expected, actual)
		}
	}
}
