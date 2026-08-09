package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/server"
)

func (s *Server) DeleteRelation(
	_ context.Context,
	r server.DeleteRelationRequestObject,
) (server.DeleteRelationResponseObject, error) {
	removed, e := s.service.DeleteRelation(
		r.Params.SourceIdentifier,
		r.Params.TargetIdentifier,
	)

	if e != nil {
		return server.DeleteRelation500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	if !removed {
		return server.DeleteRelation404JSONResponse(
			server.Error{Error: "relation not found"},
		), nil
	}

	return server.DeleteRelation200Response{}, nil
}
