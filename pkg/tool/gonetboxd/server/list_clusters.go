package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListClusters(
	_ context.Context,
	_ server.ListClustersRequestObject,
) (server.ListClustersResponseObject, error) {
	clusters, e := s.client.Clusters()

	if e != nil {
		return server.ListClusters500JSONResponse(*s.captureDetail(e)), nil
	}

	result := make([]*server.Cluster, 0, len(clusters))

	for _, cl := range clusters {
		entry := &server.Cluster{Identifier: cl.Identifier, Name: cl.Name}

		if cl.Raw.Type.Name != "" {
			entry.Type = &cl.Raw.Type.Name
		}

		result = append(result, entry)
	}

	return server.ListClusters200JSONResponse(result), nil
}
