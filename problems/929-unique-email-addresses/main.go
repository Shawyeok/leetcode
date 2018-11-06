package main

func main() {
}

func numUniqueEmails(emails []string) int {
	uniqEmails := make(map[string]bool)
	for _, addr := range emails {
		email := processAddr(addr)
		uniqEmails[email] = true
	}

	return len(uniqEmails)
}

func processAddr(addr string) string {
	bytes := []byte{}
	k := 0
	ignore := false
	for i := 0; i < len(addr); i++ {
		if addr[i] == '@' {
			k = i
			break
		}
		if addr[i] == '.' || ignore {
			continue
		}
		if addr[i] == '+' {
			ignore = true
			continue
		}
		bytes = append(bytes, addr[i])
	}

	return string(bytes) + addr[k:]
}
