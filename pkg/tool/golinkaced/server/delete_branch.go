package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/server"
)

func (s *Server) DeleteBranch(
	_ context.Context,
	r server.DeleteBranchRequestObject,
) (server.DeleteBranchResponseObject, error) {
	if e := s.service.DeleteBranch(int(r.Identifier)); e != nil {
		return server.DeleteBranch500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.DeleteBranch200Response{}, nil
}
