package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) CreateSite(
	x context.Context,
	r server.CreateSiteRequestObject,
) (server.CreateSiteResponseObject, error) {
	i, e := s.service.CreateSite(r.Body.Name)

	if validation.Is(e) {
		return server.CreateSite400JSONResponse(*clientError(e)), nil
	}

	if unavailable.Is(e) {
		return server.CreateSite503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.CreateSite500JSONResponse(*s.captureDetail(x, e)), nil
	}

	return server.CreateSite201JSONResponse{
		Identifier: i.Identifier,
		Name:       i.Name,
	}, nil
}
