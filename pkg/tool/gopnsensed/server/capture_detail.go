package server

import (
	"github.com/funtimecoding/soil/pkg/errors/classify"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/constant"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/server"
)

func (s *Server) captureDetail(e error) *server.ErrorResponse {
	detail := classify.Message(e, constant.RequestFailed)

	if !classify.Reportable(e) {
		return &server.ErrorResponse{Error: detail}
	}

	return s.captureFail(e, detail)
}
