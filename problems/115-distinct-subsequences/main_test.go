package main

import "testing"

type testcase struct {
	s        string
	t        string
	expected int
}

func Test_numDistinct(t *testing.T) {
	suit := []testcase{
		testcase{
			"rabbbit",
			"rabbit",
			3,
		},
		testcase{
			"babgbag",
			"bag",
			5,
		},
		testcase{
			"baaaaaag",
			"bag",
			6,
		},
		testcase{
			"bag",
			"bag",
			1,
		},
		testcase{
			"",
			"",
			1,
		},
		testcase{
			"adbdadeecadeadeccaeaabdabdbcdabddddabcaaadbabaaedeeddeaeebcdeabcaaaeeaeeabcddcebddebeebedaecccbdcbcedbdaeaedcdebeecdaaedaacadbdccabddaddacdddc",
			"bcddceeeebecbc",
			700531452,
		},
	}

	for _, tc := range suit {
		actual := numDistinct(tc.s, tc.t)
		if actual != tc.expected {
			t.Errorf("s: %v, t: %v, expected: %v, actual: %v", tc.s, tc.t, tc.expected, actual)
		}
	}
}
