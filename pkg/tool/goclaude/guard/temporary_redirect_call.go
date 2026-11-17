package guard

import (
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/guard/shell_script"
	"mvdan.cc/sh/v3/syntax"
	"strings"
)

func temporaryRedirectCall(c *shell_script.Call) bool {
	if c.Statement == nil {
		return false
	}

	for _, r := range c.Statement.Redirs {
		switch r.Op {
		case syntax.RdrOut, syntax.AppOut, syntax.RdrAll, syntax.AppAll,
			syntax.RdrClob:
		default:
			continue
		}

		if r.Word == nil {
			continue
		}

		target := strings.TrimLeft(shell_script.Printed(r.Word), `"'`)

		for _, d := range constant.TemporaryDirectories {
			if strings.HasPrefix(target, d) {
				return true
			}
		}
	}

	return false
}
