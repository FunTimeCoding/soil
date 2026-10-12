package model_context

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/common"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) serviceFail(
	e error,
	message string,
) (*mcp.CallToolResult, error) {
	if not_found.Is(e) || validation.Is(e) {
		return response.Fail("%s", e)
	}

	if m, okay := common.ExtractMessage(e); okay {
		return s.captureFail(e, m)
	}

	return s.captureFail(e, message)
}
