package redundant_type_name

import (
	"strings"
	"unicode"
)

func Strip(
	name string,
	affix string,
) string {
	if remainder, found := strings.CutSuffix(name, affix); found &&
		remainder != "" {
		return remainder
	}

	if remainder, found := strings.CutPrefix(name, affix); found &&
		remainder != "" &&
		unicode.IsUpper([]rune(remainder)[0]) {
		return remainder
	}

	return ""
}
