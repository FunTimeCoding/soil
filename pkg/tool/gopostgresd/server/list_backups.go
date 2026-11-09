package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/not_selected"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/generated/server"
)

func (s *Server) ListBackups(
	x context.Context,
	r server.ListBackupsRequestObject,
) (server.ListBackupsResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if not_selected.Is(e) {
		return server.ListBackups400JSONResponse(*clientError(e)), nil
	}

	if not_found.Is(e) {
		return server.ListBackups404JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.ListBackups500JSONResponse(*s.answerMessage(x, e, "")), nil
	}

	listed, e := s.backup.List(instance)

	if e != nil {
		return server.ListBackups500JSONResponse(*s.answerMessage(x, e, "")), nil
	}

	return server.ListBackups200JSONResponse(convert.Backups(listed)), nil
}
