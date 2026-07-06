package server

import (
	"errors"
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
	"github.com/funtimecoding/soil/pkg/web/detail_error"
)

func (s *Server) captureDetail(e error) *server.ErrorResponse {
	if j, okay := errors.AsType[*jira.Error](e); okay {
		if m := jiraMessage(j); m != "" {
			return s.captureFail(e, m)
		}
	}

	if d, okay := errors.AsType[*detail_error.Detail](e); okay {
		return s.captureFail(e, d.Detail)
	}

	return s.captureFail(e, constant.UnexpectedError)
}
