package service

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/face"
)

func New(
	p face.TopStatsParser,
	templatePath string,
	outputPath string,
	l *logger.Logger,
) *Service {
	return &Service{
		parser:       p,
		templatePath: templatePath,
		outputPath:   outputPath,
		logger:       l,
	}
}
