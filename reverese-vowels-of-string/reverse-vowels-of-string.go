package reverese_vowels_of_string

func isVowel(i byte) bool {
	switch i {
	case 65, 69, 73, 79, 85, 97, 101, 105, 111, 117:
		return true
	default:
		return false
	}
}
func reverseVowels(s string) string {
	var result = make([]byte, len(s))
	var last = len(s) - 1
	for i := 0; i <= last; {
		if isVowel(s[i]) && isVowel(s[last]) {
			result[i] = s[last]
			result[last] = s[i]
			i++
			last--
		} else if !isVowel(s[i]) {
			result[i] = s[i]
			i++
		} else if !isVowel(s[last]) {
			result[last] = s[last]
			last--
		}
	}
	return string(result)
}
