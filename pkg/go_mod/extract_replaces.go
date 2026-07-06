package go_mod

import "github.com/funtimecoding/soil/pkg/expression"

func ExtractReplaces(mod string) string {
	return expression.SubMatchIndex(`(?s)replace (.*)`, mod, 0)
}
