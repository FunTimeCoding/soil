package guard

import (
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
	"mvdan.cc/sh/v3/syntax"
	"path"
)

func localOapiCodegen(command string) bool {
	found := false
	e := walkCalls(
		command,
		func(call *syntax.CallExpr) {
			if path.Base(call.Args[0].Lit()) == constant.OapiCodegenCommand {
				found = true
			}
		},
	)

	return found && e == nil
}
