package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/alpine/constant"
	"github.com/funtimecoding/soil/pkg/alpine/package_server"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/convert"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/generated/server"
)

func (s *Server) GetPackages(
	_ context.Context,
	r server.GetPackagesRequestObject,
) (server.GetPackagesResponseObject, error) {
	listings, e := package_server.Indexes(constant.PackageRoot)

	if e != nil {
		return server.GetPackages500JSONResponse(
			*s.captureFail(e, "read indexes fail"),
		), nil
	}

	name := ""

	if r.Params.Name != nil {
		name = *r.Params.Name
	}

	return server.GetPackages200JSONResponse(
		convert.Listings(package_server.Filter(listings, name)),
	), nil
}
