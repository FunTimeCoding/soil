package alert

import (
	"github.com/funtimecoding/soil/pkg/console"
	monitor "github.com/funtimecoding/soil/pkg/monitor/constant"
	"github.com/funtimecoding/soil/pkg/openapi"
	prometheus "github.com/funtimecoding/soil/pkg/prometheus/constant"
	"github.com/prometheus/alertmanager/api/v2/models"
	"slices"
)

func New(
	v *models.GettableAlert,
	host string,
) *Alert {
	var remaining models.LabelSet

	if v.Labels != nil {
		remaining = make(models.LabelSet, len(v.Labels))

		for k, v := range v.Labels {
			remaining[k] = v
		}
	}

	state := *v.Status.State

	if state == "" {
		state = prometheus.None
	} else {
		if !slices.Contains(prometheus.AlertStates, state) {
			console.Format("Unexpected state: %s\n", state)
		}
	}

	var receivers []string

	for _, receiver := range v.Receivers {
		if receiver.Name == nil {
			continue
		}

		n := *receiver.Name

		switch n {
		case "":
			continue
		case "null":
			continue
		}

		receivers = append(receivers, n)
	}

	result := &Alert{
		MonitorIdentifier: monitor.GoAlert.StringIdentifier(*v.Fingerprint),
		Fingerprint:       *v.Fingerprint,
		State:             state,
		Labels:            v.Labels,
		Receivers:         receivers,
		Remaining:         remaining,
		Raw:               v,
		Start:             openapi.ConvertTime(v.StartsAt),
		instance:          StripColon,
	}
	extractKey(&remaining, prometheus.AlertnameLabel, &result.Name)
	extractKey(&remaining, prometheus.SeverityLabel, &result.Severity)
	extractKey(&remaining, prometheus.SummaryLabel, &result.Summary)
	extractKey(&remaining, prometheus.MessageLabel, &result.Message)
	extractKey(&remaining, prometheus.PrometheusLabel, &result.Prometheus)
	result.Link = result.buildLink(host)

	return result
}
