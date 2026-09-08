package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerSubscription() {
	s.server.AddTool(
		mcp.NewTool(
			constant.SubscribeThread,
			mcp.WithDescription(
				"Watch a thread and get notified when it moves. Accepts any post in the thread; the root is resolved for you. Notifications arrive as a digest naming who posted and what changed, coalesced per thread.",
			),
			mcp.WithString(
				constant.ParameterRoot,
				mcp.Required(),
				mcp.Description("Post ID of the thread, root or any reply"),
			),
			mcp.WithString(
				constant.ParameterCallsign,
				mcp.Required(),
				mcp.Description(
					"Your session callsign, the notification address",
				),
			),
			mcp.WithString(
				constant.ParameterAlias,
				mcp.Description(
					"Short name for the thread, shown in notifications",
				),
			),
		),
		mcp.NewTypedToolHandler(s.SubscribeThread),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.UnsubscribeThread,
			mcp.WithDescription(
				"Stop watching a thread. Use when it is resolved, or too noisy to be worth the context.",
			),
			mcp.WithString(
				constant.ParameterRoot,
				mcp.Required(),
				mcp.Description("Root post ID of the thread"),
			),
			mcp.WithString(
				constant.ParameterCallsign,
				mcp.Required(),
				mcp.Description("Your session callsign"),
			),
		),
		mcp.NewTypedToolHandler(s.UnsubscribeThread),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListSubscriptions,
			mcp.WithDescription(
				"List the threads your callsign is watching, newest activity first. Resolves an alias back to its root post.",
			),
			mcp.WithString(
				constant.ParameterCallsign,
				mcp.Required(),
				mcp.Description("Your session callsign"),
			),
		),
		mcp.NewTypedToolHandler(s.ListSubscriptions),
	)
}
