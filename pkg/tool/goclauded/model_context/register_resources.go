package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	web "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerResources() {
	s.server.AddResource(
		mcp.NewResource(
			constant.SessionWorkflowURI,
			"Session lifecycle, tool rhythm, coordination patterns",
			mcp.WithResourceDescription(
				"How to move through a goclauded session: announce, update, complete, summarize, amend. Short and long session flows. Coordination with other sessions.",
			),
			mcp.WithMIMEType(web.Markdown),
		),
		func(
			_ context.Context,
			_ mcp.ReadResourceRequest,
		) ([]mcp.ResourceContents, error) {
			return []mcp.ResourceContents{
				mcp.TextResourceContents{
					URI:      constant.SessionWorkflowURI,
					MIMEType: web.Markdown,
					Text:     constant.SessionWorkflow,
				},
			}, nil
		},
	)
}
