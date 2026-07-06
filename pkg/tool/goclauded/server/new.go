package server

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service"
)

func New(
	s *service.Service,
	l *logger.Logger,
	r face.Reporter,
	harborPath string,
	sessionExportPath string,
) *Server {
	return &Server{
		service:           s,
		logger:            l,
		reporter:          r,
		harborPath:        harborPath,
		sessionExportPath: sessionExportPath,
	}
}
