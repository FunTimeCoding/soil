package service

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/option"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/store"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/constant"
)

func New(
	o *option.Raid,
	l *logger.Logger,
	r face.Reporter,
) *Service {
	d := relational.Open(l, o.PostgresLocator, o.LitePath)
	p := o.Parser()

	return &Service{
		database:      d,
		store:         store.New(d),
		parser:        p,
		parserGuard:   peer.New(constant.Identity.Name(), "", p, l, r),
		authorization: authorizationClient(o).WithReporter(r),
		logger:        l,
		logCachePath:  o.LogCachePath,
		elitePath:     o.ElitePath,
		outputPath:    o.OutputPath,
	}
}
