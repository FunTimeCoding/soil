package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gooutpostd/generated/server"
)

func (s *Server) GetHost(
	_ context.Context,
	_ server.GetHostRequestObject,
) (server.GetHostResponseObject, error) {
	h := s.service.Host()

	return server.GetHost200JSONResponse{
		Hostname:          h.Name,
		Platform:          h.Platform,
		HardwareAddresses: &h.HardwareAddresses,
	}, nil
}
