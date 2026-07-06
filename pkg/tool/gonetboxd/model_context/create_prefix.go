package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/netbox/site"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) createPrefix(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	cidr, e := r.RequireString(constant.Prefix)

	if e != nil {
		return response.Fail("prefix is required: %v", e)
	}

	var si *site.Site

	if siteName := r.GetString(constant.Site, ""); siteName != "" {
		found, f := s.client.SiteByName(siteName)

		if f != nil {
			return s.captureDetail(f)
		}

		si = found
	}

	description := r.GetString(constant.Description, "")
	p, g := s.client.CreatePrefix(cidr, si, description)

	if g != nil {
		return s.captureDetail(g)
	}

	return response.SuccessAny(convert.Prefix(p))
}
