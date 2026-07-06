package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) ReadHistory(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.ReadHistory,
) (*mcp.CallToolResult, error) {
	count := int(a.Count)

	if count <= 0 {
		count = 50
	}

	v, e := s.client.ReadHistory(a.SessionIdentifier, count)

	if e != nil {
		return s.captureFail(e, constant.UnexpectedError)
	}

	return response.SuccessAny(v)
}
