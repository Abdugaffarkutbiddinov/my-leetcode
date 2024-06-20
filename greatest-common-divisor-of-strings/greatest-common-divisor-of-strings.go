package greatest_common_divisor_of_strings

func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}

	gcdLen := gcd(len(str1), len(str2))
	return str2[:gcdLen]
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
