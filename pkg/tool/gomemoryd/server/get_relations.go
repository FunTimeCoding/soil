package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	memoryConstant "github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/server"
)

func (s *Server) GetRelations(
	_ context.Context,
	r server.GetRelationsRequestObject,
) (server.GetRelationsResponseObject, error) {
	relations, e := s.service.ListRelations()

	if e != nil {
		return server.GetRelations500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	hidden, e := s.service.HiddenIdentifiers()

	if e != nil {
		return server.GetRelations500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	typeFilter := ""

	if r.Params.Type != nil {
		typeFilter = *r.Params.Type
	}

	scopeFilter := ""
	filterScope := false

	if r.Params.Scope != nil {
		filterScope = true
		scopeFilter = *r.Params.Scope

		if scopeFilter == memoryConstant.DefaultScope {
			scopeFilter = ""
		}
	}

	result := make([]server.Relation, 0, len(relations))

	for _, relation := range relations {
		if hidden[relation.SourceIdentifier] ||
			hidden[relation.TargetIdentifier] {
			continue
		}

		if typeFilter == memoryConstant.UntypedFilter &&
			relation.Type != "" {
			continue
		}

		if typeFilter != "" &&
			typeFilter != memoryConstant.UntypedFilter &&
			relation.Type != typeFilter {
			continue
		}

		if filterScope &&
			relation.SourceScope != scopeFilter &&
			relation.TargetScope != scopeFilter {
			continue
		}

		entry := server.Relation{
			SourceIdentifier: relation.SourceIdentifier,
			SourceName:       relation.SourceName,
			TargetIdentifier: relation.TargetIdentifier,
			TargetName:       relation.TargetName,
		}

		if relation.SourceScope != "" {
			entry.SourceScope = new(relation.SourceScope)
		}

		if relation.TargetScope != "" {
			entry.TargetScope = new(relation.TargetScope)
		}

		if relation.Type != "" {
			entry.Type = new(relation.Type)
		}

		if relation.CreatedAt != "" {
			entry.CreatedAt = new(relation.CreatedAt)
		}

		result = append(result, entry)
	}

	return server.GetRelations200JSONResponse(result), nil
}
