package model_context

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerConfigurationContexts() {
	assignment := []mcp.ToolOption{
		mcp.WithArray(
			constant.Roles,
			mcp.WithStringItems(),
			mcp.Description("Role names the context applies to"),
		),
		mcp.WithArray(
			constant.Clusters,
			mcp.WithStringItems(),
			mcp.Description("Cluster names the context applies to"),
		),
		mcp.WithArray(
			constant.Platforms,
			mcp.WithStringItems(),
			mcp.Description("Platform names the context applies to"),
		),
		mcp.WithArray(
			constant.Sites,
			mcp.WithStringItems(),
			mcp.Description("Site names the context applies to"),
		),
		mcp.WithArray(
			constant.Tenants,
			mcp.WithStringItems(),
			mcp.Description("Tenant names the context applies to"),
		),
		mcp.WithArray(
			constant.Tags,
			mcp.WithStringItems(),
			mcp.Description("Tag names the context applies to"),
		),
	}
	name := mcp.WithString(
		generativeConstant.ParameterName,
		mcp.Required(),
		mcp.Description("Configuration context name"),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListConfigurationContexts,
			mcp.WithDescription(
				"List NetBox configuration contexts with their payload and what they apply to",
			),
		),
		s.listConfigurationContexts,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.GetConfigurationContext,
			mcp.WithDescription("Get one NetBox configuration context by name"),
			name,
		),
		s.getConfigurationContext,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateConfigurationContext,
			append(
				[]mcp.ToolOption{
					mcp.WithDescription(
						"Create a NetBox configuration context. Its payload is merged into every device and virtual machine it applies to; with no assignment it applies to everything.",
					),
					name,
					mcp.WithObject(
						constant.Payload,
						mcp.Required(),
						mcp.Description(
							"The JSON object the context contributes",
						),
					),
					mcp.WithNumber(
						constant.Weight,
						mcp.Description(
							"Merge order, higher weights win. Default: 1000.",
						),
					),
				},
				assignment...,
			)...,
		),
		s.createConfigurationContext,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.UpdateConfigurationContext,
			append(
				[]mcp.ToolOption{
					mcp.WithDescription(
						"Update a NetBox configuration context. A given payload replaces the whole payload; any given assignment list replaces the whole assignment.",
					),
					name,
					mcp.WithObject(
						constant.Payload,
						mcp.Description(
							"The JSON object the context contributes",
						),
					),
					mcp.WithNumber(
						constant.Weight,
						mcp.Description("Merge order, higher weights win"),
					),
				},
				assignment...,
			)...,
		),
		s.updateConfigurationContext,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.DeleteConfigurationContext,
			mcp.WithDescription("Delete a NetBox configuration context"),
			name,
		),
		s.deleteConfigurationContext,
	)
}
