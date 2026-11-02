package daemon

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/option"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/service"
)

func New(
	o *option.Parser,
	l *logger.Logger,
) *Daemon {
	return &Daemon{
		service: service.New(
			o.TopStats(o.ParserPath),
			o.TemplatePath,
			o.OutputPath,
			l,
		),
	}
}
