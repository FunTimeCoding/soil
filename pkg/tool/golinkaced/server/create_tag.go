package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/server"
)

func (s *Server) CreateTag(
	_ context.Context,
	r server.CreateTagRequestObject,
) (server.CreateTagResponseObject, error) {
	result, f := s.client.CreateTag(r.Body.Name)

	if f != nil {
		return server.CreateTag500JSONResponse(*s.captureDetail(f)), nil
	}

	return server.CreateTag200JSONResponse{
		Identifier: int32(result.Identifier),
		Name:       result.Name,
	}, nil
}
