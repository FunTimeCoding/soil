package store

import "unicode/utf8"

func alignRuneStart(
	body string,
	index int,
) int {
	for index > 0 && index < len(body) && !utf8.RuneStart(body[index]) {
		index--
	}

	return index
}
