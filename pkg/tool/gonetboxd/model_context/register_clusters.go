package model_context

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerClusters() {
	s.server.AddTool(
		mcp.NewTool(
			constant.ListClusterTypes,
			mcp.WithDescription("List all NetBox cluster types"),
		),
		s.listClusterTypes,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListClusters,
			mcp.WithDescription("List all NetBox clusters"),
		),
		s.listClusters,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateClusterType,
			mcp.WithDescription("Create a NetBox cluster type"),
			mcp.WithString(
				generativeConstant.ParameterName,
				mcp.Required(),
				mcp.Description("Cluster type name"),
			),
		),
		s.createClusterType,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateCluster,
			mcp.WithDescription("Create a NetBox cluster (VM host grouping)"),
			mcp.WithString(
				generativeConstant.ParameterName,
				mcp.Required(),
				mcp.Description("Cluster name"),
			),
			mcp.WithString(
				constant.Type,
				mcp.Required(),
				mcp.Description("Cluster type name"),
			),
			mcp.WithString(
				constant.Site,
				mcp.Required(),
				mcp.Description("Site name"),
			),
		),
		s.createCluster,
	)
}
