package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/convert"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/generated/server"
)

func (s *Server) GetPackages(
	_ context.Context,
	r server.GetPackagesRequestObject,
) (server.GetPackagesResponseObject, error) {
	name := ""

	if r.Params.Name != nil {
		name = *r.Params.Name
	}

	listings, e := s.service.Packages(name)

	if e != nil {
		return server.GetPackages500JSONResponse(*s.answer(e)), nil
	}

	return server.GetPackages200JSONResponse(convert.Listings(listings)), nil
}
