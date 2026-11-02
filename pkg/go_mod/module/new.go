package module

import (
	"github.com/funtimecoding/soil/pkg/go_mod/response"
	monitorConstant "github.com/funtimecoding/soil/pkg/monitor/constant"
	"golang.org/x/mod/module"
	"time"
)

func New(
	r *response.Module,
	now time.Time,
	threshold int,
) *Module {
	latest := r

	if r.Update != nil {
		latest = r.Update
	}

	return &Module{
		MonitorIdentifier: monitorConstant.GoStale.StringIdentifier(r.Path),
		Path:              r.Path,
		Version:           r.Version,
		Latest:            latest.Version,
		Release:           latest.Time,
		Indirect:          r.Indirect,
		Untagged:          module.IsPseudoVersion(latest.Version),
		now:               now,
		threshold:         threshold,
	}
}
