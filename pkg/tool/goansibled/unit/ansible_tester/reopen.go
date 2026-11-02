package ansible_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/option"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/service"
)

func (o *Tester) Reopen() *service.Service {
	o.t.Helper()
	p := option.New()
	p.LitePath = o.litePath
	result := service.New(p, logger.New(o.t.Context()), memory.New())
	o.t.Cleanup(result.Close)

	return result
}
