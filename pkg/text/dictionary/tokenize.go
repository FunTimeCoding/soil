package dictionary

import "strings"

func tokenize(
	text string,
	result map[string]bool,
) {
	for _, token := range strings.FieldsFunc(
		text,
		func(r rune) bool { return !wordRune(r) },
	) {
		result[token] = true
	}
}
