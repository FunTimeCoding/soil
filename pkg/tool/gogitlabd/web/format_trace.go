package web

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/strings/split"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"strings"
)

func formatTrace(s string) string {
	lines := split.NewLine(stripEscapes(s))
	clocks := make([]string, 0, len(lines))
	contents := make([]string, 0, len(lines))
	dropped := false

	for _, l := range lines {
		match := constant.TracePrefixPattern.FindStringSubmatch(l)

		if match == nil {
			clocks = append(clocks, "")
			contents = append(contents, l)

			continue
		}

		content := l[len(match[0]):]

		if strings.HasPrefix(content, "section_start:") ||
			strings.HasPrefix(content, "section_end:") {
			dropped = true

			continue
		}

		if strings.HasSuffix(match[2], "+") &&
			!dropped &&
			len(contents) > 0 {
			contents[len(contents)-1] = join.Empty(
				contents[len(contents)-1],
				content,
			)

			continue
		}

		dropped = false
		clocks = append(clocks, traceClock(match[1]))
		contents = append(contents, content)
	}

	result := make([]string, 0, len(contents))

	for index, content := range contents {
		content = carriageSegment(content)

		if clocks[index] == "" {
			result = append(result, content)

			continue
		}

		result = append(result, join.Space(clocks[index], content))
	}

	return join.NewLine(result)
}
