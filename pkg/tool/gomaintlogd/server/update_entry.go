package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/store/entry"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/types/change"
)

func (s *Server) UpdateEntry(
	_ context.Context,
	r server.UpdateEntryRequestObject,
) (server.UpdateEntryResponseObject, error) {
	c := change.New()
	c.Action = &r.Body.Action
	c.User = &r.Body.User
	c.System = r.Body.System
	c.Service = r.Body.Service
	c.Description = r.Body.Description
	c.Timestamp = r.Body.Timestamp
	result, e := s.service.Update(uint(r.Id), c)

	if validation.Is(e) {
		return server.UpdateEntry400JSONResponse{Error: e.Error()}, nil
	}

	if not_found.Is(e) {
		return server.UpdateEntry404JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.UpdateEntry500JSONResponse(*s.answer(e)), nil
	}

	return server.UpdateEntry200JSONResponse(
		toResponse([]entry.Entry{*result})[0],
	), nil
}
