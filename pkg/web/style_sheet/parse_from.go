package style_sheet

import "strings"

func parseFrom(
	css string,
	base int,
) []*Rule {
	var result []*Rule

	for i := 0; i < len(css); {
		open := strings.IndexByte(css[i:], '{')

		if open < 0 {
			break
		}

		open += i
		end := matchingBrace(css, open)
		head := css[i:open]
		prelude := strings.TrimSpace(head)
		body := css[open+1 : end]
		offset := base + i + len(head) - len(strings.TrimLeft(head, " \t\r\n"))

		switch {
		case strings.HasPrefix(prelude, "@media"),
			strings.HasPrefix(prelude, "@supports"):
			result = append(result, parseFrom(body, base+open+1)...)
		case strings.HasPrefix(prelude, "@"):
		default:
			result = append(
				result,
				New(offset, SplitTop(prelude, ','), declarations(body)),
			)
		}

		i = end + 1
	}

	return result
}
