package field_changer

import "github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/field_changer/severity"

func (a *Changer) AddSeverity(
	alert string,
	from string,
	to string,
) *Changer {
	a.severity = append(a.severity, severity.New(alert, from, to))

	return a
}
