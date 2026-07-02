package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/goprometheusd/convert"
	"github.com/funtimecoding/go-library/pkg/tool/goprometheusd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) series(
	x context.Context,
	_ mcp.CallToolRequest,
	_ argument.Series,
) (*mcp.CallToolResult, error) {
	instance, e := s.service.ResolveInstance(s.activeInstanceName(x))

	if e != nil {
		return response.Fail("%s", e)
	}

	v, e := s.service.Series(instance)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(convert.TSDB(v))
}
