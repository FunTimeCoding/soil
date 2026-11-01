package style_sheet

import "strings"

func SplitTop(
	s string,
	separator byte,
) []string {
	var result []string
	depth := 0
	start := 0

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(', '[':
			depth++
		case ')', ']':
			depth--
		case separator:
			if depth == 0 {
				result = append(result, strings.TrimSpace(s[start:i]))
				start = i + 1
			}
		}
	}

	return append(result, strings.TrimSpace(s[start:]))
}
