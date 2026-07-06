package server

import (
	"context"
	"errors"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/store/entry"
)

func (s *Server) UpdateEntry(
	_ context.Context,
	r server.UpdateEntryRequestObject,
) (server.UpdateEntryResponseObject, error) {
	e, f := s.store.Get(uint(r.Id))

	if f != nil {
		if errors.Is(f, not_found.Sentinel) {
			return server.UpdateEntry404JSONResponse{Error: f.Error()}, nil
		}

		return server.UpdateEntry500JSONResponse(
			*s.captureFail(f, constant.UnexpectedError),
		), nil
	}

	e.Action = r.Body.Action
	e.User = r.Body.User

	if r.Body.System != nil {
		e.System = *r.Body.System
	}

	if r.Body.Service != nil {
		e.Service = *r.Body.Service
	}

	if r.Body.Description != nil {
		e.Description = *r.Body.Description
	}

	if r.Body.Timestamp != nil {
		e.Timestamp = *r.Body.Timestamp
	}

	if f = s.store.Update(e); f != nil {
		return server.UpdateEntry500JSONResponse(
			*s.captureFail(f, constant.UnexpectedError),
		), nil
	}

	return server.UpdateEntry200JSONResponse(
		toResponse([]entry.Entry{*e})[0],
	), nil
}
