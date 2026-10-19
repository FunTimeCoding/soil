package response

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/mark3labs/mcp-go/mcp"
)

func Answer(
	r face.Reporter,
	e error,
) (*mcp.CallToolResult, error) {
	return AnswerWith(r, e, "")
}
