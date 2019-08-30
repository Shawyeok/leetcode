package main

func main() {
}

func toLowerCase(str string) string {
	chars := []rune(str)
	for i, c := range chars {
		if c >= 65 && c <= 90 {
			chars[i] = c + 32
		}
	}
	return string(chars)
}
