package page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"strings"
)

func ToMarkdown(markup string) string {
	if !constant.RichTextMacroPattern.MatchString(markup) &&
		!constant.PlainTextMacroPattern.MatchString(markup) {
		return markupToMarkdown(markup)
	}

	var parts []string
	remaining := markup

	for {
		plainLocation := constant.PlainTextMacroPattern.FindStringIndex(
			remaining,
		)
		richLocation := constant.RichTextMacroPattern.FindStringIndex(remaining)
		l, pattern := earliest(
			plainLocation,
			richLocation,
			constant.PlainTextMacroPattern,
			constant.RichTextMacroPattern,
		)

		if l == nil {
			if s := strings.TrimSpace(markupToMarkdown(remaining)); s != "" {
				parts = append(parts, s)
			}

			break
		}

		if l[0] > 0 {
			if s := strings.TrimSpace(markupToMarkdown(remaining[:l[0]])); s != "" {
				parts = append(parts, s)
			}
		}

		match := remaining[l[0]:l[1]]
		groups := pattern.FindStringSubmatch(match)
		name := groups[1]
		body := groups[2]

		if pattern == constant.PlainTextMacroPattern {
			parts = append(
				parts,
				fmt.Sprintf(
					"<!-- ac:%s -->\n```\n%s\n```\n<!-- /ac:%s -->",
					name,
					body,
					name,
				),
			)
		} else {
			parts = append(
				parts,
				fmt.Sprintf(
					"<!-- ac:%s -->\n%s\n<!-- /ac:%s -->",
					name,
					strings.TrimSpace(markupToMarkdown(body)),
					name,
				),
			)
		}

		remaining = remaining[l[1]:]
	}

	return strings.Join(parts, "\n\n")
}
