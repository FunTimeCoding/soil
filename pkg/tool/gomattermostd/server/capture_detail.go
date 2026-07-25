package server

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/generated/server"
	"github.com/mattermost/mattermost/server/public/model"
)

func (s *Server) captureDetail(e error) *server.ErrorResponse {
	if f, okay := errors.AsType[*model.AppError](e); okay {
		if f.Message != "" {
			return s.captureFail(e, f.Message)
		}
	}

	return s.captureFail(e, constant.UnexpectedError)
}
