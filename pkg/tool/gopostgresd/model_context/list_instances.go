package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/model_context/argument"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/model_context/instance_response"
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

	var result []*instance_response.Response

	for _, i := range s.service.Instances() {
		result = append(
			result,
			instance_response.New(
				i.Name,
				i.Host,
				i.Port,
				i.Database,
				i.Name == active,
			),
		)
	}

	return response.SuccessAny(result)
}
