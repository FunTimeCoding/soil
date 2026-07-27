package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/server"
)

func (s *Server) PostRelation(
	_ context.Context,
	r server.PostRelationRequestObject,
) (server.PostRelationResponseObject, error) {
	if e := s.service.CreateRelation(
		r.Body.SourceIdentifier,
		r.Body.TargetIdentifier,
	); e != nil {
		return server.PostRelation500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	return server.PostRelation200Response{}, nil
}
