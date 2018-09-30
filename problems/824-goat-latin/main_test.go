package main

import (
	"testing"
)

type testcase struct {
	S        string
	expected string
}

func Test_toGoatLatin(t *testing.T) {
	suit := []testcase{
		testcase{
			"I speak Goat Latin",
			"Imaa peaksmaaa oatGmaaaa atinLmaaaaa",
		},
		testcase{
			"The quick brown fox jumped over the lazy dog",
			"heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa",
		},
	}

	for _, tc := range suit {
		actual := toGoatLatin(tc.S)
		if actual != tc.expected {
			t.Errorf("S: %v, expected: %v, actual: %v", tc.S, tc.expected, actual)
		}
	}
}
