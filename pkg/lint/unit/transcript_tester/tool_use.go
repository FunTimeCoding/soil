package transcript_tester

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/notation"
)

func ToolUse(
	identifier string,
	name string,
	input map[string]any,
) string {
	return string(
		notation.Marshal(
			map[string]any{
				"type": "assistant",
				"message": map[string]any{
					"content": []map[string]any{
						{
							"type":  constant.TranscriptToolUse,
							"id":    identifier,
							"name":  name,
							"input": input,
						},
					},
				},
			},
		),
	)
}
