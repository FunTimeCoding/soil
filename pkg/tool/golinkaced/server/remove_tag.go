package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/server"
)

func (s *Server) RemoveTag(
	_ context.Context,
	r server.RemoveTagRequestObject,
) (server.RemoveTagResponseObject, error) {
	result, f := s.service.RemoveTag(int(r.Identifier), r.Body.Name)

	if f != nil {
		return server.RemoveTag500JSONResponse(*s.captureDetail(f)), nil
	}

	return server.RemoveTag200JSONResponse{
		Identifier: int32(result.Identifier),
		Name:       result.Title,
		Link:       result.Link,
	}, nil
}
