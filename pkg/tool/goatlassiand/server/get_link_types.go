package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
)

func (s *Server) GetLinkTypes(
	_ context.Context,
	_ server.GetLinkTypesRequestObject,
) (server.GetLinkTypesResponseObject, error) {
	types, e := s.service.LinkTypes()

	if forbidden.Is(e) {
		return server.GetLinkTypes403JSONResponse{Error: e.Error()}, nil
	}

	if unavailable.Is(e) {
		return server.GetLinkTypes503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.GetLinkTypes500JSONResponse(*s.answer(e)), nil
	}

	var result server.GetLinkTypes200JSONResponse

	for _, t := range types {
		result = append(
			result,
			server.LinkType{
				Name:    t.Name,
				Inward:  t.Inward,
				Outward: t.Outward,
			},
		)
	}

	return result, nil
}
