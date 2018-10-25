package main

import (
	"strings"
)

func main() {
}

func decodeString(s string) string {
	i, st := 0, 0
	for i < len(s) && !isNumber(s[i]) {
		i++
	}
	str := s[st:i]
	if len(s) == i {
		return str
	}
	k := 0
	for isNumber(s[i]) {
		k = k*10 + int(s[i]-'0')
		i++
	}
	c, st := 0, i+1
	for ok := true; ok; ok = c > 0 {
		if s[i] == '[' {
			c++
		}
		if s[i] == ']' {
			c--
		}
		i++
	}
	rep := decodeString(s[st : i-1])
	str += strings.Repeat(rep, k)
	if len(s) > i {
		str += decodeString(s[i:])
	}
	return str
}

func isNumber(b byte) bool {
	return b >= '0' && b <= '9'
}
