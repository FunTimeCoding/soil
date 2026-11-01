package server

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service"
	"time"
)

func New(
	s *service.Service,
	l *logger.Logger,
	r face.Reporter,
	hold time.Duration,
) *Server {
	return &Server{service: s, logger: l, reporter: r, hold: hold}
}
