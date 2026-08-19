package model_context

import (
	"context"
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/convert"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) getAuthority(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.Name,
) (*mcp.CallToolResult, error) {
	result, e := s.store.Authority(a.Name)

	if e != nil {
		return s.captureFail(e, library.UnexpectedError)
	}

	if result == nil {
		return response.Fail("%s", constant.AuthorityMissing)
	}

	return response.SuccessAny(convert.Authority(result))
}
