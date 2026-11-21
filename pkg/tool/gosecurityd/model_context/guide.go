package model_context

import (
	"context"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) guide(
	locator string,
	name string,
	description string,
	text string,
) {
	s.server.AddResource(
		mcp.NewResource(
			locator,
			name,
			mcp.WithResourceDescription(description),
			mcp.WithMIMEType(webConstant.Markdown),
		),
		func(
			context.Context,
			mcp.ReadResourceRequest,
		) ([]mcp.ResourceContents, error) {
			return []mcp.ResourceContents{
				mcp.TextResourceContents{
					URI:      locator,
					MIMEType: webConstant.Markdown,
					Text:     text,
				},
			}, nil
		},
	)
}
