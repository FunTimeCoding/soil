package response

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/mark3labs/mcp-go/mcp"
)

func Answer(
	x context.Context,
	r face.Reporter,
	e error,
) (*mcp.CallToolResult, error) {
	return AnswerMessage(x, r, e, "")
}
