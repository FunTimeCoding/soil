package guard

import (
	"mvdan.cc/sh/v3/syntax"
	"path"
)

func localNpx(command string) (bool, error) {
	found := false
	e := walkCalls(
		command,
		func(call *syntax.CallExpr) {
			if path.Base(call.Args[0].Lit()) == "npx" {
				found = true
			}
		},
	)

	return found, e
}
