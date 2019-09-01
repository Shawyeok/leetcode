package main

import "testing"

func Test_findSubstringInWraproundString(t *testing.T) {
	params := map[string]int{
		"a":             1,
		"cac":           2,
		"zab":           6,
		"zabc":          10,
		"zabcezabcxdef": 17,
	}
	for p, expected := range params {
		actual := findSubstringInWraproundString(p)
		if actual != expected {
			t.Errorf("p: %v, expected: %v, actual: %v", p, expected, actual)
		}
	}
}
