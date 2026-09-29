package model_context

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/web/detail_error"
	"github.com/mark3labs/mcp-go/mcp"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (s *Server) captureDetail(e error) (*mcp.CallToolResult, error) {
	if d, okay := errors.AsType[*detail_error.Detail](e); okay {
		return s.captureFail(e, d.Detail)
	}

	if v, okay := errors.AsType[*validation.Detail](e); okay {
		return s.captureFail(e, v.Message)
	}

	if f, okay := errors.AsType[*gitlab.ErrorResponse](e); okay {
		if f.Message != "" {
			return s.captureFail(e, f.Message)
		}
	}

	return s.captureFail(e, constant.UnexpectedError)
}
