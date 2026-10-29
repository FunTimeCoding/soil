package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListSites(
	_ context.Context,
	_ server.ListSitesRequestObject,
) (server.ListSitesResponseObject, error) {
	sites, e := s.service.Sites()

	if forbidden.Is(e) {
		return server.ListSites403JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.ListSites500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.ListSites200JSONResponse(convert.Sites(sites)), nil
}
