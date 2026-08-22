package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
)

func (s *Server) GetLinkTypes(
	_ context.Context,
	_ server.GetLinkTypesRequestObject,
) (server.GetLinkTypesResponseObject, error) {
	types, e := s.service.LinkTypes()

	if e != nil {
		return server.GetLinkTypes500JSONResponse(*s.captureDetail(e)), nil
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
