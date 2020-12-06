package leetcode

func calculate(s string) int {
	l := len(s)
	var cs []int
	var ops []int
	cursor := 0
	for i := 0; i < l; i++ {
		switch s[i] {
		case ' ':
			// ignore
		case '(':
			cs = append(cs, cursor)
		case ')':
			lr := len(cs)
			cursor = cs[lr-1]
			n := calc(ops[cursor:])
			ops = ops[:cursor]
			ops = append(ops, n)
			cursor++
			cs = cs[:lr-1]
		case '+', '-':
			ops = append(ops, int(s[i]))
			cursor++
		default:
			n := int(s[i] - '0')
			j := i + 1
			for ; j < l && s[j] >= '0'; j++ {
				n = n*10 + int(s[j]-'0')
			}
			i = j - 1
			ops = append(ops, n)
			cursor++
		}
	}
	return calc(ops)
}

func calc(ops []int) int {
	l := len(ops)
	if l == 1 {
		return ops[0]
	}
	n := ops[0]
	for i, j := 1, 2; j < l; {
		n = doCalc(n, ops[j], ops[i])
		i, j = i+2, j+2
	}
	return n
}

func doCalc(n1, n2 int, op int) int {
	var m int
	switch op {
	case '+':
		m = n1 + n2
	case '-':
		m = n1 - n2
	}
	return m
}
