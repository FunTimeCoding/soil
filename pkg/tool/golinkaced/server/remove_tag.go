package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/server"
)

func (s *Server) RemoveTag(
	_ context.Context,
	r server.RemoveTagRequestObject,
) (server.RemoveTagResponseObject, error) {
	result, f := s.service.RemoveTag(int(r.Identifier), r.Body.Name)

	if not_found.Is(f) {
		return server.RemoveTag404JSONResponse{Error: f.Error()}, nil
	}

	if f != nil {
		return server.RemoveTag500JSONResponse(*s.answer(f)), nil
	}

	return server.RemoveTag200JSONResponse{
		Identifier: int32(result.Identifier),
		Name:       result.Title,
		Link:       result.Link,
	}, nil
}
