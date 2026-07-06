package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	web "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerResources() {
	s.server.AddResource(
		mcp.NewResource(
			constant.MemoryWorkflowURI,
			"Memory lifecycle, profile tiers, tags, search index",
			mcp.WithResourceDescription(
				"How to use gomemoryd: profile on arrival, save and update memories, tag management, search and retrieval, goqueryd integration.",
			),
			mcp.WithMIMEType(web.Markdown),
		),
		func(
			_ context.Context,
			_ mcp.ReadResourceRequest,
		) ([]mcp.ResourceContents, error) {
			return []mcp.ResourceContents{
				mcp.TextResourceContents{
					URI:      constant.MemoryWorkflowURI,
					MIMEType: web.Markdown,
					Text:     constant.MemoryWorkflow,
				},
			}, nil
		},
	)
}
