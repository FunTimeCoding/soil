package server

import (
	"context"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
)

func (s *Server) GetSessionsFind(
	_ context.Context,
	r server.GetSessionsFindRequestObject,
) (server.GetSessionsFindResponseObject, error) {
	var result []server.FindMatch

	for _, m := range s.transcript.SessionsByTool(r.Params.Tool) {
		name := m.Session.Slug
		e, found, f := s.roster.FindSession(m.Session.Identifier)

		if f != nil {
			return server.GetSessionsFind500JSONResponse(
				*s.captureFail(f, soil.UnexpectedError),
			), nil
		}

		if found && e.Alias != nil {
			name = *e.Alias
		}

		result = append(
			result,
			server.FindMatch{
				Count:     m.Count,
				SessionId: m.Session.Identifier,
				Name:      name,
			},
		)
	}

	return server.GetSessionsFind200JSONResponse{Matches: result}, nil
}
