package model_context

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) register() {
	s.server.AddTool(
		mcp.NewTool(
			constant.ListPackages,
			mcp.WithDescription(
				"List packages and versions from the repository index, grouped by version, repository, and architecture.",
			),
			mcp.WithString(
				generativeConstant.ParameterName,
				mcp.Description("Filter to one package name."),
			),
		),
		mcp.NewTypedToolHandler(s.listPackages),
	)
}
