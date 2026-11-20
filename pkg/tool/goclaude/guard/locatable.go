package guard

import (
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/guard/shell_script"
)

func locatable(
	s *shell_script.Script,
	command string,
) bool {
	return !s.Unparsed() &&
		!s.Any(directoryCall) &&
		!constant.GitEnvironment.MatchString(command)
}
