package server

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/common"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) captureDetail(e error) *server.ErrorResponse {
	if m, okay := common.ExtractMessage(e); okay {
		return s.captureFail(e, m)
	}

	return s.captureFail(e, constant.UnexpectedError)
}
