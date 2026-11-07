package policy

import "mvdan.cc/sh/v3/syntax"

func literalWords(words []*syntax.Word) []string {
	var result []string

	for _, word := range words {
		value, okay := literal(word)

		if !okay {
			break
		}

		result = append(result, value)
	}

	return result
}
