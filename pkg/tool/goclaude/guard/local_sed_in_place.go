package guard

import (
	"mvdan.cc/sh/v3/syntax"
	"path"
	"strings"
)

func localSedInPlace(command string) (bool, error) {
	found := false
	e := walkCalls(
		command,
		func(call *syntax.CallExpr) {
			if path.Base(call.Args[0].Lit()) != "sed" {
				return
			}

			for _, argument := range call.Args[1:] {
				flag := argument.Lit()

				if strings.HasPrefix(flag, "-i") ||
					strings.HasPrefix(flag, "--in-place") {
					found = true
				}
			}
		},
	)

	return found, e
}
