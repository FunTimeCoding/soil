package alert_filter

import (
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/constant"
	"github.com/funtimecoding/soil/pkg/strings/contains"
)

func (f *Filter) match(a *alert.Alert) bool {
	p := f.option

	if p.All {
		return true
	}

	switch a.Severity {
	case constant.NoneSeverity:
		return false
	}

	if len(p.Receiver) > 0 && !contains.Any(p.Receiver, a.Receivers) {
		return false
	}

	if !p.Suppressed && a.Suppressed() {
		return false
	}

	if p.CriticalOnly && a.Severity != constant.CriticalSeverity {
		return false
	}

	if p.WarningOnly && a.Severity != constant.WarningSeverity {
		return false
	}

	if p.InformationOnly && a.Severity != constant.InformationSeverity {
		return false
	}

	return true
}
