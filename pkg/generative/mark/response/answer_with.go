package response

import (
	"github.com/funtimecoding/soil/pkg/errors/classify"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/mark3labs/mcp-go/mcp"
)

func AnswerWith(
	r face.Reporter,
	e error,
	message string,
) (*mcp.CallToolResult, error) {
	text := classify.Message(e, message)

	if !classify.Reportable(e) {
		return Fail("%s", text)
	}

	return FailAny(
		map[string]any{
			"error":            text,
			"event_identifier": r.CaptureException(e),
		},
	)
}
