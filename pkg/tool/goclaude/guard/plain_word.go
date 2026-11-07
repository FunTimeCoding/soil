package guard

import (
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
	"mvdan.cc/sh/v3/syntax"
	"strings"
)

func plainWord(w *syntax.Word) (string, bool) {
	for _, part := range w.Parts {
		switch p := part.(type) {
		case *syntax.Lit:
			if strings.ContainsAny(p.Value, constant.UnquotedExpansionCharacters) {
				return "", false
			}
		case *syntax.SglQuoted:
			if p.Dollar {
				return "", false
			}
		}
	}

	return wordValue(w)
}
