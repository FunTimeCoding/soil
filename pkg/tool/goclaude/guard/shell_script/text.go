package shell_script

import (
	"mvdan.cc/sh/v3/syntax"
	"strings"
)

func Text(w *syntax.Word) (string, bool) {
	var b strings.Builder

	for _, part := range w.Parts {
		switch p := part.(type) {
		case *syntax.Lit:
			b.WriteString(p.Value)
		case *syntax.SglQuoted:
			b.WriteString(p.Value)
		case *syntax.DblQuoted:
			for _, inner := range p.Parts {
				l, okay := inner.(*syntax.Lit)

				if !okay {
					return "", false
				}

				b.WriteString(l.Value)
			}
		default:
			return "", false
		}
	}

	return b.String(), true
}
