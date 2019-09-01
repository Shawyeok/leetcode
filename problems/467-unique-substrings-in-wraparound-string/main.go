package main

func main() {
}

func findSubstringInWraproundString(p string) int {
	cs := make([]int, 26)
	index := 0
	for index < len(p) {
		i := index + 1
		for i < len(p) && (p[i]-'a') == (p[i-1]-'a'+1)%26 {
			i++
		}
		for j := index; j < i; j++ {
			pos := p[j] - 'a'
			cs[pos] = max(cs[pos], i-j)
		}
		index = i
	}
	return sum(cs)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func sum(arr []int) int {
	sum := 0
	for _, value := range arr {
		sum += value
	}
	return sum
}
