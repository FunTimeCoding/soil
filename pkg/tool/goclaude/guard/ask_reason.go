package guard

import (
	"github.com/funtimecoding/soil/pkg/tool/goclaude/guard/git_call"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/guard/ripgrep_call"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/guard/shell_script"
	"strings"
)

func AskReason(command string) string {
	s := shell_script.Parse(command)
	var reasons []string

	if s.Parsed {
		for _, c := range git_call.Calls(s) {
			if a := c.AskReason(); a != "" {
				reasons = append(reasons, a)
			}
		}
	}

	if r := ripgrep_call.AskReason(s, command); r != "" {
		reasons = append(reasons, r)
	}

	return strings.Join(reasons, "; ")
}
