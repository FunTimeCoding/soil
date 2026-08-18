package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/strings/join"
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

	if guard, g := s.guardDocumentSourced(identifier); guard != nil {
		return guard, g
	}

	addRaw := q.GetString(constant.Add, "")
	removeRaw := q.GetString(constant.Remove, "")
	replaceRaw := q.GetString(constant.ReplaceAll, "")

	if addRaw == "" && removeRaw == "" && replaceRaw == "" {
		return response.Fail(
			"at least one of add, remove, or replace_all is required",
		)
	}

	var cleaned []string

	if replaceRaw != "" {
		tags, stripped := splitTags(replaceRaw)

		if stripped {
			cleaned = append(cleaned, tags...)
		}

		e := s.service.ReplaceTags(identifier, tags)

		if e != nil {
			return s.captureFail(e, "failed to replace tags")
		}
	}

	if addRaw != "" {
		tags, stripped := splitTags(addRaw)

		if stripped {
			cleaned = append(cleaned, tags...)
		}

		e := s.service.AddTags(identifier, tags)

		if e != nil {
			return s.captureFail(e, "failed to add tags")
		}
	}

	if removeRaw != "" {
		tags, stripped := splitTags(removeRaw)

		if stripped {
			cleaned = append(cleaned, tags...)
		}

		e := s.service.RemoveTags(identifier, tags)

		if e != nil {
			return s.captureFail(e, "failed to remove tags")
		}
	}

	m, e := s.service.GetMemory(identifier)

	if e != nil {
		return s.captureFail(e, "failed to fetch memory")
	}

	stored := fmt.Sprintf("Memory %d tags: %v", identifier, m.Tags)

	if len(cleaned) == 0 {
		return response.Success(stored)
	}

	return response.Success(
		join.NewLine(
			[]string{
				join.Space(constant.TagStripNotice, join.CommaSpace(cleaned)),
				stored,
			},
		),
	)
}
