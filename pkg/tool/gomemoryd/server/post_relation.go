package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/server"
)

func (s *Server) PostRelation(
	x context.Context,
	r server.PostRelationRequestObject,
) (server.PostRelationResponseObject, error) {
	relationType := ""

	if r.Body.Type != nil {
		relationType = *r.Body.Type
	}

	e := s.service.CreateRelation(
		r.Body.SourceIdentifier,
		r.Body.TargetIdentifier,
		relationType,
	)

	if validation.Is(e) {
		return server.PostRelation400JSONResponse{Error: e.Error()}, nil
	}

	if not_found.Is(e) {
		return server.PostRelation404JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.PostRelation500JSONResponse(*s.answer(x, e)), nil
	}

	return server.PostRelation200Response{}, nil
}
