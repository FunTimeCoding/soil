package model_context

import (
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/convert"
	memoryResponse "github.com/funtimecoding/soil/pkg/tool/gomemoryd/model_context/response"
	"github.com/mark3labs/mcp-go/mcp"
	"strconv"
	"strings"
)

func (s *Server) getMany(
	raw string,
	detail bool,
) (*mcp.CallToolResult, error) {
	var result []any

	for _, part := range strings.Split(raw, constant.Comma) {
		identifier, e := strconv.ParseInt(strings.TrimSpace(part), 10, 64)

		if e != nil {
			return response.Fail("invalid memory id %q", part)
		}

		m, f := s.service.GetMemory(identifier)

		if f != nil {
			return response.Fail("memory %d not found", identifier)
		}

		related, g := s.service.ListRelated(identifier)

		if g != nil {
			return s.captureFail(g, "load relations")
		}

		if detail {
			result = append(
				result,
				memoryResponse.NewMemoryWithHistory(*m, related, nil),
			)

			continue
		}

		result = append(
			result,
			memoryResponse.NewSlimMemoryWithHistory(
				*convert.Memory(m),
				related,
				nil,
			),
		)
	}

	return response.Success(notation.MarshalIndent(result))
}
