package edge

import "github.com/funtimecoding/soil/pkg/reacher/constant"

func (e *Edge) Context() map[string]any {
	if e.Down {
		return map[string]any{
			constant.HostKey:   e.Host,
			constant.StateKey:  constant.DownState,
			constant.ReasonKey: e.Reason,
		}
	}

	return map[string]any{
		constant.HostKey:     e.Host,
		constant.StateKey:    constant.UpState,
		constant.DurationKey: e.Duration.String(),
	}
}
