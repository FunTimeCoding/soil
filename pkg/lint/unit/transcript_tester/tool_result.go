package transcript_tester

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/notation"
)

func ToolResult(
	identifier string,
	text string,
) string {
	return string(
		notation.Marshal(
			map[string]any{
				"type": "user",
				"message": map[string]any{
					"content": []map[string]any{
						{
							"type":        constant.TranscriptToolResult,
							"tool_use_id": identifier,
							"content": []map[string]any{
								{"type": "text", "text": text},
							},
						},
					},
				},
			},
		),
	)
}
