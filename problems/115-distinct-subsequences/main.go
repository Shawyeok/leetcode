package main

func main() {
}

func numDistinct(s string, t string) int {
	matrix := makeMatrix(len(t)+1, len(s)+1)
	for i := 0; i <= len(s); i++ {
		matrix[0][i] = 1
	}

	for i := range t {
		for j := range s {
			if t[i] == s[j] {
				matrix[i+1][j+1] = matrix[i][j] + matrix[i+1][j]
			} else {
				matrix[i+1][j+1] = matrix[i+1][j]
			}
		}
	}

	return matrix[len(t)][len(s)]
}

func makeMatrix(n, m int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, m)
	}

	return matrix
}
