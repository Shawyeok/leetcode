package leetcode

func calculate(s string) int {
	ops := [2]int32{}
	nums := [3]int32{}
	ni := 0
	oi := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case ' ':
			// ignore
		case '+':
			if oi > 0 {
				nums[0] = doCalculate(nums[0], nums[1], ops[0])
				oi, ni = 0, 1
			}
			ops[oi] = '+'
			oi++
		case '-':
			if oi > 0 {
				nums[0] = doCalculate(nums[0], nums[1], ops[0])
				oi, ni = 0, 1
			}
			ops[oi] = '-'
			oi++
		case '*':
			if oi > 0 && (ops[0] == '*' || ops[0] == '/') {
				nums[0] = doCalculate(nums[0], nums[1], ops[0])
				oi, ni = 0, 1
			}
			ops[oi] = '*'
			oi++
		case '/':
			if oi > 0 && (ops[0] == '*' || ops[0] == '/') {
				nums[0] = doCalculate(nums[0], nums[1], ops[0])
				oi, ni = 0, 1
			}
			ops[oi] = '/'
			oi++
		default:
			var n int32
			for ; i < len(s) && s[i] >= '0'; i++ {
				n = n*10 + int32(s[i]-'0')
			}
			i--
			nums[ni] = n
			ni++
			if ni == 3 {
				nums[1] = doCalculate(nums[1], nums[2], ops[1])
				oi, ni = 1, 2
			}
		}
	}
	if oi == 0 {
		return int(nums[0])
	}
	return int(doCalculate(nums[0], nums[1], ops[0]))
}

func doCalculate(s1, s2 int32, op rune) int32 {
	var n int32
	switch op {
	case '+':
		n = s1 + s2
	case '-':
		n = s1 - s2
	case '*':
		n = s1 * s2
	case '/':
		n = s1 / s2
	}
	return n
}
