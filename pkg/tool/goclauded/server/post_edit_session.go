package server

import (
	"context"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/types/argument/edit_session"
)

func (s *Server) PostEditSession(
	x context.Context,
	r server.PostEditSessionRequestObject,
) (server.PostEditSessionResponseObject, error) {
	a := edit_session.New()
	a.Alias = r.Body.Name
	a.Description = r.Body.Description
	a.Topic = r.Body.Topic
	a.Files = r.Body.Files

	if e := s.roster.EditSession(r.Body.Session, a); e != nil {
		if conflict.Is(e) {
			return server.PostEditSession409JSONResponse(
				server.Error{Error: e.Error()},
			), nil
		}

		return server.PostEditSession500JSONResponse(
			*s.captureFail(x, e, soil.UnexpectedError),
		), nil
	}

	return server.PostEditSession200Response{}, nil
}
