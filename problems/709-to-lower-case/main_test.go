package main

import (
	"testing"
)

func Test_toLowerCase(t *testing.T) {
	params := map[string]string{
		"PPALLP": "ppallp",
		"PPAlll": "ppalll",
		"abc":    "abc",
		"你好A":    "你好a",
	}
	for s, expected := range params {
		actual := toLowerCase(s)
		if actual != expected {
			t.Errorf("s: %v, expected: %v, actual: %v", s, expected, actual)
		}
	}
}
