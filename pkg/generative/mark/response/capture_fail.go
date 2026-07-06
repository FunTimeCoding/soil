package response

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/mark3labs/mcp-go/mcp"
)

func CaptureFail(
	r face.Reporter,
	e error,
	message string,
) (*mcp.CallToolResult, error) {
	return FailAny(
		map[string]any{
			"error":            message,
			"event_identifier": r.CaptureException(e),
		},
	)
}
