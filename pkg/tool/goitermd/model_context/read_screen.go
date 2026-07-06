package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) ReadScreen(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.ReadScreen,
) (*mcp.CallToolResult, error) {
	v, e := s.client.ReadScreen(a.SessionIdentifier)

	if e != nil {
		return s.captureFail(e, constant.UnexpectedError)
	}

	return response.SuccessAny(v)
}
