package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/not_selected"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/generated/server"
)

func (s *Server) CreateBackup(
	x context.Context,
	r server.CreateBackupRequestObject,
) (server.CreateBackupResponseObject, error) {
	instance, e := s.resolveInstance(r.Body.Instance)

	if not_selected.Is(e) {
		return server.CreateBackup400JSONResponse(*clientError(e)), nil
	}

	if not_found.Is(e) {
		return server.CreateBackup404JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.CreateBackup500JSONResponse(*s.answerMessage(x, e, "")), nil
	}

	b, e := s.backup.Create(x, instance)

	if not_found.Is(e) {
		return server.CreateBackup404JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.CreateBackup500JSONResponse(*s.answerMessage(x, e, "")), nil
	}

	return server.CreateBackup200JSONResponse(convert.Backup(b)), nil
}
