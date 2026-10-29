package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/generated/server"
)

func (s *Server) PostPackage(
	_ context.Context,
	r server.PostPackageRequestObject,
) (server.PostPackageResponseObject, error) {
	e := s.service.Upload(
		r.Version,
		r.Repository,
		r.Architecture,
		r.File,
		r.Body,
	)

	if validation.Is(e) {
		return server.PostPackage400JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.PostPackage500JSONResponse(*s.answer(e)), nil
	}

	return server.PostPackage201Response{}, nil
}
