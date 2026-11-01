package model_context

import (
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerVariableChanges() {
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateProjectVariable,
			mcp.WithDescription("Create a CI/CD variable"),
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
			mcp.WithString(
				"value",
				mcp.Required(),
				mcp.Description("Variable value"),
			),
			mcp.WithBoolean(
				"protected",
				mcp.Description("Only expose to protected branches/tags"),
			),
			mcp.WithBoolean(
				"masked",
				mcp.Description("Mask value in job logs"),
			),
			mcp.WithBoolean(
				"expand",
				mcp.Description(
					"Expand $VARIABLE references in the value; enable only when the value actually contains a reference to expand (default: raw)",
				),
			),
		),
		mcp.NewTypedToolHandler(s.CreateProjectVariable),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.UpdateProjectVariable,
			mcp.WithDescription("Update a CI/CD variable"),
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
			mcp.WithString(
				"value",
				mcp.Required(),
				mcp.Description("New variable value"),
			),
			mcp.WithBoolean(
				"protected",
				mcp.Description("Only expose to protected branches/tags"),
			),
			mcp.WithBoolean(
				"masked",
				mcp.Description("Mask value in job logs"),
			),
			mcp.WithBoolean(
				"expand",
				mcp.Description(
					"Expand $VARIABLE references in the value; enable only when the value actually contains a reference to expand (default: raw)",
				),
			),
		),
		mcp.NewTypedToolHandler(s.UpdateProjectVariable),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.DeleteProjectVariable,
			mcp.WithDescription("Delete a CI/CD variable"),
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
		mcp.NewTypedToolHandler(s.DeleteProjectVariable),
	)
}
