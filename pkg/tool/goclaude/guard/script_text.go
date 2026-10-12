package guard

import (
	"mvdan.cc/sh/v3/syntax"
	"strings"
)

func scriptText(words []*syntax.Word) string {
	var b strings.Builder

	for i, w := range words {
		if i > 0 {
			b.WriteString(" ")
		}

		for _, part := range w.Parts {
			switch p := part.(type) {
			case *syntax.Lit:
				b.WriteString(p.Value)
			case *syntax.SglQuoted:
				b.WriteString(p.Value)
			case *syntax.DblQuoted:
				for _, inner := range p.Parts {
					if l, okay := inner.(*syntax.Lit); okay {
						b.WriteString(l.Value)
					} else {
						b.WriteString("_")
					}
				}
			default:
				b.WriteString("_")
			}
		}
	}

	return b.String()
}
