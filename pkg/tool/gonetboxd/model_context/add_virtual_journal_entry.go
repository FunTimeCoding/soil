package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) addVirtualJournalEntry(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	name, f := r.RequireString(constant.VirtualMachine)

	if f != nil {
		return response.Fail("virtual_machine is required: %v", f)
	}

	comments, g := r.RequireString(constant.Comments)

	if g != nil {
		return response.Fail("comments is required: %v", g)
	}

	kind := r.GetString(constant.Kind, "")
	result, h := s.client.AddVirtualJournalEntry(name, kind, comments)

	if h != nil {
		return s.captureFail(h, "journal entry not added to virtual machine")
	}

	return response.SuccessAny(result)
}
