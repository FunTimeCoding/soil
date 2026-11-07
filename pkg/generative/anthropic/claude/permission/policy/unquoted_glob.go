package policy

import (
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"mvdan.cc/sh/v3/syntax"
	"strings"
)

func unquotedGlob(arguments []*syntax.Word) bool {
	for _, a := range arguments {
		for _, p := range a.Parts {
			l, okay := p.(*syntax.Lit)

			if okay && strings.ContainsAny(l.Value, constant.GlobCharacters) {
				return true
			}
		}
	}

	return false
}
