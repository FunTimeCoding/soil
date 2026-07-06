package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/generative/model_context/parameter"
	"github.com/funtimecoding/soil/pkg/netbox/tenant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) createDevice(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	name, f := r.RequireString(parameter.Name)

	if f != nil {
		return response.Fail("name is required: %v", f)
	}

	roleName, g := r.RequireString(constant.Role)

	if g != nil {
		return response.Fail("role is required: %v", g)
	}

	typeName, h := r.RequireString(constant.Type)

	if h != nil {
		return response.Fail("type is required: %v", h)
	}

	siteName, i := r.RequireString(constant.Site)

	if i != nil {
		return response.Fail("site is required: %v", i)
	}

	role, j := s.client.DeviceRoleByName(roleName)

	if j != nil {
		return s.captureFail(j, "device role not found")
	}

	deviceType, k := s.client.DeviceTypeByName(typeName)

	if k != nil {
		return s.captureFail(k, "device type not found")
	}

	site, l := s.client.SiteByName(siteName)

	if l != nil {
		return s.captureFail(l, "site not found")
	}

	var ten *tenant.Tenant

	if t := r.GetString(constant.Tenant, ""); t != "" {
		var n error
		ten, n = s.client.TenantByName(t)

		if n != nil {
			return s.captureFail(n, "tenant not found")
		}
	}

	result, m := s.client.CreateDevice(name, role, nil, deviceType, site, ten)

	if m != nil {
		return s.captureFail(m, "device not created")
	}

	return response.SuccessAny(convert.Device(result))
}
