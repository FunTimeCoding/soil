package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/server"
)

func (s *Server) DownloadLocator(
	_ context.Context,
	r server.DownloadLocatorRequestObject,
) (server.DownloadLocatorResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if e != nil {
		return server.DownloadLocator400JSONResponse(
			*clientError(e),
		), nil
	}

	c, e := s.service.Client(instance)

	if e != nil {
		return server.DownloadLocator500JSONResponse(
			*s.captureDetail(e),
		), nil
	}

	e = s.service.DownloadLocator(
		c,
		r.Name,
		r.Storage,
		r.Body.Content,
		r.Body.Filename,
		r.Body.Locator,
	)

	if e != nil {
		return server.DownloadLocator500JSONResponse(
			*s.captureDetail(e),
		), nil
	}

	status := "downloaded"

	return server.DownloadLocator200JSONResponse{
		Status: status,
	}, nil
}
