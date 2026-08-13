package go_mod

import "regexp"

func ReplaceReplaces(
	mod string,
	replaces string,
) string {
	return regexp.MustCompile("(?s)replace (.*)").ReplaceAllString(
		mod,
		replaces,
	)
}
