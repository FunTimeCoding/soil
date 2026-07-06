package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/funtimecoding/go-library/pkg/tool/gomemoryd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) get(
	_ context.Context,
	q mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	identifier := int64(q.GetFloat(constant.MemoryIdentifier, 0))

	if identifier == 0 {
		return response.Fail("memory_id is required")
	}

	m, e := s.service.GetMemory(identifier)

	if e != nil {
		return s.captureFail(e, "memory not found")
	}

	related, f := s.service.ListRelated(identifier)

	if f != nil {
		return s.captureFail(f, "failed to load relations")
	}

	result := memoryWithHistory{Memory: *m, Related: related}

	if q.GetBool(constant.IncludeHistory, false) {
		result.History, e = s.service.GetMemoryHistory(identifier)

		if e != nil {
			return s.captureFail(e, "failed to load history")
		}
	}

	return response.Success(notation.MarshalIndent(result))
}
