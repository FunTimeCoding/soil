package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gogated/convert"
	"github.com/funtimecoding/soil/pkg/tool/gogated/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
	"strings"
)

func (s *Server) updateClient(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.UpdateClient,
) (*mcp.CallToolResult, error) {
	result, e := s.service.UpdateClient(
		a.Identifier,
		strings.Fields(a.RedirectLocator),
		strings.Fields(a.Scope),
		strings.Fields(a.GrantType),
		strings.Fields(a.ResponseType),
		a.TokenEndpointAuthMethod,
	)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(convert.Client(result))
}
