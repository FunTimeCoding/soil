package service

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational/lite/connection"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/option"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"
)

func New(
	o *option.Memory,
	l *logger.Logger,
) *Service {
	d := connection.New(l, o.LitePath)

	return &Service{
		logger:        l,
		database:      d,
		store:         store.New(d),
		query:         o.Query(),
		hiddenTag:     o.HiddenTag,
		referenceRoot: o.ReferenceRoot,
	}
}
