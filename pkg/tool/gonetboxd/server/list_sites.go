package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListSites(
	x context.Context,
	_ server.ListSitesRequestObject,
) (server.ListSitesResponseObject, error) {
	sites, e := s.service.Sites()

	if unavailable.Is(e) {
		return server.ListSites503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.ListSites500JSONResponse(*s.captureDetail(x, e)), nil
	}

	return server.ListSites200JSONResponse(convert.Sites(sites)), nil
}
