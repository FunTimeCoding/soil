package server

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
)

func New(
	parserPath string,
	templatePath string,
	outputPath string,
	l *logger.Logger,
	r face.Reporter,
) *Server {
	return &Server{
		parserPath:   parserPath,
		templatePath: templatePath,
		outputPath:   outputPath,
		logger:       l,
		reporter:     r,
	}
}
