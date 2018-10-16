package main

import (
	"math"
)

func main() {
}

func superPow(a int, b []int) int {
	mod := 0
	for _, exp := range b {
		math.Pow(a, exp) % 1337
	}
}
