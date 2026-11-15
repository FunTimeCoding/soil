package coverage_tester

import "github.com/funtimecoding/soil/pkg/generative/anthropic/claude/tool_call"

func Call(
	name string,
	timestamp string,
) *tool_call.Call {
	return tool_call.New(name, "", timestamp)
}
