package server

import (
	"context"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
)

func (s *Server) PostCompaction(
	x context.Context,
	r server.PostCompactionRequestObject,
) (server.PostCompactionResponseObject, error) {
	trigger := ""

	if r.Body.Trigger != nil {
		trigger = *r.Body.Trigger
	}

	if e := s.roster.RecordCompaction(
		r.Body.Session,
		constant.CompactionPhase(r.Body.Phase),
		trigger,
	); e != nil {
		if validation.Is(e) {
			return server.PostCompaction400JSONResponse(
				server.Error{Error: e.Error()},
			), nil
		}

		return server.PostCompaction500JSONResponse(
			*s.captureFail(x, e, soil.UnexpectedError),
		), nil
	}

	return server.PostCompaction200Response{}, nil
}
