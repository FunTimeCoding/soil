package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
)

func (s *Server) GetTransitions(
	_ context.Context,
	r server.GetTransitionsRequestObject,
) (server.GetTransitionsResponseObject, error) {
	transitions, e := s.service.Transitions(r.Key)

	if forbidden.Is(e) {
		return server.GetTransitions403JSONResponse{Error: e.Error()}, nil
	}

	if not_found.Is(e) {
		return server.GetTransitions404JSONResponse{Error: e.Error()}, nil
	}

	if unavailable.Is(e) {
		return server.GetTransitions503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.GetTransitions500JSONResponse(*s.answer(e)), nil
	}

	result := make([]*server.JiraTransition, 0, len(transitions))

	for _, t := range transitions {
		a := &server.JiraTransition{Identifier: t.ID, Name: t.Name}

		if t.To.Name != "" {
			a.ToStatus = &t.To.Name
		}

		result = append(result, a)
	}

	return server.GetTransitions200JSONResponse(result), nil
}
