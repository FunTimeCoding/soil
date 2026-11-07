package guard

import (
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
	"mvdan.cc/sh/v3/syntax"
	"strings"
)

func leadingCharacter(w *syntax.Word) (byte, bool) {
	if len(w.Parts) == 0 {
		return 0, false
	}

	switch p := w.Parts[0].(type) {
	case *syntax.Lit:
		if p.Value == "" ||
			strings.IndexByte(constant.UnquotedExpansionCharacters, p.Value[0]) >= 0 {
			return 0, false
		}

		return p.Value[0], true
	case *syntax.SglQuoted:
		if p.Dollar || p.Value == "" {
			return 0, false
		}

		return p.Value[0], true
	case *syntax.DblQuoted:
		if len(p.Parts) == 0 {
			return 0, false
		}

		l, okay := p.Parts[0].(*syntax.Lit)

		if !okay || l.Value == "" {
			return 0, false
		}

		return l.Value[0], true
	default:
		return 0, false
	}
}
