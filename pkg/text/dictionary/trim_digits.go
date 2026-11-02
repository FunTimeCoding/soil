package dictionary

import (
	"github.com/funtimecoding/soil/pkg/text/constant"
	"strings"
	"unicode"
)

func TrimDigits(word string) string {
	if b := strings.TrimRightFunc(word, unicode.IsDigit); len(b) >=
		constant.MinimumStem {
		return b
	}

	return word
}
