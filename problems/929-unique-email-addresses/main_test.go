package main

import "testing"

type testcase struct {
	emails   []string
	expected int
}

func Test_numUniqueEmails(t *testing.T) {
	suit := []testcase{
		testcase{
			[]string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"},
			2,
		},
		testcase{
			[]string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com", "kmter@live.com", "kmter.com@gmail.com", "kmter.com+1@gmail.com"},
			4,
		},
	}
	for _, tc := range suit {
		actual := numUniqueEmails(tc.emails)
		if actual != tc.expected {
			t.Errorf("emails: %v, expected: %v, actual: %v", tc.emails, tc.expected, actual)
		}
	}
}
