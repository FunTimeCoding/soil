package model_context_tester

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/store"
)

func (o *Tester) Seed(
	tool string,
	surface string,
	actor string,
) {
	e := store.NewUsageEvent()
	e.Tool = tool
	e.Surface = surface
	e.Actor = actor
	e.Outcome = constant.OutcomeSuccess
	errors.PanicOnError(o.Store.Create(e))
}
