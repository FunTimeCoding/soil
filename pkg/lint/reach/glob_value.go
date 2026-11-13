package reach

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"mvdan.cc/sh/v3/syntax"
	"strings"
)

func globValue(w *syntax.Word) (string, bool) {
	var b strings.Builder
	glob := false

	for _, part := range w.Parts {
		switch p := part.(type) {
		case *syntax.Lit:
			if strings.ContainsAny(p.Value, constant.GlobCharacters) {
				glob = true
			}

			b.WriteString(p.Value)
		case *syntax.SglQuoted:
			b.WriteString(escapeGlob(p.Value))
		case *syntax.DblQuoted:
			for _, q := range p.Parts {
				l, okay := q.(*syntax.Lit)

				if !okay {
					return "", false
				}

				b.WriteString(escapeGlob(l.Value))
			}
		default:
			return "", false
		}
	}

	return b.String(), glob
}
