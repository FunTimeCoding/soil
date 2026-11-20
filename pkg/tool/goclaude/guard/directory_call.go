package guard

import (
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/guard/shell_script"
	"slices"
)

func directoryCall(c *shell_script.Call) bool {
	return slices.Contains(constant.DirectoryCommands, c.Name)
}
