package middle

func IsPalindorme(n string) bool {
	for i := 0; i < len(n)/2; i++ {
		if n[i] != n[len(n)-1-i] {
			return false
		}
	}
	return true
}
