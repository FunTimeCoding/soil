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
		if e := d.PostTrigger(c); e != nil {
			r.reportDownstream(e, c)

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
