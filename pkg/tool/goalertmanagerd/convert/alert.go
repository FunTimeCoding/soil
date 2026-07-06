package convert

import (
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert"
	"time"
)

func Alert(a *alert.Alert) *SlimAlert {
	var start string

	if a.Start != nil {
		start = a.Start.Format(time.RFC3339)
	}

	labels := make(map[string]string, len(a.Remaining))

	for k, v := range a.Remaining {
		labels[k] = v
	}

	return &SlimAlert{
		Name:        a.Name,
		State:       a.State,
		Severity:    a.Severity,
		Summary:     a.Summary,
		Fingerprint: a.Fingerprint,
		Start:       start,
		Receivers:   a.Receivers,
		Labels:      labels,
		Link:        a.Link,
	}
}
