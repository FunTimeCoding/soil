package guard

import (
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
	"mvdan.cc/sh/v3/syntax"
	"path"
	"slices"
	"strings"
)

func scriptWords(call *syntax.CallExpr) []*syntax.Word {
	base := path.Base(call.Args[0].Lit())

	if base == constant.EvalCommand {
		return call.Args[1:]
	}

	if !slices.Contains(constant.ScriptShells, base) || len(call.Args) < 3 {
		return nil
	}

	for i, w := range call.Args[1 : len(call.Args)-1] {
		flag := w.Lit()

		if strings.HasPrefix(flag, "-") && !strings.HasPrefix(flag, "--") &&
			strings.Contains(flag, "c") {
			return call.Args[i+2 : i+3]
		}
	}

	return nil
}
