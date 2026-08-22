package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/argument/edit_session"
)

func (s *Server) PostEditSession(
	_ context.Context,
	r server.PostEditSessionRequestObject,
) (server.PostEditSessionResponseObject, error) {
	a := edit_session.New()
	a.Alias = r.Body.Name
	a.Description = r.Body.Description
	a.Topic = r.Body.Topic
	a.Files = r.Body.Files

	if e := s.service.EditSession(r.Body.Session, a); e != nil {
		if conflict.Is(e) {
			return server.PostEditSession409JSONResponse(
				server.Error{Error: e.Error()},
			), nil
		}

		return server.PostEditSession500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	return server.PostEditSession200Response{}, nil
}
