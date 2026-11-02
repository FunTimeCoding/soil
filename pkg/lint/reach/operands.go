package reach

import (
	"mvdan.cc/sh/v3/syntax"
	"strings"
)

func operands(words []*syntax.Word) []*syntax.Word {
	var result []*syntax.Word

	for _, w := range words {
		if !strings.HasPrefix(w.Lit(), "-") {
			result = append(result, w)
		}
	}

	return result
}
