package main

func main() {
}

func validUtf8(data []int) bool {
	k := 0
	for _, n := range data {
		if k == 0 {
			for t, i := 0, uint(1); i < 8; i++ {
				t = 1 << (8 - i)
				if t&n != t {
					break
				}
				k++
			}
			if k == 1 || k > 4 {
				return false
			}
			if k > 0 {
				k--
			}
			continue
		}

		if !startsWith10(n) {
			return false
		}
		k--
	}

	return k == 0
}

func startsWith10(n int) bool {
	x := 1 << (8 - 1)
	y := 1 << (8 - 2)
	return x&n == x && y&n == 0
}
