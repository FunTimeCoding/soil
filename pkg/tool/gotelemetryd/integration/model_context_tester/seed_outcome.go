package model_context_tester

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/model/usage_event"
)

func (o *Tester) SeedOutcome(
	tool string,
	actor string,
	outcome string,
	detail string,
) {
	e := usage_event.New()
	e.Tool = tool
	e.Surface = constant.SurfaceCommandLine
	e.Actor = actor
	e.Outcome = outcome
	e.Detail = new(detail)
	errors.PanicOnError(o.Store.Create(e))
}
