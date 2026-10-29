package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) CreatePrefix(
	_ context.Context,
	r server.CreatePrefixRequestObject,
) (server.CreatePrefixResponseObject, error) {
	site := ""

	if r.Body.Site != nil {
		site = *r.Body.Site
	}

	description := ""

	if r.Body.Description != nil {
		description = *r.Body.Description
	}

	p, e := s.service.CreatePrefix(r.Body.Prefix, site, description)

	if validation.Is(e) {
		return server.CreatePrefix400JSONResponse(*clientError(e)), nil
	}

	if forbidden.Is(e) {
		return server.CreatePrefix403JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.CreatePrefix500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.CreatePrefix201JSONResponse(*convert.Prefix(p)), nil
}
