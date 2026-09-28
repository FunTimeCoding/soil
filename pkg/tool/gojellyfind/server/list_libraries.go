package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/generated/server"
)

func (s *Server) ListLibraries(
	_ context.Context,
	_ server.ListLibrariesRequestObject,
) (server.ListLibrariesResponseObject, error) {
	libraries, e := s.client.Libraries()

	if e != nil {
		return server.ListLibraries500JSONResponse(*s.captureDetail(e)), nil
	}

	var result server.ListLibraries200JSONResponse

	for _, l := range libraries {
		row := &server.Library{Id: l.Identifier, Name: l.Name}

		if l.CollectionType != "" {
			row.CollectionType = new(l.CollectionType)
		}

		result = append(result, row)
	}

	return result, nil
}
