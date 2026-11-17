package base

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/model/usage_event"
)

func (s *Server) SeedOutcome(
	tool string,
	actor string,
	outcome string,
	detail string,
) {
	e := usage_event.Stub()
	e.Tool = tool
	e.Surface = constant.SurfaceCommandLine
	e.Actor = actor
	e.Outcome = outcome
	e.Detail = new(detail)
	errors.PanicOnError(s.Store().Create(e))
}
