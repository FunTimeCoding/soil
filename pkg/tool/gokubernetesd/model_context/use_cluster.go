package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func (s *Server) UseCluster(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.UseCluster,
) (*mcp.CallToolResult, error) {
	if a.Cluster == "" {
		return response.Fail("cluster is required")
	}

	if _, e := s.service.ClusterByName(a.Cluster); e != nil {
		return response.Fail(e.Error())
	}

	session := server.ClientSessionFromContext(x)

	if session == nil {
		return response.Fail("no session")
	}

	s.sessions.Store(session.SessionID(), a.Cluster)

	return response.Success("active cluster set to %s", a.Cluster)
}
