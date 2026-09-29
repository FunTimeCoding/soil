package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/server"
)

func (s *Server) CreateList(
	_ context.Context,
	r server.CreateListRequestObject,
) (server.CreateListResponseObject, error) {
	description := ""

	if r.Body.Description != nil {
		description = *r.Body.Description
	}

	result, f := s.client.CreateList(r.Body.Name, description)

	if f != nil {
		return server.CreateList500JSONResponse(*s.captureDetail(f)), nil
	}

	return server.CreateList200JSONResponse{
		Identifier: int32(result.Identifier),
		Name:       result.Name,
	}, nil
}
