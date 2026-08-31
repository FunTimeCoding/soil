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

	changes := r.changeFunction(value)

	if len(changes) == 0 {
		return value
	}

	for _, d := range r.downstream {
		if e := d.Trigger(changes); e != nil {
			r.logger.Structured("downstream_trigger_error", "cause", e.Error())
		} else {
			r.logger.Structured(
				"downstream_trigger",
				constant.DownstreamChanges,
				len(changes),
			)
		}
	}

	return value
}
