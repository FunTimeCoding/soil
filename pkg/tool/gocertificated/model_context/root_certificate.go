package model_context

import (
	"context"
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) rootCertificate(
	_ context.Context,
	_ mcp.CallToolRequest,
	_ struct{},
) (*mcp.CallToolResult, error) {
	result, e := s.store.Authority(constant.RootAuthority)

	if e != nil {
		return s.captureFail(e, library.UnexpectedError)
	}

	if result == nil {
		return response.Fail("%s", constant.RootMissing)
	}

	return response.Success("%s", result.Certificate)
}
