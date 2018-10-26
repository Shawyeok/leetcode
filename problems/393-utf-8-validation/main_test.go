package main

import "testing"

type testcase struct {
	data     []int
	expected bool
}

func Test_validUtf8(t *testing.T) {
	suit := []testcase{
		testcase{
			[]int{197, 130, 1},
			true,
		},
		testcase{
			[]int{235, 140, 4},
			false,
		},
		testcase{
			[]int{1},
			true,
		},
		testcase{
			[]int{130},
			false,
		},
		testcase{
			[]int{197},
			false,
		},
		testcase{
			[]int{240},
			false,
		},
		testcase{
			[]int{130, 4},
			false,
		},
		testcase{
			[]int{250, 145, 145, 145, 145},
			false,
		},
		testcase{
			[]int{250, 145, 145, 145, 4},
			false,
		},
		testcase{
			[]int{250, 140, 140, 140},
			false,
		},
	}
	for _, tc := range suit {
		actual := validUtf8(tc.data)
		if actual != tc.expected {
			t.Errorf("data: %v, expected: %v, actual: %v", tc.data, tc.expected, actual)
		}
	}
}
