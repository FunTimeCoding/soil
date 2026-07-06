package page

import (
	"github.com/funtimecoding/soil/pkg/system/writer"
	"strings"
)

func markersToMacros(markdown string) string {
	var result strings.Builder
	remaining := markdown

	for {
		l := macroCommentPattern.FindStringIndex(remaining)

		if l == nil {
			result.WriteString(markdownToHTML(remaining))

			break
		}

		if l[0] > 0 {
			result.WriteString(markdownToHTML(remaining[:l[0]]))
		}

		match := remaining[l[0]:l[1]]
		groups := macroCommentPattern.FindStringSubmatch(match)
		name := groups[1]
		body := strings.TrimSpace(groups[2])

		if strings.HasPrefix(body, "```") && strings.HasSuffix(body, "```") {
			inner := strings.TrimPrefix(body, "```")
			inner = strings.TrimSuffix(inner, "```")
			inner = strings.TrimSpace(inner)
			writer.Print(
				&result,
				`<ac:structured-macro ac:name="%s"><ac:plain-text-body><![CDATA[%s]]></ac:plain-text-body></ac:structured-macro>`,
				name,
				inner,
			)
		} else {
			writer.Print(
				&result,
				`<ac:structured-macro ac:name="%s"><ac:rich-text-body>%s</ac:rich-text-body></ac:structured-macro>`,
				name,
				strings.TrimRight(markdownToHTML(body), "\n"),
			)
		}

		remaining = remaining[l[1]:]
	}

	return result.String()
}
