package guard

import (
	"mvdan.cc/sh/v3/syntax"
	"path"
	"strings"
)

func localSedInPlace(command string) (bool, error) {
	file, e := syntax.NewParser().Parse(
		strings.NewReader(command),
		"",
	)

	if e != nil {
		return false, e
	}

	found := false
	syntax.Walk(
		file,
		func(node syntax.Node) bool {
			call, okay := node.(*syntax.CallExpr)

			if !okay || len(call.Args) == 0 {
				return true
			}

			if path.Base(call.Args[0].Lit()) != "sed" {
				return true
			}

			for _, argument := range call.Args[1:] {
				flag := argument.Lit()

				if strings.HasPrefix(flag, "-i") ||
					strings.HasPrefix(flag, "--in-place") {
					found = true
				}
			}

			return true
		},
	)

	return found, nil
}
