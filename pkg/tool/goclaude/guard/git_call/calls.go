package git_call

import (
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/guard/shell_script"
)

func Calls(s *shell_script.Script) []*Call {
	var result []*Call

	for _, c := range s.Calls {
		if !c.Readable {
			continue
		}

		if g := Find(c); g != nil {
			result = append(result, g)
		}

		n := c.Nested

		if n == nil {
			continue
		}

		if !n.Readable && constant.GitMention.MatchString(n.Printed) {
			result = append(result, Hidden())
		}

		if n.Readable && !n.Parsed && constant.GitMention.MatchString(n.Text) {
			result = append(result, Hidden())
		}
	}

	return result
}
