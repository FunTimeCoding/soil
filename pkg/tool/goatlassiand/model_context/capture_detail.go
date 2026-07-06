package model_context

import (
	"errors"
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/web/detail_error"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) captureDetail(e error) (*mcp.CallToolResult, error) {
	if j, okay := errors.AsType[*jira.Error](e); okay {
		if len(j.ErrorMessages) > 0 {
			return s.captureFail(e, j.ErrorMessages[0])
		}
	}

	if d, okay := errors.AsType[*detail_error.Detail](e); okay {
		return s.captureFail(e, d.Detail)
	}

	return s.captureFail(e, constant.UnexpectedError)
}
