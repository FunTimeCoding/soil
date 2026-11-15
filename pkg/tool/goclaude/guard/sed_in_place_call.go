package guard

import (
	"github.com/funtimecoding/soil/pkg/tool/goclaude/guard/shell_script"
	"strings"
)

func sedInPlaceCall(c *shell_script.Call) bool {
	if c.Name != "sed" {
		return false
	}

	for _, argument := range c.Words[1:] {
		flag := argument.Lit()

		if strings.HasPrefix(flag, "-i") ||
			strings.HasPrefix(flag, "--in-place") {
			return true
		}
	}

	return false
}
