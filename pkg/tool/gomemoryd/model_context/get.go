package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) get(
	_ context.Context,
	q mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	if ids := q.GetString(constant.MemoryIdentifiers, ""); ids != "" {
		return s.getMany(ids, q.GetBool(constant.Detail, false))
	}

	identifier := int64(q.GetFloat(constant.MemoryIdentifier, 0))

	if identifier == 0 {
		return response.Fail("memory_id or memory_ids is required")
	}

	m, e := s.service.GetMemory(identifier)

	if e != nil {
		return s.captureFail(e, "memory not found")
	}

	related, f := s.service.ListRelated(identifier)

	if f != nil {
		return s.captureFail(f, "failed to load relations")
	}

	var history []store.Version

	if q.GetBool(constant.IncludeHistory, false) {
		history, e = s.service.GetMemoryHistory(identifier)

		if e != nil {
			return s.captureFail(e, "failed to load history")
		}
	}

	if q.GetBool(constant.Detail, false) {
		return response.Success(
			notation.MarshalIndent(
				memoryWithHistory{
					Memory:  *m,
					Related: related,
					History: history,
				},
			),
		)
	}

	return response.Success(
		notation.MarshalIndent(
			slimMemoryWithHistory{
				SlimMemory: *convert.Memory(m),
				Related:    related,
				History:    history,
			},
		),
	)
}
