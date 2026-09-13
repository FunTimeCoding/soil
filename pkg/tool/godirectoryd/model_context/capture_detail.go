package model_context

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) captureDetail(e error) (*mcp.CallToolResult, error) {
	if not_found.Is(e) || conflict.Is(e) {
		return response.Fail("%s", e)
	}

	return s.captureFail(e, constant.UnexpectedError)
}
