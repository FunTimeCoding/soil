package service

import (
	"github.com/dave/dst"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/removal"
)

func removeCallArguments(
	c *dst.CallExpr,
	t *removal.Target,
) {
	var arguments []dst.Expr

	for i, a := range c.Args {
		if !t.Indices[argumentParameter(t, i)] {
			arguments = append(arguments, a)
		}
	}

	if t.Variadic >= 0 && t.Indices[t.Variadic] {
		c.Ellipsis = false
	}

	c.Args = arguments
}
