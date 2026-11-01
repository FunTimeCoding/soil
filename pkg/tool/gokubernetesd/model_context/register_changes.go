package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerChanges() {
	s.server.AddTool(
		mcp.NewTool(
			constant.Delete,
			mcp.WithDescription("Delete a Kubernetes resource"),
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
		),
		mcp.NewTypedToolHandler(s.Delete),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.RolloutRestart,
			mcp.WithDescription(
				"Restart a deployment, daemonset, or statefulset",
			),
			mcp.WithString(
				"resourceType",
				mcp.Required(),
				mcp.Description(
					"Resource type: deployment, daemonset, or statefulset",
				),
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
		),
		mcp.NewTypedToolHandler(s.RolloutRestart),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.Patch,
			mcp.WithDescription(
				"Patch a Kubernetes resource. Returns the patched resource as YAML.",
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
			mcp.WithString(
				constant.Patch,
				mcp.Required(),
				mcp.Description("Patch body as JSON"),
			),
			mcp.WithString(
				"type",
				mcp.Description(
					"Patch type: strategic, merge, or json (default: strategic)",
				),
			),
			mcp.WithBoolean(
				"dryRun",
				mcp.Description("Preview the change without applying"),
			),
		),
		mcp.NewTypedToolHandler(s.Patch),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.Apply,
			mcp.WithDescription(
				"Apply a YAML manifest to create a resource. Fails if the resource already exists unless override is true.",
			),
			mcp.WithString(
				"manifest",
				mcp.Required(),
				mcp.Description("YAML manifest"),
			),
			mcp.WithString("namespace", mcp.Description("Namespace override")),
			mcp.WithBoolean(
				"override",
				mcp.Description(
					"Allow updating existing resources (default: false)",
				),
			),
			mcp.WithBoolean(
				"dryRun",
				mcp.Description("Preview the change without applying"),
			),
		),
		mcp.NewTypedToolHandler(s.Apply),
	)
}
