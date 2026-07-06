package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListPrefixes(
	_ context.Context,
	_ server.ListPrefixesRequestObject,
) (server.ListPrefixesResponseObject, error) {
	prefixes, e := s.client.Prefixes()

	if e != nil {
		return server.ListPrefixes500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.ListPrefixes200JSONResponse(
		convert.Prefixes(prefixes),
	), nil
}
