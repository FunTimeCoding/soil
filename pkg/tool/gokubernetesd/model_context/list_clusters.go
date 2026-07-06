package model_context

import (
	"context"
	mark "github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/model_context/argument"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/response"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func (s *Server) ListClusters(
	x context.Context,
	_ mcp.CallToolRequest,
	_ argument.ListClusters,
) (*mcp.CallToolResult, error) {
	var active string

	if session := server.ClientSessionFromContext(x); session != nil {
		if v, okay := s.sessions.Load(session.SessionID()); okay {
			active = v.(string)
		}
	}

	result := []response.ClusterEntry{}

	for _, c := range s.service.Clusters() {
		result = append(
			result,
			response.ClusterEntry{
				Name:   c.Name(),
				Active: c.Name() == active,
			},
		)
	}

	return mark.SuccessAny(result)
}
