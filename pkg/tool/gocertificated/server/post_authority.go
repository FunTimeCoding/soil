package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/convert"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/generated/server"
)

func (s *Server) PostAuthority(
	_ context.Context,
	r server.PostAuthorityRequestObject,
) (server.PostAuthorityResponseObject, error) {
	result, e := s.service.CreateAuthority(r.Body)

	if validation.Is(e) {
		return server.PostAuthority400JSONResponse{Error: e.Error()}, nil
	}

	if not_found.Is(e) {
		return server.PostAuthority404JSONResponse{Error: e.Error()}, nil
	}

	if conflict.Is(e) {
		return server.PostAuthority409JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.PostAuthority500JSONResponse(*s.answer(e)), nil
	}

	return server.PostAuthority200JSONResponse(*convert.Authority(result)), nil
}
