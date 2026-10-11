package server

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service"
	"time"
)

type Server struct {
	service           *service.Service
	logger            *logger.Logger
	reporter          face.Reporter
	harborPath        string
	sessionExportPath string
	hold              time.Duration
}
