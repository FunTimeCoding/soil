package model_context

import (
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerRepository() {
	s.server.AddTool(
		mcp.NewTool(
			constant.GetRepositoryTree,
			mcp.WithDescription("List files and directories in a repository"),
			mcp.WithString(
				argumentConstant.Project,
				mcp.Required(),
				mcp.Description("Project path (owner/repo) or ID"),
			),
			mcp.WithString(
				"path",
				mcp.Description("Directory path within the repository"),
			),
			mcp.WithString(
				constant.ReferenceParameter,
				mcp.Description("Branch name, tag, or commit SHA"),
			),
			mcp.WithBoolean(
				"recursive",
				mcp.Description("List files recursively"),
			),
		),
		mcp.NewTypedToolHandler(s.GetRepositoryTree),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.GetFileContents,
			mcp.WithDescription("Read a file from a GitLab repository"),
			mcp.WithString(
				argumentConstant.Project,
				mcp.Required(),
				mcp.Description("Project path (owner/repo) or ID"),
			),
			mcp.WithString(
				"path",
				mcp.Required(),
				mcp.Description("File path within the repository"),
			),
			mcp.WithString(
				constant.ReferenceParameter,
				mcp.Description("Branch name, tag, or commit SHA"),
			),
		),
		mcp.NewTypedToolHandler(s.GetFileContents),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateBranch,
			mcp.WithDescription("Create a branch in a GitLab project"),
			mcp.WithString(
				argumentConstant.Project,
				mcp.Required(),
				mcp.Description("Project path (owner/repo) or ID"),
			),
			mcp.WithString(
				"branch",
				mcp.Required(),
				mcp.Description("New branch name"),
			),
			mcp.WithString(
				constant.ReferenceParameter,
				mcp.Required(),
				mcp.Description(
					"Source branch, tag, or commit SHA to branch from",
				),
			),
		),
		mcp.NewTypedToolHandler(s.CreateBranch),
	)
}
