package runner

import (
	"github.com/funtimecoding/soil/pkg/errors/classify"
	"github.com/funtimecoding/soil/pkg/provision/constant"
	"github.com/funtimecoding/soil/pkg/provision/types/change"
)

func (r *Runner) reportDownstream(
	e error,
	c *change.Change,
) {
	if !classify.Reportable(e) {
		r.logger.Structured(
			"downstream_refused",
			constant.RunnerError,
			e.Error(),
			constant.DownstreamChanges,
			len(c.Addresses),
		)

		return
	}

	r.reporter.CaptureWithContext(
		e,
		constant.DownstreamContext,
		map[string]any{
			constant.DownstreamChanges: c.Addresses,
			constant.DownstreamHosts:   c.Hosts,
		},
	)
}
