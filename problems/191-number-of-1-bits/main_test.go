package leetcode

import (
	"testing"
)

type testcase struct {
	num      uint32
	expected int
}

func Test_hammingWeight(t *testing.T) {
	suit := []testcase{
		{
			1,
			1,
		},
		{
			1024,
			1,
		},
		{
			7,
			3,
		},
	}
	for _, tc := range suit {
		actual := hammingWeight(tc.num)
		if tc.expected != actual {
			t.Logf("hammingWeight(%d), expected: %v, actual: %v", tc.num, tc.expected, actual)
			t.Fail()
		}
	}
}
