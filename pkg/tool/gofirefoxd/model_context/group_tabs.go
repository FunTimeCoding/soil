package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gofirefoxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) GroupTabs(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.GroupTabs,
) (*mcp.CallToolResult, error) {
	if len(a.TabIdentifiers) == 0 {
		return response.Fail("tab_ids is required")
	}

	groupIdentifier, e := s.client.GroupTabs(
		a.TabIdentifiers,
		int(a.GroupIdentifier),
		a.Title,
		a.Color,
	)

	if e != nil {
		return s.captureFail(e, constant.UnexpectedError)
	}

	return response.Success("grouped into group %d", groupIdentifier)
}
