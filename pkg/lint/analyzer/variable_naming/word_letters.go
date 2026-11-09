package variable_naming

func wordLetters(word string) []string {
	seen := make(map[string]bool)
	var result []string

	for _, r := range word {
		letter := string(r)

		if !seen[letter] {
			seen[letter] = true
			result = append(result, letter)
		}
	}

	return result
}
