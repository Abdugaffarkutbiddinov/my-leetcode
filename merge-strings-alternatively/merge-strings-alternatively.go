package merge_strings_alternatively

func mergeAlternately(word1 string, word2 string) string {

	//define the variables
	var result string
	minLength := len(word1)

	//check for the shorter string length
	if len(word2) < len(word1) {
		minLength = len(word2)
	}

	//merge upto the shorter stringlength
	for i := 0; i < minLength; i++ {
		result += string(word1[i]) + string(word2[i])
	}

	//add the remaining chars from the longer string
	result += word1[minLength:] + word2[minLength:]

	return result
}
