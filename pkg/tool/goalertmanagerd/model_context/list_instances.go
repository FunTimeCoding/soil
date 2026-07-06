package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/convert"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func (s *Server) listInstances(
	x context.Context,
	_ mcp.CallToolRequest,
	_ argument.ListInstances,
) (*mcp.CallToolResult, error) {
	var active string

	if session := server.ClientSessionFromContext(x); session != nil {
		active, _ = s.service.ActiveInstance(session.SessionID())
	}

	return response.SuccessAny(
		convert.Instances(s.service.Instances(), active),
	)
}
