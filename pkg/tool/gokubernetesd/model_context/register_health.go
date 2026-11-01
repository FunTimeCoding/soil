package model_context

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context/option"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerHealth() {
	s.server.AddTool(
		mcp.NewTool(
			constant.Top,
			mcp.WithDescription("Show CPU and memory usage for nodes or pods"),
			mcp.WithString(
				"resourceType",
				mcp.Required(),
				mcp.Description("Resource type: nodes or pods"),
			),
			mcp.WithString(
				"namespace",
				mcp.Description("Namespace for pods (default: default)"),
			),
			mcp.WithBoolean(
				"allNamespaces",
				mcp.Description("Show pods across all namespaces"),
			),
			mcp.WithBoolean(
				"containers",
				mcp.Description("Show per-container usage for pods"),
			),
		),
		mcp.NewTypedToolHandler(s.Top),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.Argocd,
			mcp.WithDescription(
				"ArgoCD application sync and health status. Omit name to list all applications.",
			),
			mcp.WithString(
				"name",
				mcp.Description("Application name for detail view"),
			),
			option.Unfiltered(),
		),
		mcp.NewTypedToolHandler(s.Argocd),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.Certificates,
			mcp.WithDescription(
				"Certificate status, expiry, and renewal health. Omit name to list all certificates.",
			),
			mcp.WithString(
				"name",
				mcp.Description("Certificate name for detail view"),
			),
			mcp.WithString(
				"namespace",
				mcp.Description("Namespace (required for detail view)"),
			),
			option.Unfiltered(),
		),
		mcp.NewTypedToolHandler(s.Certificates),
	)
}
