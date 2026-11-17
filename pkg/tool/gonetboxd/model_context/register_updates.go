package model_context

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerUpdates() {
	s.server.AddTool(
		mcp.NewTool(
			constant.UpdateDevice,
			mcp.WithDescription(
				"Update a NetBox device; fields left empty stay as they are",
			),
			mcp.WithString(
				constant.Device,
				mcp.Required(),
				mcp.Description("Device name"),
			),
			mcp.WithString(
				generativeConstant.ParameterName,
				mcp.Description("New device name"),
			),
			mcp.WithString(
				constant.PrimaryAddress,
				mcp.Description(
					"Address to set as primary (must be assigned to the device)",
				),
			),
			mcp.WithString(
				constant.Location,
				mcp.Description("Location name (must already exist)"),
			),
			mcp.WithString(
				constant.Platform,
				mcp.Description("Platform name (must already exist)"),
			),
			mcp.WithString(
				constant.Type,
				mcp.Description(
					"Device type model (must already exist); components stay as they are",
				),
			),
			mcp.WithString(
				constant.Role,
				mcp.Description("Device role name (must already exist)"),
			),
			mcp.WithString(
				constant.Site,
				mcp.Description("Site name (must already exist)"),
			),
			mcp.WithString(
				constant.Tenant,
				mcp.Description("Tenant name (must already exist)"),
			),
			mcp.WithString(
				constant.Status,
				mcp.Description(
					"active, offline, planned, staged, failed, inventory or decommissioning",
				),
			),
			mcp.WithString(constant.Serial, mcp.Description("Serial number")),
			mcp.WithString(
				constant.Description,
				mcp.Description("Device description"),
			),
			mcp.WithObject(
				constant.CustomFields,
				mcp.Description(
					"Custom field values by field name; fields not named keep their value",
				),
			),
			mcp.WithObject(
				constant.LocalContext,
				mcp.Description(
					"Local context data replacing the device's whole; an empty object clears it",
				),
			),
		),
		mcp.NewTypedToolHandler(s.UpdateDevice),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.UpdateVirtualMachine,
			mcp.WithDescription(
				"Update a NetBox virtual machine; fields left empty stay as they are",
			),
			mcp.WithString(
				constant.VirtualMachine,
				mcp.Required(),
				mcp.Description("Virtual machine name"),
			),
			mcp.WithString(
				generativeConstant.ParameterName,
				mcp.Description("New virtual machine name"),
			),
			mcp.WithString(
				constant.PrimaryAddress,
				mcp.Description(
					"Address to set as primary (must be assigned to the virtual machine)",
				),
			),
			mcp.WithString(
				constant.Platform,
				mcp.Description("Platform name (must already exist)"),
			),
			mcp.WithString(
				constant.Role,
				mcp.Description("Role name (must already exist)"),
			),
			mcp.WithString(
				constant.Tenant,
				mcp.Description("Tenant name (must already exist)"),
			),
			mcp.WithNumber(constant.Cores, mcp.Description("Virtual CPU count")),
			mcp.WithNumber(
				constant.Memory,
				mcp.Description("Memory in megabytes"),
			),
			mcp.WithString(
				constant.Status,
				mcp.Description(
					"active, offline, planned, staged, failed or decommissioning",
				),
			),
			mcp.WithObject(
				constant.CustomFields,
				mcp.Description(
					"Custom field values by field name; fields not named keep their value",
				),
			),
			mcp.WithObject(
				constant.LocalContext,
				mcp.Description(
					"Local context data replacing the virtual machine's whole; an empty object clears it",
				),
			),
		),
		mcp.NewTypedToolHandler(s.UpdateVirtualMachine),
	)
}
