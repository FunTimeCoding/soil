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
			mcp.WithDescription("List merge requests for a GitLab project"),
			mcp.WithString(
				argumentConstant.Project,
				mcp.Required(),
				mcp.Description("Project path (owner/repo) or ID"),
			),
			mcp.WithString(
				"state",
				mcp.Description("Filter by state (opened, closed, merged, all)"),
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
		),
		mcp.NewTypedToolHandler(s.CreateMergeRequestNote),
	)
}
