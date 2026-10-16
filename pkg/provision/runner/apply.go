package runner

import "github.com/funtimecoding/soil/pkg/provision/constant"

func (r *Runner) apply(
	parameters map[string]any,
	triggerSource string,
) any {
	value := r.applyFunction(parameters, triggerSource)

	if r.changeFunction == nil || len(r.downstream) == 0 {
		return value
	}

	c := r.changeFunction(value)

	if c == nil || len(c.Addresses) == 0 {
		return value
	}

	for _, d := range r.downstream {
		if e := d.Trigger(c); e != nil {
			r.reporter.CaptureWithContext(
				e,
				constant.DownstreamContext,
				map[string]any{
					constant.DownstreamChanges: c.Addresses,
					constant.DownstreamHosts:   c.Hosts,
				},
			)

			continue
		}

		r.logger.Structured(
			"downstream_trigger",
			constant.DownstreamChanges,
			len(c.Addresses),
			constant.DownstreamHosts,
			len(c.Hosts),
		)
	}

	return value
}
