package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
)

func (s *Server) GetTransitions(
	_ context.Context,
	r server.GetTransitionsRequestObject,
) (server.GetTransitionsResponseObject, error) {
	transitions, e := s.jira.Transitions(r.Key)

	if e != nil {
		return server.GetTransitions500JSONResponse(
			*s.captureDetail(e),
		), nil
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
