package service_tester

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/option"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/service"
)

func (o *Tester) Service() *service.Service {
	o.t.Helper()

	if o.service != nil {
		return o.service
	}

	p := option.New()
	o.service = service.New(
		o.Store,
		logger.New(o.t.Context()),
		p.Predicate,
		p.Retention,
		p.Clock,
		o.System.Run,
		o.System.Stream,
	)

	return o.service
}
