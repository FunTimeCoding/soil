package transcript_tester

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/notation"
)

func BashUse(
	identifier string,
	cwd string,
	command string,
) string {
	return string(
		notation.Marshal(
			map[string]any{
				"type": "assistant",
				"cwd":  cwd,
				"message": map[string]any{
					"content": []map[string]any{
						{
							"type":  constant.TranscriptToolUse,
							"id":    identifier,
							"name":  constant.BashTool,
							"input": map[string]any{"command": command},
						},
					},
				},
			},
		),
	)
}
