package response

import (
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/mark3labs/mcp-go/mcp"
)

func SuccessAny(v any) (*mcp.CallToolResult, error) {
	return Success(notation.MarshalIndent(v))
}
