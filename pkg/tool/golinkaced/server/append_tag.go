package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/server"
)

func (s *Server) AppendTag(
	_ context.Context,
	r server.AppendTagRequestObject,
) (server.AppendTagResponseObject, error) {
	result, f := s.service.AppendTag(int(r.Identifier), r.Body.Name)

	if f != nil {
		return server.AppendTag500JSONResponse(*s.captureDetail(f)), nil
	}

	return server.AppendTag200JSONResponse{
		Identifier: int32(result.Identifier),
		Name:       result.Title,
		Link:       result.Link,
	}, nil
}
