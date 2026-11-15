package guard

import "github.com/funtimecoding/soil/pkg/tool/goclaude/guard/shell_script"

func npxCall(c *shell_script.Call) bool {
	return c.Name == "npx"
}
