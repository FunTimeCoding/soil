package reach

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"mvdan.cc/sh/v3/syntax"
	"strings"
)

func literalValue(w *syntax.Word) (string, bool) {
	var b strings.Builder

	for _, part := range w.Parts {
		switch p := part.(type) {
		case *syntax.Lit:
			if strings.ContainsAny(p.Value, constant.GlobCharacters) {
				return "", false
			}

			b.WriteString(p.Value)
		case *syntax.SglQuoted:
			b.WriteString(p.Value)
		case *syntax.DblQuoted:
			for _, q := range p.Parts {
				l, okay := q.(*syntax.Lit)

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
