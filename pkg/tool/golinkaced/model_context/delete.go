package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) Delete(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.Delete,
) (*mcp.CallToolResult, error) {
	var e error

	switch a.Type {
	case "link":
		e = s.client.DeleteLink(a.Identifier)
	case "list":
		e = s.client.DeleteList(a.Identifier)
	case "tag":
		e = s.client.DeleteTag(a.Identifier)
	case "note":
		e = s.client.DeleteNote(a.Identifier)
	case "branch":
		e = s.service.DeleteBranch(a.Identifier)
	default:
		return response.Fail("unknown type: %s", a.Type)
	}

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(
		map[string]any{
			"type":       a.Type,
			"identifier": a.Identifier,
			"deleted":    true,
		},
	)
}
