package policy

import (
	"mvdan.cc/sh/v3/syntax"
	"strings"
)

func (w *Walk) quotedPart(part syntax.WordPart) (string, bool) {
	switch p := part.(type) {
	case *syntax.Lit:
		if strings.ContainsRune(p.Value, '\\') {
			return "", false
		}

		return p.Value, true
	case *syntax.ParamExp:
		return w.parameter(p)
	default:
		return "", false
	}
}
