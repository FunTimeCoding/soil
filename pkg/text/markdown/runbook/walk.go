package runbook

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/text/markdown/runbook/command"
	"github.com/funtimecoding/soil/pkg/text/markdown/runbook/section"
	"github.com/yuin/goldmark/v2/ast"
	"strings"
)

func (r *Runbook) Walk(n ast.Node) {
	var s *section.Section
	var description string

	for child := n.FirstChild(); child != nil; child = child.NextSibling() {
		switch child.Kind() {
		case ast.KindHeading:
			h := child.(*ast.Heading)
			title := extractText(r.source, h)

			if h.Level == 2 {
				s = section.New(title)
				r.Sections = append(r.Sections, s)
			} else {
				console.Format(
					"Unexpected heading level %d: %s\n",
					h.Level,
					title,
				)
			}
		case ast.KindParagraph:
			if r.Title == "" {
				r.Title = strings.TrimSuffix(
					r.Filename,
					constant.MarkdownExtension,
				)
			}

			if s == nil {
				s = section.New("Uncategorized")
				r.Sections = append(r.Sections, s)
			}

			description = extractText(r.source, child)
		case ast.KindCodeBlock:
			code := child.(*ast.CodeBlock)

			if code.CodeBlockKind == ast.CodeBlockKindFenced &&
				s != nil &&
				description != "" {
				s.Commands = append(
					s.Commands,
					command.New(description, extractCode(r.source, code)),
				)
				description = ""
			}
		}
	}
}
