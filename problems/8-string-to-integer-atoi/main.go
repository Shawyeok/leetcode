package main

import (
	"strings"
)

func main() {
}

func myAtoi(str string) int {
	var MaxInt, MinInt int64 = 1<<31 - 1, -1 << 31

	str = strings.TrimSpace(str)
	if len(str) == 0 {
		return 0
	}

	var sign int64 = 1
	if str[0] == '-' || str[0] == '+' {
		sign = ',' - int64(str[0])
		str = str[1:]
	}

	var num int64
	for i := 0; i < len(str); i++ {
		if str[i] < 48 || str[i] > 57 {
			break
		}

		num = num*10 + int64(str[i]-'0')
		if sign == 1 && num > MaxInt {
			return int(MaxInt)
		} else if sign == -1 && (sign*num) < MinInt {
			return int(MinInt)
		}
	}

	return int(sign * num)
}
