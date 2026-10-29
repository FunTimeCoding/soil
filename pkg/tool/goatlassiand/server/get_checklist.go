package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
)

func (s *Server) GetChecklist(
	_ context.Context,
	r server.GetChecklistRequestObject,
) (server.GetChecklistResponseObject, error) {
	items, e := s.service.ReadChecklist(r.Key)

	if forbidden.Is(e) {
		return server.GetChecklist403JSONResponse{Error: e.Error()}, nil
	}

	if not_found.Is(e) {
		return server.GetChecklist404JSONResponse{Error: e.Error()}, nil
	}

	if unavailable.Is(e) {
		return server.GetChecklist503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.GetChecklist500JSONResponse(*s.answer(e)), nil
	}

	return server.GetChecklist200JSONResponse(convertChecklist(items)), nil
}
