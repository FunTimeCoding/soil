package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/notation"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/convert"
	memoryResponse "github.com/funtimecoding/soil/pkg/tool/gomemoryd/model_context/response"
	"github.com/mark3labs/mcp-go/mcp"
	"strconv"
	"strings"
)

func (s *Server) getMany(
	x context.Context,
	raw string,
	detail bool,
) (*mcp.CallToolResult, error) {
	var result []any

	for _, part := range strings.Split(raw, stringsConstant.Comma) {
		identifier, e := strconv.ParseInt(strings.TrimSpace(part), 10, 64)

		if e != nil {
			return response.Fail("invalid memory id %q", part)
		}

		m, f := s.service.GetMemory(identifier)

		if f != nil {
			return s.answer(x, f)
		}

		related, g := s.service.ListRelated(identifier)

		if g != nil {
			return s.answerWith(x, g, "load relations")
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
