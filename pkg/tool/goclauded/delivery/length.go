package delivery

import "unicode/utf8"

func length(s string) int {
	return utf8.RuneCountInString(s)
}
