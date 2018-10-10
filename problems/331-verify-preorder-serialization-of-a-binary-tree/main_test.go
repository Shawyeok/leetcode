package main

import "testing"

type testcase struct {
	preorder string
	expected bool
}

func Test_isValidSerialization(t *testing.T) {
	suit := []testcase{
		testcase{
			"9,3,4,#,#,1,#,#,2,#,6,#,#",
			true,
		},
		testcase{
			"#",
			true,
		},
		testcase{
			"#,#",
			false,
		},
		testcase{
			"9,3,4,#,#,#,#",
			true,
		},
		testcase{
			"1,#",
			false,
		},
		testcase{
			"9,#,#,1",
			false,
		},
		testcase{
			"",
			false,
		},
		testcase{
			"1",
			false,
		},
	}

	for _, tc := range suit {
		actual := isValidSerialization(tc.preorder)
		if actual != tc.expected {
			t.Errorf("preorder: %v, expected: %v, actual: %v", tc.preorder, tc.expected, actual)
		}
	}
}
