package policy

import (
	"mvdan.cc/sh/v3/syntax"
	"strings"
)

func (w *Walk) resolve(word *syntax.Word) (string, bool) {
	var b strings.Builder

	for _, part := range word.Parts {
		switch p := part.(type) {
		case *syntax.ParamExp:
			value, okay := w.parameter(p)

			if !okay || strings.ContainsAny(value, " \t\n*?[") {
				return "", false
			}

			b.WriteString(value)
		case *syntax.DblQuoted:
			for _, inner := range p.Parts {
				value, okay := w.quotedPart(inner)

				if !okay {
					return "", false
				}

				b.WriteString(value)
			}
		default:
			value, okay := literal(&syntax.Word{Parts: []syntax.WordPart{p}})

			if !okay {
				return "", false
			}

			b.WriteString(value)
		}
	}

	return b.String(), true
}
