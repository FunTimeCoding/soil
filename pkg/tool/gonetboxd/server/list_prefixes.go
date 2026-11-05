package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListPrefixes(
	x context.Context,
	_ server.ListPrefixesRequestObject,
) (server.ListPrefixesResponseObject, error) {
	prefixes, e := s.service.Prefixes()

	if unavailable.Is(e) {
		return server.ListPrefixes503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.ListPrefixes500JSONResponse(*s.captureDetail(x, e)), nil
	}

	return server.ListPrefixes200JSONResponse(convert.Prefixes(prefixes)), nil
}
