package reach

import (
	"mvdan.cc/sh/v3/syntax"
	"strings"
)

func hasFlag(
	words []*syntax.Word,
	prefixes ...string,
) bool {
	for _, w := range words {
		for _, p := range prefixes {
			if strings.HasPrefix(w.Lit(), p) {
				return true
			}
		}
	}

	return false
}
