package guard

import (
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/guard/shell_script"
	"path"
	"strings"
)

func xargsUnsupportedFlagCall(c *shell_script.Call) bool {
	if path.Base(c.Expr.Args[0].Lit()) != constant.XargsCommand {
		return false
	}

	for _, argument := range c.Expr.Args[1:] {
		flag := argument.Lit()

		if strings.HasPrefix(flag, "-a") ||
			strings.HasPrefix(flag, "-d") ||
			strings.HasPrefix(flag, "--arg-file") ||
			strings.HasPrefix(flag, "--delimiter") {
			return true
		}
	}

	return false
}
