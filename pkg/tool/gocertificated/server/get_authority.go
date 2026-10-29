package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/convert"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/generated/server"
)

func (s *Server) GetAuthority(
	_ context.Context,
	r server.GetAuthorityRequestObject,
) (server.GetAuthorityResponseObject, error) {
	result, e := s.service.Authority(r.Name)

	if not_found.Is(e) {
		return server.GetAuthority404JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.GetAuthority500JSONResponse(*s.answer(e)), nil
	}

	return server.GetAuthority200JSONResponse(*convert.Authority(result)), nil
}
