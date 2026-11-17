package base

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/model/usage_event"
)

func (s *Server) Seed(
	tool string,
	surface string,
	actor string,
) {
	e := usage_event.Stub()
	e.Tool = tool
	e.Surface = surface
	e.Actor = actor
	e.Outcome = constant.OutcomeSuccess
	errors.PanicOnError(s.Store().Create(e))
}
