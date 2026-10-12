package guard

import (
	"mvdan.cc/sh/v3/syntax"
	"path"
	"strings"
)

func localXargsUnsupportedFlag(command string) (bool, error) {
	found := false
	e := walkCalls(
		command,
		func(call *syntax.CallExpr) {
			if path.Base(call.Args[0].Lit()) != "xargs" {
				return
			}

			for _, argument := range call.Args[1:] {
				flag := argument.Lit()

				if strings.HasPrefix(flag, "-a") ||
					strings.HasPrefix(flag, "-d") ||
					strings.HasPrefix(flag, "--arg-file") ||
					strings.HasPrefix(flag, "--delimiter") {
					found = true
				}
			}
		},
	)

	return found, e
}
