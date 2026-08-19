package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) publish(
	_ context.Context,
	_ mcp.CallToolRequest,
	_ struct{},
) (*mcp.CallToolResult, error) {
	commit, change, e := s.service.Publish()

	if e != nil {
		return s.captureFail(e, constant.PublishFail)
	}

	if commit == "" {
		return response.Success("Nothing to publish")
	}

	return response.SuccessAny(
		map[string]any{
			"commit":    commit,
			"published": convert.Changes(change),
		},
	)
}
