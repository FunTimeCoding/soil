package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/server"
)

func (s *Server) RemoveList(
	_ context.Context,
	r server.RemoveListRequestObject,
) (server.RemoveListResponseObject, error) {
	result, f := s.service.RemoveList(int(r.Identifier), r.Body.Name)

	if not_found.Is(f) {
		return server.RemoveList404JSONResponse{Error: f.Error()}, nil
	}

	if f != nil {
		return server.RemoveList500JSONResponse(*s.answer(f)), nil
	}

	return server.RemoveList200JSONResponse{
		Identifier: int32(result.Identifier),
		Name:       result.Title,
		Link:       result.Link,
	}, nil
}
