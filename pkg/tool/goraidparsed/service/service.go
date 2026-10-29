package service

import "github.com/funtimecoding/soil/pkg/log/logger"

type Service struct {
	parserPath   string
	templatePath string
	outputPath   string
	logger       *logger.Logger
}
