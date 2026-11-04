package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/server"
)

func (s *Server) ListInstances(
	x context.Context,
	_ server.ListInstancesRequestObject,
) (server.ListInstancesResponseObject, error) {
	var result server.ListInstances200JSONResponse

	for _, i := range s.inventory.Instances {
		result = append(
			result,
			server.InstanceDetail{Name: i.Name, Host: i.Host},
		)
	}

	return result, nil
}
