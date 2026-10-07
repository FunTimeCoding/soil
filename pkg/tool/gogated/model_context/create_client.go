package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gogated/model_context/argument"
	"github.com/funtimecoding/soil/pkg/tool/gogated/types/credential"
	"github.com/funtimecoding/soil/pkg/tool/gogated/types/register_client"
	"github.com/mark3labs/mcp-go/mcp"
	"strings"
)

func (s *Server) createClient(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.CreateClient,
) (*mcp.CallToolResult, error) {
	result, e := s.service.RegisterClient(
		register_client.NewFleetRequest(
			strings.Fields(a.RedirectLocator),
			strings.Fields(a.Scope),
		),
	)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(
		credential.New(result.ClientIdentifier, result.ClientSecret),
	)
}
