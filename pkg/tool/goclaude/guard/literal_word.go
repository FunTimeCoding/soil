package guard

import (
	"mvdan.cc/sh/v3/syntax"
	"strings"
)

func literalWord(w *syntax.Word) (string, bool) {
	for _, part := range w.Parts {
		l, okay := part.(*syntax.Lit)

		if okay && strings.ContainsAny(l.Value, "*?[") {
			return "", false
		}
	}

	return wordValue(w)
}
