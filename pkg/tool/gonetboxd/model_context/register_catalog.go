package model_context

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerCatalog() {
	s.server.AddTool(
		mcp.NewTool(
			constant.ListSites,
			mcp.WithDescription("List all NetBox sites"),
		),
		s.listSites,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListTenants,
			mcp.WithDescription("List all NetBox tenants"),
		),
		s.listTenants,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListDeviceRoles,
			mcp.WithDescription("List all NetBox device roles"),
		),
		s.listDeviceRoles,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListManufacturers,
			mcp.WithDescription("List all NetBox manufacturers"),
		),
		s.listManufacturers,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListDeviceTypes,
			mcp.WithDescription("List all NetBox device types"),
		),
		s.listDeviceTypes,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateSite,
			mcp.WithDescription("Create a NetBox site"),
			mcp.WithString(
				generativeConstant.ParameterName,
				mcp.Required(),
				mcp.Description("Site name"),
			),
		),
		s.createSite,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateTenant,
			mcp.WithDescription("Create a NetBox tenant"),
			mcp.WithString(
				generativeConstant.ParameterName,
				mcp.Required(),
				mcp.Description("Tenant name"),
			),
		),
		s.createTenant,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateDeviceRole,
			mcp.WithDescription("Create a NetBox device role"),
			mcp.WithString(
				generativeConstant.ParameterName,
				mcp.Required(),
				mcp.Description("Role name"),
			),
			mcp.WithBoolean(
				constant.VirtualMachineRole,
				mcp.Description(
					"Allow virtual machines to take this role. Default: false.",
				),
			),
		),
		s.createDeviceRole,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateManufacturer,
			mcp.WithDescription("Create a NetBox manufacturer"),
			mcp.WithString(
				generativeConstant.ParameterName,
				mcp.Required(),
				mcp.Description("Manufacturer name"),
			),
		),
		s.createManufacturer,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateDeviceType,
			mcp.WithDescription("Create a NetBox device type"),
			mcp.WithString(
				constant.Model,
				mcp.Required(),
				mcp.Description("Model name"),
			),
			mcp.WithString(
				constant.Manufacturer,
				mcp.Required(),
				mcp.Description("Manufacturer name (must exist)"),
			),
		),
		s.createDeviceType,
	)
}
