package model_context

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context/option"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerObjects() {
	s.server.AddTool(
		mcp.NewTool(
			constant.Get,
			mcp.WithDescription(
				"Get or list Kubernetes resources by type, name, and namespace",
			),
			mcp.WithString(
				"resourceType",
				mcp.Required(),
				mcp.Description(
					"Resource type (e.g. pods, deployments, services, configmaps, nodes)",
				),
			),
			mcp.WithString(
				"name",
				mcp.Description("Resource name (omit to list all)"),
			),
			mcp.WithString(
				"namespace",
				mcp.Description("Namespace (default: default)"),
			),
			mcp.WithBoolean(
				"allNamespaces",
				mcp.Description("List across all namespaces"),
			),
			mcp.WithString(
				"labelSelector",
				mcp.Description("Filter by label selector (e.g. app=nginx)"),
			),
			mcp.WithString(
				"fieldSelector",
				mcp.Description(
					"Filter by field selector (e.g. metadata.name=my-pod)",
				),
			),
			option.Unfiltered(),
		),
		mcp.NewTypedToolHandler(s.Get),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.Describe,
			mcp.WithDescription(
				"Describe a Kubernetes resource with status, conditions, and events",
			),
			mcp.WithString(
				"resourceType",
				mcp.Required(),
				mcp.Description("Resource type"),
			),
			mcp.WithString(
				"name",
				mcp.Required(),
				mcp.Description("Resource name"),
			),
			mcp.WithString(
				"namespace",
				mcp.Description("Namespace (default: default)"),
			),
			option.Unfiltered(),
		),
		mcp.NewTypedToolHandler(s.Describe),
	)
}
