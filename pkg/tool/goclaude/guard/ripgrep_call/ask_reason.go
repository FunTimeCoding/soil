package ripgrep_call

import (
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/guard/shell_script"
	"slices"
	"strings"
)

func AskReason(
	s *shell_script.Script,
	command string,
) string {
	configured := strings.Contains(
		command,
		constant.RipgrepConfigurationVariable,
	)
	var reasons []string

	for _, c := range s.Calls {
		r := callReason(c, configured)

		if r != "" && !slices.Contains(reasons, r) {
			reasons = append(reasons, r)
		}
	}

	if s.Unparsed() && constant.RipgrepInvocation.MatchString(command) {
		return constant.RipgrepUnreadableReason
	}

	return strings.Join(reasons, "; ")
}
