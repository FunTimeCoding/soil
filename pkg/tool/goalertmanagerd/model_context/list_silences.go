package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/goalertmanagerd/convert"
	"github.com/funtimecoding/go-library/pkg/tool/goalertmanagerd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
	"strings"
)

func (s *Server) listSilences(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.ListSilences,
) (*mcp.CallToolResult, error) {
	instance, e := s.service.ResolveInstance(s.activeInstanceName(x))

	if e != nil {
		return response.Fail("%s", e)
	}

	v, e := s.service.Silences(instance, a.Expired)

	if e != nil {
		return s.captureDetail(e)
	}

	result := convert.Silences(v)

	if a.Filter != "" {
		var filtered []*convert.SlimSilence

		for _, s := range result {
			if strings.Contains(s.Rule, a.Filter) ||
				strings.Contains(s.Match, a.Filter) {
				filtered = append(filtered, s)
			}
		}

		result = filtered
	}

	return response.SuccessAny(
		convert.NewSilenceResult(result, int(a.Limit), int(a.Offset)),
	)
}
