package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/not_selected"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/generated/server"
)

func (s *Server) RestoreBackup(
	x context.Context,
	r server.RestoreBackupRequestObject,
) (server.RestoreBackupResponseObject, error) {
	instance, e := s.resolveInstance(r.Body.Instance)

	if not_selected.Is(e) {
		return server.RestoreBackup400JSONResponse(*clientError(e)), nil
	}

	if not_found.Is(e) {
		return server.RestoreBackup404JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.RestoreBackup500JSONResponse(*s.answerMessage(x, e, "")), nil
	}

	var target string

	if r.Body.Target != nil {
		target = *r.Body.Target
	}

	database, e := s.backup.Restore(x, instance, r.Name, target)

	if validation.Is(e) {
		return server.RestoreBackup400JSONResponse(*clientError(e)), nil
	}

	if not_found.Is(e) {
		return server.RestoreBackup404JSONResponse(*clientError(e)), nil
	}

	if conflict.Is(e) {
		return server.RestoreBackup409JSONResponse(*clientError(e)), nil
	}

	if unavailable.Is(e) {
		return server.RestoreBackup503JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.RestoreBackup500JSONResponse(*s.answerMessage(x, e, "")), nil
	}

	return server.RestoreBackup200JSONResponse{Database: database}, nil
}
