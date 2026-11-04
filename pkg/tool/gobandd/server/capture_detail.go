package server

import (
	"context"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/generated/server"
)

func (s *Server) captureDetail(x context.Context, e error) *server.ErrorResponse {
	if unexpected.Is(e) {
		return s.captureFail(x, e, e.Error())
	}

	return s.captureFail(x, e, soil.UnexpectedError)
}
