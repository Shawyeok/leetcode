package main

import "testing"

type testcase struct {
	s        string
	expected string
}

func Test_decodeString(t *testing.T) {
	suit := []testcase{
		testcase{
			"abc",
			"abc",
		},
		testcase{
			"3[a]2[bc]",
			"aaabcbc",
		},
		testcase{
			"3[a2[c]]",
			"accaccacc",
		},
		testcase{
			"2[abc]3[cd]ef",
			"abcabccdcdcdef",
		},
	}

	for _, tc := range suit {
		actual := decodeString(tc.s)
		if actual != tc.expected {
			t.Errorf("s: %v, expected: %v, actual: %v", tc.s, tc.expected, actual)
		}
	}
}
