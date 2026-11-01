package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerPublication() {
	s.server.AddTool(
		mcp.NewTool(
			constant.PendingPublication,
			mcp.WithDescription(
				"Show which files a publish would commit to the repository.",
			),
		),
		mcp.NewTypedToolHandler(s.pendingPublication),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.Publish,
			mcp.WithDescription(
				"Commit every unpublished authority to the repository in one commit.",
			),
		),
		mcp.NewTypedToolHandler(s.publish),
	)
}
