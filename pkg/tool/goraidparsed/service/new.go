package service

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/option"
)

func New(
	o *option.Parser,
	l *logger.Logger,
) *Service {
	return &Service{
		parserPath:   o.ParserPath,
		templatePath: o.TemplatePath,
		outputPath:   o.OutputPath,
		logger:       l,
	}
}
