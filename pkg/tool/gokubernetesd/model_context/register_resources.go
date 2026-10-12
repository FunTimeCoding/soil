package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerResources() {
	s.server.AddResource(
		mcp.NewResource(
			constant.KubernetesWorkflowURI,
			"Cluster selection, response format, log resolution",
			mcp.WithResourceDescription(
				"Non-obvious behaviors: YAML responses, filtered comments, deployment/job log resolution, multi-pod and multi-container disambiguation.",
			),
			mcp.WithMIMEType(webConstant.Markdown),
		),
		func(
			_ context.Context,
			_ mcp.ReadResourceRequest,
		) ([]mcp.ResourceContents, error) {
			return []mcp.ResourceContents{
				mcp.TextResourceContents{
					URI:      constant.KubernetesWorkflowURI,
					MIMEType: webConstant.Markdown,
					Text:     constant.KubernetesWorkflow,
				},
			}, nil
		},
	)
}
