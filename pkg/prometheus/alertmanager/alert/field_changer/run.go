package field_changer

import "github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert"

func (a *Changer) Run(v []*alert.Alert) []*alert.Alert {
	for _, e := range v {
		e.Name = a.Name(e.Name)
	}

	for _, e := range v {
		e.Severity = a.Severity(e.Name, e.Severity)
	}

	return v
}
