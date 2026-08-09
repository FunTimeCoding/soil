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
	relationType := ""

	if r.Body.Type != nil {
		relationType = *r.Body.Type
	}

	if e := s.service.CreateRelation(
		r.Body.SourceIdentifier,
		r.Body.TargetIdentifier,
		relationType,
	); e != nil {
		return server.PostRelation500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	return server.PostRelation200Response{}, nil
}
