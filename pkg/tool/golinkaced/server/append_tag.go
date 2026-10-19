package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/server"
)

func (s *Server) AppendTag(
	_ context.Context,
	r server.AppendTagRequestObject,
) (server.AppendTagResponseObject, error) {
	result, f := s.service.AppendTag(int(r.Identifier), r.Body.Name)

	if not_found.Is(f) {
		return server.AppendTag404JSONResponse{Error: f.Error()}, nil
	}

	if f != nil {
		return server.AppendTag500JSONResponse(*s.answer(f)), nil
	}

	return server.AppendTag200JSONResponse{
		Identifier: int32(result.Identifier),
		Name:       result.Title,
		Link:       result.Link,
	}, nil
}
