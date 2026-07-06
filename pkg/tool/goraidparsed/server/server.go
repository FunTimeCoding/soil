package server

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
)

type Server struct {
	parserPath   string
	templatePath string
	outputPath   string
	logger       *logger.Logger
	reporter     face.Reporter
}
