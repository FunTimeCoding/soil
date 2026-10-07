package model_context_tester

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/model/usage_event"
)

func (o *Tester) Seed(
	tool string,
	surface string,
	actor string,
) {
	e := usage_event.New()
	e.Tool = tool
	e.Surface = surface
	e.Actor = actor
	e.Outcome = constant.OutcomeSuccess
	errors.PanicOnError(o.Store.Create(e))
}
