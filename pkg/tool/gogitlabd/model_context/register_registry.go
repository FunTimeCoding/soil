package model_context

import (
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerRegistry() {
	s.server.AddTool(
		mcp.NewTool(
			constant.ListRegistryRepositories,
			mcp.WithDescription(
				"List container registry repositories of a GitLab project",
			),
			mcp.WithString(
				argumentConstant.Project,
				mcp.Required(),
				mcp.Description("Project path (owner/repo) or ID"),
			),
		),
		mcp.NewTypedToolHandler(s.ListRegistryRepositories),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.DeleteRegistryRepository,
			mcp.WithDescription(
				"Delete a container registry repository and its cache sibling (<path>/cache) when present - the ritual after an image moves to another registry",
			),
			mcp.WithString(
				argumentConstant.Project,
				mcp.Required(),
				mcp.Description("Project path (owner/repo) or ID"),
			),
			mcp.WithString(
				"path",
				mcp.Required(),
				mcp.Description(
					"Full registry repository path as shown by list_registry_repositories",
				),
			),
		),
		mcp.NewTypedToolHandler(s.DeleteRegistryRepository),
	)
}
