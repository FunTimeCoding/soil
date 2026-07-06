package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/generative/model_context/parameter"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) createCluster(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	name, f := r.RequireString(parameter.Name)

	if f != nil {
		return response.Fail("name is required: %v", f)
	}

	typeName, g := r.RequireString(constant.Type)

	if g != nil {
		return response.Fail("type is required: %v", g)
	}

	siteName, h := r.RequireString(constant.Site)

	if h != nil {
		return response.Fail("site is required: %v", h)
	}

	t, i := s.client.ClusterTypeByName(typeName)

	if i != nil {
		return s.captureFail(i, "cluster type not found")
	}

	site, j := s.client.SiteByName(siteName)

	if j != nil {
		return s.captureFail(j, "site not found")
	}

	result, k := s.client.CreateCluster(name, t, site)

	if k != nil {
		return s.captureFail(k, "cluster not created")
	}

	return response.SuccessAny(convert.Cluster(result))
}
