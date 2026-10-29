package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/server"
)

func (s *Server) AppendList(
	_ context.Context,
	r server.AppendListRequestObject,
) (server.AppendListResponseObject, error) {
	result, f := s.service.AppendList(int(r.Identifier), r.Body.Name)

	if not_found.Is(f) {
		return server.AppendList404JSONResponse{Error: f.Error()}, nil
	}

	if forbidden.Is(f) {
		return server.AppendList403JSONResponse{Error: f.Error()}, nil
	}

	if unavailable.Is(f) {
		return server.AppendList503JSONResponse{Error: f.Error()}, nil
	}

	if f != nil {
		return server.AppendList500JSONResponse(*s.answer(f)), nil
	}

	return server.AppendList200JSONResponse{
		Identifier: int32(result.Identifier),
		Name:       result.Title,
		Link:       result.Link,
	}, nil
}
