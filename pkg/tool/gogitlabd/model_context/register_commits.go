package model_context

import (
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerCommits() {
	s.server.AddTool(
		mcp.NewTool(
			constant.ListCommits,
			mcp.WithDescription("List commits for a GitLab project"),
			mcp.WithString(
				argumentConstant.Project,
				mcp.Required(),
				mcp.Description("Project path (owner/repo) or ID"),
			),
			mcp.WithString(
				constant.ReferenceParameter,
				mcp.Description("Branch or tag name"),
			),
		),
		mcp.NewTypedToolHandler(s.ListCommits),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.GetCommit,
			mcp.WithDescription("Get details of a commit"),
			mcp.WithString(
				argumentConstant.Project,
				mcp.Required(),
				mcp.Description("Project path (owner/repo) or ID"),
			),
			mcp.WithString(
				"sha",
				mcp.Required(),
				mcp.Description("Commit SHA"),
			),
		),
		mcp.NewTypedToolHandler(s.GetCommit),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.GetCommitDiff,
			mcp.WithDescription("Get the diff of a commit"),
			mcp.WithString(
				argumentConstant.Project,
				mcp.Required(),
				mcp.Description("Project path (owner/repo) or ID"),
			),
			mcp.WithString(
				"sha",
				mcp.Required(),
				mcp.Description("Commit SHA"),
			),
		),
		mcp.NewTypedToolHandler(s.GetCommitDiff),
	)
}
