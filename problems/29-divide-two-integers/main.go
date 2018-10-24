package main

func main() {
}

func divide(dividend int, divisor int) int {
	MaxInt, MinInt := 1<<31-1, -1<<31

	sign := 1
	if dividend < 1 {
		sign ^= -1
	} else {
		sign ^= 1
	}
	if divisor < 1 {
		sign ^= -1
	} else {
		sign ^= 1
	}
	res := int(divideSigned(abs(dividend), abs(divisor)))
	if sign > 0 {
		if res > MaxInt {
			return MaxInt
		}
		return res
	}
	if res < MinInt {
		return MinInt
	}
	return -res
}

func divideSigned(dividend int, divisor int) uint {
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
		return 1<<(e-1) + divideSigned(diff, divisor)
	}

	return 1 << (e - 1)
}

func abs(num int) int {
	if num > 0 {
		return num
	}
	return -num
}
