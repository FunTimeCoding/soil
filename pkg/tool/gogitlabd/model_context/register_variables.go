package model_context

import (
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerVariables() {
	s.server.AddTool(
		mcp.NewTool(
			constant.ListProjectVariables,
			mcp.WithDescription("List CI/CD variables for a GitLab project"),
			mcp.WithString(
				argumentConstant.Project,
				mcp.Required(),
				mcp.Description("Project path (owner/repo) or ID"),
			),
		),
		mcp.NewTypedToolHandler(s.ListProjectVariables),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.GetProjectVariable,
			mcp.WithDescription("Get a CI/CD variable by key"),
			mcp.WithString(
				argumentConstant.Project,
				mcp.Required(),
				mcp.Description("Project path (owner/repo) or ID"),
			),
			mcp.WithString(
				"key",
				mcp.Required(),
				mcp.Description("Variable key"),
			),
		),
		mcp.NewTypedToolHandler(s.GetProjectVariable),
	)
}
