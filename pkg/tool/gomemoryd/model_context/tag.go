package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) tag(
	_ context.Context,
	q mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	identifier := int64(q.GetFloat(constant.MemoryIdentifier, 0))

	if identifier == 0 {
		return response.Fail("memory_id is required")
	}

	addRaw := q.GetString(constant.Add, "")
	removeRaw := q.GetString(constant.Remove, "")
	replaceRaw := q.GetString(constant.ReplaceAll, "")

	if addRaw == "" && removeRaw == "" && replaceRaw == "" {
		return response.Fail("at least one of add, remove, or replace_all is required")
	}

	if replaceRaw != "" {
		e := s.service.ReplaceTags(identifier, splitTags(replaceRaw))

		if e != nil {
			return s.captureFail(e, "failed to replace tags")
		}
	}

	if addRaw != "" {
		e := s.service.AddTags(identifier, splitTags(addRaw))

		if e != nil {
			return s.captureFail(e, "failed to add tags")
		}
	}

	if removeRaw != "" {
		e := s.service.RemoveTags(identifier, splitTags(removeRaw))

		if e != nil {
			return s.captureFail(e, "failed to remove tags")
		}
	}

	m, e := s.service.GetMemory(identifier)

	if e != nil {
		return s.captureFail(e, "failed to fetch memory")
	}

	return response.Success(
		fmt.Sprintf("Memory %d tags: %v", identifier, m.Tags),
	)
}
