package server

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/generated/server"
)

func (s *Server) captureDetail(e error) *server.ErrorResponse {
	if unexpected.Is(e) {
		return s.captureFail(e, e.Error())
	}

	return s.captureFail(e, constant.UnexpectedError)
}
