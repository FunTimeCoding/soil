package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/not_selected"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/server"
)

func (s *Server) DownloadLocator(
	_ context.Context,
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
		return server.DownloadLocator500JSONResponse(
			*s.captureFail(e, e.Error()),
		), nil
	}

	e = s.service.DownloadLocator(
		instance,
		r.Name,
		r.Storage,
		r.Body.Content,
		r.Body.Filename,
		r.Body.Locator,
	)

	if e != nil {
		return server.DownloadLocator500JSONResponse(*s.captureDetail(e)), nil
	}

	status := "downloaded"

	return server.DownloadLocator200JSONResponse{Status: status}, nil
}
