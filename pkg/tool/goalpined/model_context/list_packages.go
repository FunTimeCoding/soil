package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/system/linux/alpine/types/listing"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/convert"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) listPackages(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.ListPackages,
) (*mcp.CallToolResult, error) {
	listings, e := listing.Indexes(constant.PackageRoot)

	if e != nil {
		return s.captureFail(e, "read indexes fail")
	}

	return response.SuccessAny(
		convert.Listings(listing.Filter(listings, a.Name)),
	)
}
