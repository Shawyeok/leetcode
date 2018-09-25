package main

import (
	"testing"
)

func Test_checkRecord(t *testing.T) {
	params := map[string]bool{
		"PPALLP":  true,
		"PPALLL":  false,
		"LALL":    true,
		"AA":      false,
		"LLALLPP": true,
		"LLALLA":  false,
		"LLALLAL": false,
	}
	for s, expected := range params {
		actual := checkRecord(s)
		if actual != expected {
			t.Errorf("s: %v, expected: %v, actual: %v", s, expected, actual)
		}
	}
}
