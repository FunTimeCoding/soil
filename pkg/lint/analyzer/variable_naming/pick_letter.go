package variable_naming

import "go/types"

func pickLetter(
	t types.Type,
	taken map[string]bool,
) string {
	for _, letter := range typeLetters(t) {
		if !taken[letter] {
			return letter
		}
	}

	for c := byte('a'); c <= byte('z'); c++ {
		letter := string(c)

		if !taken[letter] {
			return letter
		}
	}

	return ""
}
