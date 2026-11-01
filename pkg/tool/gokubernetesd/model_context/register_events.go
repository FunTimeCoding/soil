package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerEvents() {
	s.server.AddTool(
		mcp.NewTool(
			constant.Events,
			mcp.WithDescription(
				"List Kubernetes events, optionally scoped to a resource",
			),
			mcp.WithString(
				"namespace",
				mcp.Description("Namespace (omit for all namespaces)"),
			),
			mcp.WithString(
				"kind",
				mcp.Description(
					"Filter by involved object kind (e.g. Pod, Deployment)",
				),
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
				constant.IncludeMuted,
				mcp.Description(
					"Include events that are muted (default false)",
				),
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
}
