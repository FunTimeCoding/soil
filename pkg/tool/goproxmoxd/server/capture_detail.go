package server

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/classify"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/server"
	"github.com/luthermonson/go-proxmox"
)

func (s *Server) captureDetail(e error) *server.ErrorResponse {
	if errors.Is(e, proxmox.ErrNotAuthorized) ||
		errors.Is(e, proxmox.ErrTimeout) ||
		classify.Known(e) {
		return s.captureFail(e, e.Error())
	}

	return s.captureFail(e, constant.UnexpectedError)
}
