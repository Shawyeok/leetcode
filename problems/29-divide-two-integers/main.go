package main

import (
	"math"
)

func main() {
}

func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	positive := true
	if (dividend < 0 && divisor > 0) || (dividend > 0 && divisor < 0) {
		positive = false
	}
	res := divideUnsigned(abs(dividend), abs(divisor))

	if positive {
		return res
	}
	return -res
}

func divideUnsigned(dividend int, divisor int) int {
	if dividend < divisor {
		return 0
	} else if dividend == divisor {
		return 1
	}

	var e uint
	n := divisor
	for n < dividend {
		n <<= 1
		e++
	}
	diff := dividend - (n >> 1)
	if diff >= divisor {
		return 1<<(e-1) + divideUnsigned(diff, divisor)
	}

	return 1 << (e - 1)
}

func abs(num int) int {
	if num > 0 {
		return num
	}
	return -num
}
