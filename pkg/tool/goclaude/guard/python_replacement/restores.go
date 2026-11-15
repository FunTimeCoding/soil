package python_replacement

import (
	"github.com/funtimecoding/soil/pkg/tool/goclaude/guard/shell_script"
	"path/filepath"
	"slices"
)

func restores(
	c *shell_script.Call,
	target string,
) bool {
	var operands []string

	for _, a := range c.Words[1:] {
		if value, okay := shell_script.Text(a); okay {
			operands = append(operands, filepath.Clean(value))
		}
	}

	if len(operands) == 0 {
		return false
	}

	switch c.Name {
	case "cp", "mv":
		return operands[len(operands)-1] == target
	case "git":
		return (operands[0] == "checkout" || operands[0] == "restore") &&
			slices.Contains(operands[1:], target)
	default:
		return false
	}
}
