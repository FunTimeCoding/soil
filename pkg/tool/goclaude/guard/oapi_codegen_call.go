package guard

import (
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/guard/shell_script"
)

func oapiCodegenCall(c *shell_script.Call) bool {
	return c.Name == constant.OapiCodegenCommand
}
