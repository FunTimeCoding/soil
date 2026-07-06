package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) register() {
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
	s.server.AddTool(
		mcp.NewTool(
			constant.Get,
			mcp.WithDescription("Get or list Kubernetes resources by type, name, and namespace"),
			mcp.WithString(
				"resourceType",
				mcp.Required(),
				mcp.Description("Resource type (e.g. pods, deployments, services, configmaps, nodes)"),
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
				mcp.Description("Filter by field selector (e.g. metadata.name=my-pod)"),
			),
			mcp.WithBoolean(
				"unfiltered",
				mcp.Description("Return the full object without filtering noise (managedFields, last-applied-configuration)"),
			),
		),
		mcp.NewTypedToolHandler(s.Get),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.Describe,
			mcp.WithDescription("Describe a Kubernetes resource with status, conditions, and events"),
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
			mcp.WithBoolean(
				"unfiltered",
				mcp.Description("Return the full object without filtering noise (managedFields, last-applied-configuration)"),
			),
		),
		mcp.NewTypedToolHandler(s.Describe),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.Logs,
			mcp.WithDescription("Get logs from a pod or resource that owns pods (e.g. deployment/sentry, job/migrate)"),
			mcp.WithString(
				"name",
				mcp.Required(),
				mcp.Description("Pod name or resource/name (e.g. sentry-abc123, deployment/sentry, job/migrate)"),
			),
			mcp.WithString(
				"namespace",
				mcp.Required(),
				mcp.Description("Namespace"),
			),
			mcp.WithString(
				"container",
				mcp.Description("Container name (required for multi-container pods)"),
			),
			mcp.WithNumber(
				"tail",
				mcp.Description("Number of lines from the end"),
			),
			mcp.WithString(
				"since",
				mcp.Description("Relative time (e.g. 5m, 1h)"),
			),
			mcp.WithBoolean(
				"previous",
				mcp.Description("Logs from previously terminated container"),
			),
			mcp.WithBoolean(
				"timestamps",
				mcp.Description("Include timestamps"),
			),
			mcp.WithBoolean(
				"all",
				mcp.Description("Get logs from all matching pods (default: error with list if multiple)"),
			),
		),
		mcp.NewTypedToolHandler(s.Logs),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.Events,
			mcp.WithDescription("List Kubernetes events, optionally scoped to a resource"),
			mcp.WithString(
				"namespace",
				mcp.Description("Namespace (omit for all namespaces)"),
			),
			mcp.WithString(
				"kind",
				mcp.Description("Filter by involved object kind (e.g. Pod, Deployment)"),
			),
			mcp.WithString(
				"name",
				mcp.Description("Filter by involved object name"),
			),
			mcp.WithString(
				"type",
				mcp.Description("Filter by event type: Normal or Warning"),
			),
			mcp.WithNumber(
				"limit",
				mcp.Description("Maximum events to return"),
			),
			mcp.WithBoolean(
				"unfiltered",
				mcp.Description("Show all events including muted"),
			),
		),
		mcp.NewTypedToolHandler(s.Events),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.MuteEvent,
			mcp.WithDescription("Add a mute rule to filter events by reason"),
			mcp.WithString(
				"reason",
				mcp.Required(),
				mcp.Description("Event reason to mute (exact match)"),
			),
			mcp.WithString(
				"message",
				mcp.Description("Optional message pattern (contains match)"),
			),
			mcp.WithString(
				"cluster",
				mcp.Description("Cluster name (omit for all clusters)"),
			),
		),
		mcp.NewTypedToolHandler(s.MuteEvent),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.UnmuteEvent,
			mcp.WithDescription("Remove a mute rule by identifier"),
			mcp.WithNumber(
				"identifier",
				mcp.Required(),
				mcp.Description("Mute rule identifier from list_muted_events"),
			),
		),
		mcp.NewTypedToolHandler(s.UnmuteEvent),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListMutedEvents,
			mcp.WithDescription("List all event mute rules"),
		),
		mcp.NewTypedToolHandler(s.ListMutedEvents),
	)
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
			mcp.WithDescription("ArgoCD application sync and health status. Omit name to list all applications."),
			mcp.WithString(
				"name",
				mcp.Description("Application name for detail view"),
			),
			mcp.WithBoolean(
				"unfiltered",
				mcp.Description("Show all resources even when synced"),
			),
		),
		mcp.NewTypedToolHandler(s.Argocd),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.Certificates,
			mcp.WithDescription("Certificate status, expiry, and renewal health. Omit name to list all certificates."),
			mcp.WithString(
				"name",
				mcp.Description("Certificate name for detail view"),
			),
			mcp.WithString(
				"namespace",
				mcp.Description("Namespace (required for detail view)"),
			),
			mcp.WithBoolean(
				"unfiltered",
				mcp.Description("Return the full object without filtering"),
			),
		),
		mcp.NewTypedToolHandler(s.Certificates),
	)
}
