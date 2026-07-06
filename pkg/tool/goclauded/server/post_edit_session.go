package server

import (
	"context"
	"errors"
	"github.com/funtimecoding/soil/pkg/constant"
	goclauded "github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
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
		if errors.Is(e, goclauded.ErrorAliasCollision) {
			return server.PostEditSession500JSONResponse(
				server.ErrorResponse{Error: e.Error()},
			), nil
		}

		return server.PostEditSession500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	return server.PostEditSession200Response{}, nil
}
