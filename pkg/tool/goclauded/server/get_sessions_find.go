package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
)

func (s *Server) GetSessionsFind(
	_ context.Context,
	r server.GetSessionsFindRequestObject,
) (server.GetSessionsFindResponseObject, error) {
	var result []server.FindMatch

	for _, m := range s.service.SessionsByTool(r.Params.Tool) {
		name := m.Session.Slug
		e, f := s.service.GetSession(m.Session.Identifier)

		if f != nil {
			return server.GetSessionsFind500JSONResponse(
				*s.captureFail(f, constant.UnexpectedError),
			), nil
		}

		if e != nil && e.Alias != nil {
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

	return server.GetSessionsFind200JSONResponse{
		Matches: result,
	}, nil
}
