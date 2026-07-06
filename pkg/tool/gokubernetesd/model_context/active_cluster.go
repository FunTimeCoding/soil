package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/cluster"
	"github.com/mark3labs/mcp-go/server"
)

func (s *Server) activeCluster(x context.Context) (*cluster.Cluster, error) {
	session := server.ClientSessionFromContext(x)

	if session == nil {
		return nil, fmt.Errorf("no session")
	}

	v, okay := s.sessions.Load(session.SessionID())

	if !okay {
		return nil, fmt.Errorf(
			"no cluster selected - use use_cluster first",
		)
	}

	return s.service.ClusterByName(v.(string))
}
