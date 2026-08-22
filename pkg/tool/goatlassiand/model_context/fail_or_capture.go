package model_context

import (
	"github.com/funtimecoding/soil/pkg/errors/ambiguous"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/not_selected"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) failOrCapture(
	e error,
	message string,
) (*mcp.CallToolResult, error) {
	if not_found.Is(e) || ambiguous.Is(e) || not_selected.Is(e) {
		return response.Fail("%s", e)
	}

	return s.captureFail(e, message)
}
