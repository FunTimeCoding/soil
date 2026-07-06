package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerWrite() {
	s.server.AddTool(
		mcp.NewTool(
			constant.ExecInPod,
			mcp.WithDescription("Execute a command in a pod. Command must be an array of strings - no shell interpretation"),
			mcp.WithString(
				"name",
				mcp.Required(),
				mcp.Description("Pod name"),
			),
			mcp.WithString(
				"namespace",
				mcp.Required(),
				mcp.Description("Namespace"),
			),
			mcp.WithArray(
				"command",
				mcp.Required(),
				mcp.Description("Command as array of strings (e.g. [\"ls\", \"-la\"])"),
			),
			mcp.WithString(
				"container",
				mcp.Description("Container name"),
			),
			mcp.WithNumber(
				"timeout",
				mcp.Description("Timeout in milliseconds (default: 60000)"),
			),
		),
		mcp.NewTypedToolHandler(s.ExecInPod),
	)
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
			mcp.WithDescription("Restart a deployment, daemonset, or statefulset"),
			mcp.WithString(
				"resourceType",
				mcp.Required(),
				mcp.Description("Resource type: deployment, daemonset, or statefulset"),
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
			constant.PortForward,
			mcp.WithDescription("Forward a local port to a port on a Kubernetes resource"),
			mcp.WithString(
				"resourceType",
				mcp.Required(),
				mcp.Description("Resource type (e.g. pod, service)"),
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
			mcp.WithNumber(
				"localPort",
				mcp.Required(),
				mcp.Description("Local port"),
			),
			mcp.WithNumber(
				"targetPort",
				mcp.Required(),
				mcp.Description("Target port on the resource"),
			),
		),
		mcp.NewTypedToolHandler(s.PortForward),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.StopPortForward,
			mcp.WithDescription("Stop an active port forward"),
			mcp.WithString(
				"id",
				mcp.Required(),
				mcp.Description("Port forward ID"),
			),
		),
		mcp.NewTypedToolHandler(s.StopPortForward),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.Patch,
			mcp.WithDescription("Patch a Kubernetes resource. Returns the patched resource as YAML."),
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
				mcp.Description("Patch type: strategic, merge, or json (default: strategic)"),
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
			mcp.WithDescription("Apply a YAML manifest to create a resource. Fails if the resource already exists unless override is true."),
			mcp.WithString(
				"manifest",
				mcp.Required(),
				mcp.Description("YAML manifest"),
			),
			mcp.WithString(
				"namespace",
				mcp.Description("Namespace override"),
			),
			mcp.WithBoolean(
				"override",
				mcp.Description("Allow updating existing resources (default: false)"),
			),
			mcp.WithBoolean(
				"dryRun",
				mcp.Description("Preview the change without applying"),
			),
		),
		mcp.NewTypedToolHandler(s.Apply),
	)
}
