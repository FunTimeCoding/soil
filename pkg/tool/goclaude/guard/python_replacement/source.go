package python_replacement

import (
	"github.com/funtimecoding/soil/pkg/tool/goclaude/guard/shell_script"
	"mvdan.cc/sh/v3/syntax"
)

func source(c *shell_script.Call) (string, bool) {
	if c.Name != "python" && c.Name != "python3" {
		return "", false
	}

	for i, argument := range c.Words[1:] {
		value := argument.Lit()

		if value == "-c" && i+2 < len(c.Words) {
			return shell_script.Text(c.Words[i+2])
		}

		if value != "-" && (value == "" || value[0] != '-') {
			return "", false
		}
	}

	for _, r := range c.Statement.Redirs {
		if r.Hdoc != nil && (r.Op == syntax.Hdoc || r.Op == syntax.DashHdoc) {
			return shell_script.Text(r.Hdoc)
		}
	}

	return "", false
}
