package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/not_selected"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/server"
)

func (s *Server) DownloadLocator(
	x context.Context,
	r server.DownloadLocatorRequestObject,
) (server.DownloadLocatorResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if not_found.Is(e) {
		return server.DownloadLocator404JSONResponse(*clientError(e)), nil
	}

	if not_selected.Is(e) {
		return server.DownloadLocator400JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.DownloadLocator500JSONResponse(*s.answer(x, e)), nil
	}

	e = s.service.DownloadLocator(
		instance,
		r.Name,
		r.Storage,
		r.Body.Content,
		r.Body.Filename,
		r.Body.Locator,
	)

	switch {
	case validation.Is(e):
		return server.DownloadLocator400JSONResponse(*clientError(e)), nil
	case not_found.Is(e):
		return server.DownloadLocator404JSONResponse(*clientError(e)), nil
	case forbidden.Is(e):
		return server.DownloadLocator403JSONResponse(*clientError(e)), nil
	case unavailable.Is(e):
		return server.DownloadLocator503JSONResponse(*clientError(e)), nil
	case e != nil:
		return server.DownloadLocator500JSONResponse(*s.answer(x, e)), nil
	}

	status := "downloaded"

	return server.DownloadLocator200JSONResponse{Status: status}, nil
}
