package chunk

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"strings"
)

func findTables(content string) []*table {
	var result []*table
	lines := lineSpans(content)
	inside := false
	heading := ""

	for i := 0; i < len(lines); i++ {
		text := trimmed(content, lines[i])

		if isFence(text) {
			inside = !inside

			continue
		}

		if inside {
			continue
		}

		if constant.AnyHeadingPattern.MatchString(text) {
			heading = text

			continue
		}

		if !strings.HasPrefix(text, constant.TableRowPrefix) ||
			i+1 >= len(lines) ||
			!constant.TableSeparatorPattern.MatchString(
				trimmed(content, lines[i+1]),
			) {
			continue
		}

		t := &table{
			start:     lines[i].start,
			headerEnd: lines[i+1].end,
			heading:   heading,
			caption:   caption(content, lines, i),
		}
		j := i + 2

		for j < len(lines) &&
			strings.HasPrefix(trimmed(content, lines[j]), constant.TableRowPrefix) {
			t.rows = append(t.rows, lines[j])
			j++
		}

		t.end = lines[j-1].end
		result = append(result, t)
		i = j - 1
	}

	return result
}
