package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/server"
)

func (s *Server) GetStatistics(
	_ context.Context,
	_ server.GetStatisticsRequestObject,
) (server.GetStatisticsResponseObject, error) {
	scopes, e := s.service.ScopeCounts()

	if e != nil {
		return server.GetStatistics500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	tags, f := s.service.ListTags()

	if f != nil {
		return server.GetStatistics500JSONResponse(
			*s.captureFail(f, constant.UnexpectedError),
		), nil
	}

	result := server.Statistics{
		Scopes: make([]server.NamedCount, 0, len(scopes)),
		Tags:   make([]server.NamedCount, 0, len(tags)),
	}

	for _, one := range scopes {
		result.Scopes = append(
			result.Scopes,
			server.NamedCount{Name: one.Scope, Count: one.Count},
		)
	}

	membership, g := s.service.TagMembership()

	if g != nil {
		return server.GetStatistics500JSONResponse(
			*s.captureFail(g, constant.UnexpectedError),
		), nil
	}

	for _, one := range tags {
		identifiers := membership[one.Tag]
		result.Tags = append(
			result.Tags,
			server.NamedCount{
				Name:        one.Tag,
				Count:       one.Count,
				Identifiers: &identifiers,
			},
		)
	}

	return server.GetStatistics200JSONResponse(result), nil
}
