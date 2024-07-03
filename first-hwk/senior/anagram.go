package main

func Anagram(n, m string) bool {
	if len(n) != len(m) {
		return false
	}
	count := make(map[rune]int)
	for _, c := range n {
		count[c]++
	}
	for _, c := range m {
		if count[c] == 0 {
			return false
		}
		count[c]--
	}
	return true
}
