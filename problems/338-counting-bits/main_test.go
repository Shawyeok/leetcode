package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	num      int
	expected []int
}

func Test_countBits(t *testing.T) {
	suit := []testcase{
		{
			0,
			[]int{0},
		},
		{
			2,
			[]int{0, 1, 1},
		},
		{

			5,
			[]int{0, 1, 1, 2, 1, 2},
		},
		{
			15,
			[]int{0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4},
		},
		{
			100,
			[]int{0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4, 1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5, 1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5, 2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6, 1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5, 2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6, 2, 3, 3, 4, 3},
		},
	}
	for _, tc := range suit {
		actual := countBits(tc.num)
		if !assert.ObjectsAreEqual(tc.expected, actual) {
			t.Logf("For countBits(%d), expected: %v, actual: %v\n", tc.num, tc.expected, actual)
			t.Fail()
		}
	}
}

func Test_smallestK(t *testing.T) {
	k := smallestK(1024)
	println(k)
}
