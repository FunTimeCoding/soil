package guard

import "mvdan.cc/sh/v3/syntax"

func localPipInstall(command string) (bool, error) {
	found := false
	e := walkCalls(
		command,
		func(call *syntax.CallExpr) {
			if pipInstallCall(call) {
				found = true
			}
		},
	)

	return found, e
}
