package main

func main() {
}

func numJewelsInStones(J string, S string) int {
	jewels := make(map[rune]bool)
	for _, b := range J {
		jewels[b] = true
	}
	c := 0
	for _, b := range S {
		if jewels[b] {
			c++
		}
	}

	return c
}
