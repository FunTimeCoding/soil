package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/chromium/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/model_context/argument"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/model_context/tab_response"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) ListTabs(
	_ context.Context,
	_ mcp.CallToolRequest,
	_ argument.ListTabs,
) (*mcp.CallToolResult, error) {
	tabs, e := s.client.Tabs()

	if e != nil {
		return s.captureDetail(e)
	}

	var result []*tab_response.Response

	for _, t := range tabs {
		if t.Type != constant.PageTabType {
			continue
		}

		result = append(
			result,
			tab_response.New(t.Identifier, t.Title, t.Locator),
		)
	}

	for _, t := range tabs {
		if t.Type != constant.IframeTabType || t.Locator == "" {
			continue
		}

		r := tab_response.New(t.Identifier, t.Title, t.Locator)
		r.Type = t.Type
		r.Parent = t.ParentIdentifier
		result = append(result, r)
	}

	return response.SuccessAny(result)
}
