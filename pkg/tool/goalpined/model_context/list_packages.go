package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/alpine/constant"
	"github.com/funtimecoding/soil/pkg/alpine/index"
	"github.com/funtimecoding/soil/pkg/alpine/package_server"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) listPackages(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.ListPackages,
) (*mcp.CallToolResult, error) {
	listings, e := package_server.Indexes(constant.PackageRoot)

	if e != nil {
		return s.captureFail(e, "read indexes fail")
	}

	if a.Name == "" {
		return response.SuccessAny(listings)
	}

	var result []*package_server.Listing

	for _, l := range listings {
		var entries []*index.Entry

		for _, entry := range l.Packages {
			if entry.Name == a.Name {
				entries = append(entries, entry)
			}
		}

		if len(entries) == 0 {
			continue
		}

		filtered := *l
		filtered.Packages = entries
		result = append(result, &filtered)
	}

	return response.SuccessAny(result)
}
