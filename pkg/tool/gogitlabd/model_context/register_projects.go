package model_context

import (
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerProjects() {
	s.server.AddTool(
		mcp.NewTool(
			constant.GetProject,
			mcp.WithDescription("Get a GitLab project by path or ID"),
			mcp.WithString(
				argumentConstant.Project,
				mcp.Required(),
				mcp.Description("Project path (owner/repo) or ID"),
			),
		),
		mcp.NewTypedToolHandler(s.GetProject),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListProjects,
			mcp.WithDescription("List accessible GitLab projects"),
			mcp.WithString("search", mcp.Description("Filter by name")),
		),
		mcp.NewTypedToolHandler(s.ListProjects),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.SearchRepositories,
			mcp.WithDescription("Search GitLab projects by name"),
			mcp.WithString(
				"query",
				mcp.Required(),
				mcp.Description("Search query"),
			),
		),
		mcp.NewTypedToolHandler(s.SearchRepositories),
	)
}
