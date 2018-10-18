package main

func main() {
}

func uniquePaths(m int, n int) int {
	matrix := makeMatrix(m+1, n+1)
	matrix[0][1] = 1
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			matrix[i][j] = matrix[i][j-1] + matrix[i-1][j]
		}
	}

	return matrix[m][n]
}

func makeMatrix(n, m int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, m)
	}

	return matrix
}
