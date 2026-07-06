package server

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
)

func (s *Server) GetResolve(
	_ context.Context,
	r server.GetResolveRequestObject,
) (server.GetResolveResponseObject, error) {
	result, e := s.service.ResolveSessionIdentifier(r.Params.Query)

	if e != nil {
		return server.GetResolve500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	if result.Ambiguous() {
		var matches []server.ResolveMatch

		for _, m := range result.Matches {
			entry := server.ResolveMatch{
				Identifier: m.Identifier,
				Field:      m.Field,
			}

			if m.Name != "" {
				entry.Name = &m.Name
			}

			if m.Alias != "" {
				entry.Alias = &m.Alias
			}

			matches = append(matches, entry)
		}

		return server.GetResolve409JSONResponse{Matches: matches}, nil
	}

	if !result.Found() {
		return server.GetResolve404JSONResponse{
			Error: fmt.Sprintf("session not found: %s", r.Params.Query),
		}, nil
	}

	return server.GetResolve200JSONResponse{
		Identifier: result.Identifier(),
	}, nil
}
