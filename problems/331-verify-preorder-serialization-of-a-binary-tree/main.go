package main

import (
	"strings"
)

func main() {
}

func isValidSerialization(preorder string) bool {
	nodes := strings.Split(preorder, ",")
	if preorder == "" {
		return false
	}

	heap := make([]string, len(nodes))
	p := 0
	for i := range nodes {
		heap[p] = nodes[i]
		for last := p - 1; heap[p] == "#" && last > -1 && heap[last] == "#"; last = p - 1 {
			p -= 2
			if p < 0 {
				return false
			}

			heap[p] = "#"
		}
		p++
	}

	return p == 1 && heap[p-1] == "#"
}
