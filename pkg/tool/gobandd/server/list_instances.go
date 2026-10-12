package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/generated/server"
)

func (s *Server) ListInstances(
	_ context.Context,
	_ server.ListInstancesRequestObject,
) (server.ListInstancesResponseObject, error) {
	var result []server.Instance

	for _, i := range s.service.Instances() {
		result = append(result, server.Instance{Name: i.Name, Host: i.Host})
	}

	return server.ListInstances200JSONResponse(result), nil
}
