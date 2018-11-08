package main

import "testing"

type testcase struct {
	J        string
	S        string
	expected int
}

func Test_numJewelsInStones(t *testing.T) {
	suit := []testcase{
		testcase{
			"aA",
			"aAAbbbb",
			3,
		},
	}
	for _, tc := range suit {
		actual := numJewelsInStones(tc.J, tc.S)
		if actual != tc.expected {
			t.Errorf("J: %v, S: %v, expected: %v, actual: %v", tc.J, tc.S, tc.expected, actual)
		}
	}
}
