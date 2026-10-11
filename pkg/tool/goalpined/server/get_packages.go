package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/system/linux/alpine/types/listing"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/convert"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/generated/server"
)

func (s *Server) GetPackages(
	_ context.Context,
	r server.GetPackagesRequestObject,
) (server.GetPackagesResponseObject, error) {
	listings, e := listing.Indexes(constant.PackageRoot)

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
		convert.Listings(listing.Filter(listings, name)),
	), nil
}
