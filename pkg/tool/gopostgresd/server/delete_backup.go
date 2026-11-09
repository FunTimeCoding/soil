package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/not_selected"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/generated/server"
)

func (s *Server) DeleteBackup(
	x context.Context,
	r server.DeleteBackupRequestObject,
) (server.DeleteBackupResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if not_selected.Is(e) {
		return server.DeleteBackup400JSONResponse(*clientError(e)), nil
	}

	if not_found.Is(e) {
		return server.DeleteBackup404JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.DeleteBackup500JSONResponse(*s.answerMessage(x, e, "")), nil
	}

	e = s.backup.Delete(instance, r.Name)

	if validation.Is(e) {
		return server.DeleteBackup400JSONResponse(*clientError(e)), nil
	}

	if not_found.Is(e) {
		return server.DeleteBackup404JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.DeleteBackup500JSONResponse(*s.answerMessage(x, e, "")), nil
	}

	return server.DeleteBackup204Response{}, nil
}
