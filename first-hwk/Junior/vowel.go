package Junior

var isVowel = map[rune]bool{
	'a': true,
	'e': true,
	'i': true,
	'o': true,
	'u': true,
	'A': true,
	'E': true,
	'I': true,
	'O': true,
	'U': true,
}

func IsVowel(symbol rune) bool {
	if _, ok := isVowel[symbol]; ok {
		return true
	}
	return false
}
