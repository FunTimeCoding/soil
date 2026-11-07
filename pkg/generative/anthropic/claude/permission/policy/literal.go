package policy

import (
	"mvdan.cc/sh/v3/syntax"
	"strings"
)

func literal(word *syntax.Word) (string, bool) {
	var b strings.Builder

	for _, part := range word.Parts {
		switch p := part.(type) {
		case *syntax.Lit:
			if strings.ContainsRune(p.Value, '\\') {
				return "", false
			}

			b.WriteString(p.Value)
		case *syntax.SglQuoted:
			if p.Dollar {
				return "", false
			}

			b.WriteString(p.Value)
		case *syntax.DblQuoted:
			for _, inner := range p.Parts {
				l, okay := inner.(*syntax.Lit)

				if !okay || strings.ContainsRune(l.Value, '\\') {
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
