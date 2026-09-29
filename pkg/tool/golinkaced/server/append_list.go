package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/server"
)

func (s *Server) AppendList(
	_ context.Context,
	r server.AppendListRequestObject,
) (server.AppendListResponseObject, error) {
	result, f := s.service.AppendList(int(r.Identifier), r.Body.Name)

	if f != nil {
		return server.AppendList500JSONResponse(*s.captureDetail(f)), nil
	}

	return server.AppendList200JSONResponse{
		Identifier: int32(result.Identifier),
		Name:       result.Title,
		Link:       result.Link,
	}, nil
}
