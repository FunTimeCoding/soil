package markup_tester

import (
	"github.com/funtimecoding/soil/pkg/web/style_sheet"
	"strings"
)

func Specificity(selector string) [3]int {
	var result [3]int
	s := selector

	for i := 0; i < len(s); {
		c := s[i]

		switch {
		case c == '#':
			result[0]++
			i = nameEnd(s, i+1)
		case c == '.':
			result[1]++
			i = nameEnd(s, i+1)
		case c == '[':
			result[1]++
			i = closing(s, i) + 1
		case c == ':' && i+1 < len(s) && s[i+1] == ':':
			result[2]++
			i = nameEnd(s, i+2)
		case c == ':':
			end := nameEnd(s, i+1)
			name := s[i+1 : end]

			if end >= len(s) || s[end] != '(' {
				result[1]++
				i = end

				continue
			}

			last := closing(s, end)

			switch name {
			case "where":
			case "is", "not", "has":
				var weightiest [3]int

				for _, a := range style_sheet.SplitTop(s[end+1:last], ',') {
					w := Specificity(strings.TrimLeft(a, " >+~"))

					if heavier(w, weightiest) {
						weightiest = w
					}
				}

				for k := range result {
					result[k] += weightiest[k]
				}
			default:
				result[1]++
			}

			i = last + 1
		case c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z':
			result[2]++
			i = nameEnd(s, i)
		default:
			i++
		}
	}

	return result
}
