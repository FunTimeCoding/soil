package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerResources() {
	s.server.AddResource(
		mcp.NewResource(
			constant.SearchWorkflowURI,
			"Search pipeline, collections, source types, index management",
			mcp.WithResourceDescription(
				"How to use goqueryd: collection types, hybrid search pipeline, source type filtering, context system, index management.",
			),
			mcp.WithMIMEType(webConstant.Markdown),
		),
		func(
			_ context.Context,
			_ mcp.ReadResourceRequest,
		) ([]mcp.ResourceContents, error) {
			return []mcp.ResourceContents{
				mcp.TextResourceContents{
					URI:      constant.SearchWorkflowURI,
					MIMEType: webConstant.Markdown,
					Text:     constant.SearchWorkflow,
				},
			}, nil
		},
	)
}
