package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) revocationList(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.Authority,
) (*mcp.CallToolResult, error) {
	result, e := s.service.RevocationList(a.Authority)

	if errors.Is(e, constant.ErrorNotFound) {
		return response.Fail("%s", constant.AuthorityMissing)
	}

	if e != nil {
		return s.captureFail(e, constant.RevocationListFail)
	}

	return response.Success("%s", string(result))
}
