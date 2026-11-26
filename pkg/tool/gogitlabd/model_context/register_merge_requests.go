package model_context

import (
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerMergeRequests() {
	s.server.AddTool(
		mcp.NewTool(
			constant.ListMergeRequests,
			mcp.WithDescription(
				"List a GitLab project's merge requests, most recently updated first",
			),
			mcp.WithString(
				argumentConstant.Project,
				mcp.Required(),
				mcp.Description("Project path (owner/repo) or ID"),
			),
			mcp.WithString(
				"state",
				mcp.Description("opened (default), closed, merged or all"),
			),
			mcp.WithNumber(
				"limit",
				mcp.Description("Maximum merge requests, 1 to 100 (default 10)"),
			),
			mcp.WithNumber(
				"offset",
				mcp.Description("Merge requests to skip (default 0)"),
			),
			mcp.WithString(
				"order_by",
				mcp.Description(
					"updated_at (default) or created_at, newest first",
				),
			),
			mcp.WithString(
				"updated_after",
				mcp.Description(
					"Only merge requests updated since: a duration (7d, 36h) or RFC3339",
				),
			),
			mcp.WithString("author", mcp.Description("Author username")),
			mcp.WithString("reviewer", mcp.Description("Reviewer username")),
			mcp.WithString("assignee", mcp.Description("Assignee username")),
			mcp.WithString("source_branch", mcp.Description("Source branch")),
			mcp.WithString("target_branch", mcp.Description("Target branch")),
			mcp.WithString(
				"search",
				mcp.Description("Text in title or description"),
			),
			mcp.WithBoolean(
				"draft",
				mcp.Description("Only drafts (true) or only ready ones (false)"),
			),
		),
		mcp.NewTypedToolHandler(s.ListMergeRequests),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.GetMergeRequest,
			mcp.WithDescription("Get details of a merge request"),
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
		),
		mcp.NewTypedToolHandler(s.GetMergeRequest),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.GetMergeRequestDiffs,
			mcp.WithDescription("Get the diff of a merge request"),
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
		),
		mcp.NewTypedToolHandler(s.GetMergeRequestDiffs),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.MergeRequestDiscussions,
			mcp.WithDescription("List discussion threads on a merge request"),
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
		),
		mcp.NewTypedToolHandler(s.MergeRequestDiscussions),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateMergeRequestNote,
			mcp.WithDescription("Add a comment to a merge request"),
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
				"body",
				mcp.Required(),
				mcp.Description("Comment body"),
			),
			mcp.WithString(
				constant.ParameterCallsign,
				mcp.Description(
					"Your session callsign. Your own watch on this merge request then leaves the comment out, and other watching sessions see it under your callsign",
				),
			),
		),
		mcp.NewTypedToolHandler(s.CreateMergeRequestNote),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ReplyToDiscussion,
			mcp.WithDescription(
				"Reply inside an existing merge request discussion thread, as mr_discussions lists them, instead of a new top-level comment",
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
				"discussion",
				mcp.Required(),
				mcp.Description(
					"Discussion ID from mr_discussions, 40 hex characters",
				),
			),
			mcp.WithString(
				"body",
				mcp.Required(),
				mcp.Description("Reply body"),
			),
			mcp.WithString(
				constant.ParameterCallsign,
				mcp.Description(
					"Your session callsign. Your own watch on this merge request then leaves the reply out, and other watching sessions see it under your callsign",
				),
			),
			mcp.WithBoolean(
				"resolve",
				mcp.Description(
					"Also mark the discussion resolved after replying (default false)",
				),
			),
		),
		mcp.NewTypedToolHandler(s.ReplyToMergeRequestDiscussion),
	)
}
