package alert

import (
	"fmt"
	item "github.com/funtimecoding/soil/pkg/monitor/item/constant"
	"github.com/funtimecoding/soil/pkg/openapi"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/constant"
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
		state = constant.None
	} else {
		if !slices.Contains(constant.AlertStates, state) {
			fmt.Printf("Unexpected state: %s\n", state)
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
		MonitorIdentifier: item.GoAlert.StringIdentifier(*v.Fingerprint),
		Fingerprint:       *v.Fingerprint,
		State:             state,
		Labels:            v.Labels,
		Receivers:         receivers,
		Remaining:         remaining,
		Raw:               v,
		Start:             openapi.ConvertTime(v.StartsAt),
		instance:          StripColon,
	}
	extractKey(&remaining, constant.AlertnameLabel, &result.Name)
	extractKey(&remaining, constant.SeverityLabel, &result.Severity)
	extractKey(&remaining, constant.SummaryLabel, &result.Summary)
	extractKey(&remaining, constant.MessageLabel, &result.Message)
	extractKey(&remaining, constant.PrometheusLabel, &result.Prometheus)
	result.Link = result.buildLink(host)

	return result
}
