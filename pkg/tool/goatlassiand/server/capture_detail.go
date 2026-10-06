package server

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/errors/classify"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
	"github.com/funtimecoding/soil/pkg/web/detail_error"
)

func (s *Server) captureDetail(e error) *server.ErrorResponse {
	if d, okay := errors.AsType[*detail_error.Detail](e); okay {
		return s.captureFail(e, d.Detail)
	}

	detail := classify.Message(e, constant.RequestFailed)

	if !classify.Reportable(e) {
		return &server.ErrorResponse{Error: detail}
	}

	return s.captureFail(e, detail)
}
