package service

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/face"
)

type Service struct {
	parser       face.TopStatsParser
	templatePath string
	outputPath   string
	logger       *logger.Logger
}
