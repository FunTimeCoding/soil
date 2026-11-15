package guard

import (
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/guard/shell_script"
)

func Hint(command string) string {
	s := shell_script.Parse(command)

	if !s.Unparsed() && s.Any(oapiCodegenCall) {
		return constant.OapiCodegenHint
	}

	return ""
}
