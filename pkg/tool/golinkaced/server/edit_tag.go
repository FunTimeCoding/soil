package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/server"
)

func (s *Server) EditTag(
	_ context.Context,
	r server.EditTagRequestObject,
) (server.EditTagResponseObject, error) {
	patch := map[string]any{}

	if r.Body.Name != nil {
		patch["name"] = *r.Body.Name
	}

	result, f := s.client.UpdateTag(int(r.Identifier), patch)

	if f != nil {
		return server.EditTag500JSONResponse(*s.captureDetail(f)), nil
	}

	return server.EditTag200JSONResponse{
		Identifier: int32(result.Identifier),
		Name:       result.Name,
	}, nil
}
