package runbook

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/yuin/goldmark/ast"
	"strings"
)

func (r *Runbook) Walk(n ast.Node) {
	var section *Section
	var description string

	for child := n.FirstChild(); child != nil; child = child.NextSibling() {
		switch child.Kind() {
		case ast.KindHeading:
			h := child.(*ast.Heading)
			title := extractText(r.source, h)

			if h.Level == 2 {
				section = &Section{Title: title}
				r.Sections = append(r.Sections, *section)
			} else {
				fmt.Printf(
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

			if section == nil {
				section = &Section{Title: "Uncategorized"}
				r.Sections = append(r.Sections, *section)
			}

			description = extractText(r.source, child)
		case ast.KindFencedCodeBlock:
			if section != nil && description != "" {
				if len(r.Sections) > 0 {
					r.Sections[len(r.Sections)-1].Commands = append(
						r.Sections[len(r.Sections)-1].Commands,
						Command{
							Description: description,
							Code: extractCode(
								r.source,
								child.(*ast.FencedCodeBlock),
							),
						},
					)
				}

				description = ""
			}
		}
	}
}
