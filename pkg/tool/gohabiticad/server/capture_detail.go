package server

import (
	"github.com/funtimecoding/soil/pkg/errors/classify"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/constant"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/generated/server"
)

func (s *Server) captureDetail(e error) *server.ErrorResponse {
	detail := classify.Message(e, constant.RequestFailed)

	if !classify.Reportable(e) {
		return &server.ErrorResponse{Error: detail}
	}

	return s.captureFail(e, detail)
}
