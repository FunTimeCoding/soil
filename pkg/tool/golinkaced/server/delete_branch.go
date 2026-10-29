package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/server"
)

func (s *Server) DeleteBranch(
	_ context.Context,
	r server.DeleteBranchRequestObject,
) (server.DeleteBranchResponseObject, error) {
	e := s.service.DeleteBranch(int(r.Identifier))

	if not_found.Is(e) {
		return server.DeleteBranch404JSONResponse{Error: e.Error()}, nil
	}

	if forbidden.Is(e) {
		return server.DeleteBranch403JSONResponse{Error: e.Error()}, nil
	}

	if unavailable.Is(e) {
		return server.DeleteBranch503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.DeleteBranch500JSONResponse(*s.answer(e)), nil
	}

	return server.DeleteBranch200Response{}, nil
}
