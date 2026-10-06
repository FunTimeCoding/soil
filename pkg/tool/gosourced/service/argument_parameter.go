package service

import "github.com/funtimecoding/soil/pkg/tool/gosourced/service/removal"

func argumentParameter(
	t *removal.Target,
	argument int,
) int {
	if t.Variadic >= 0 && argument >= t.Variadic {
		return t.Variadic
	}

	return argument
}
