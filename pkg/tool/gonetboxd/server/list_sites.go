package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListSites(
	_ context.Context,
	_ server.ListSitesRequestObject,
) (server.ListSitesResponseObject, error) {
	sites, e := s.client.Sites()

	if e != nil {
		return server.ListSites500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.ListSites200JSONResponse(convert.Sites(sites)), nil
}
