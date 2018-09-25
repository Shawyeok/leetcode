package main

func main() {
}

func checkRecord(s string) bool {
	var absentCount int
	var cLate bool

	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == 'L' {
			var j int
			for j = i + 1; j < len(s) && s[j] == 'L'; j++ {
			}
			if (j - i) > 2 {
				cLate = true
				break
			}
			i = j - 1
		} else if c == 'A' {
			absentCount++
			if absentCount > 1 {
				break
			}
		}
	}

	return absentCount <= 1 && !cLate
}
