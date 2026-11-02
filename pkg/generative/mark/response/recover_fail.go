package response

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/mark3labs/mcp-go/mcp"
)

func RecoverFail(
	r face.Reporter,
	v any,
	tool string,
) (*mcp.CallToolResult, error) {
	return FailAny(
		map[string]any{
			"error":            fmt.Sprintf("%s panicked", tool),
			"event_identifier": r.Recover(v),
		},
	)
}
