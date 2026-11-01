package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerClusters() {
	s.server.AddTool(
		mcp.NewTool(
			constant.ListClusters,
			mcp.WithDescription("List available Kubernetes clusters"),
		),
		mcp.NewTypedToolHandler(s.ListClusters),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.UseCluster,
			mcp.WithDescription("Set the active cluster for this session"),
			mcp.WithString(
				"cluster",
				mcp.Required(),
				mcp.Description("Cluster name"),
			),
		),
		mcp.NewTypedToolHandler(s.UseCluster),
	)
}
