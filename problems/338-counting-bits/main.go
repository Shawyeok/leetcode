package leetcode

func countBits(num int) []int {
	r := make([]int, num+1, num+1)
	r[0] = 0
	m := 0
	for i := 1; i <= num; i <<= 1 {
		r[i] = 1
		m = i
	}
	for j := 1; j < m; j++ {
		k := smallestK(j)
		for i := k + j; i <= num; i = k + j {
			r[i] = 1 + r[j]
			k <<= 1
		}
	}
	return r
}

// compute the smallest power of two that grater than integer v
func smallestK(v int) int {
	v |= v >> 1
	v |= v >> 2
	v |= v >> 4
	v |= v >> 8
	v |= v >> 16
	v++
	return v
}
