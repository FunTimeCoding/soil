package model_context

import (
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerSubscription() {
	s.server.AddTool(
		mcp.NewTool(
			constant.SubscribeMergeRequest,
			mcp.WithDescription(
				"Watch a merge request and get notified when it moves: comments, approvals, pipeline status, pushes, draft, mergeability, merged or closed. Checked every minute; each change arrives as one digest. The watch ends on its own when the merge request is merged or closed, or a day after your session leaves the roster.",
			),
			mcp.WithString(
				argumentConstant.Project,
				mcp.Required(),
				mcp.Description("Project path (owner/repo) or ID"),
			),
			mcp.WithNumber(
				"merge_request",
				mcp.Required(),
				mcp.Description("Merge request IID"),
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
					"Short name for the merge request, shown in notifications instead of group/project!12",
				),
			),
			mcp.WithString(
				constant.ParameterDelivery,
				mcp.Enum("immediate", "queued"),
				mcp.Description(
					"immediate (default) wakes the session through its channel; queued waits for its next prompt",
				),
			),
		),
		mcp.NewTypedToolHandler(s.SubscribeMergeRequest),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.UnsubscribeMergeRequest,
			mcp.WithDescription("Stop watching a merge request"),
			mcp.WithString(
				argumentConstant.Project,
				mcp.Required(),
				mcp.Description("Project path (owner/repo) or ID"),
			),
			mcp.WithNumber(
				"merge_request",
				mcp.Required(),
				mcp.Description("Merge request IID"),
			),
			mcp.WithString(
				constant.ParameterCallsign,
				mcp.Required(),
				mcp.Description("Your session callsign"),
			),
		),
		mcp.NewTypedToolHandler(s.UnsubscribeMergeRequest),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListSubscriptions,
			mcp.WithDescription(
				"List the merge requests your callsign is watching",
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
