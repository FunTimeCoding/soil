package alertmanager

import (
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/refusal"
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
	"github.com/go-openapi/strfmt"
	"github.com/prometheus/alertmanager/api/v2/client/alert"
	"github.com/prometheus/alertmanager/api/v2/models"
	"time"
)

func (c *Client) Create(
	name string,
	instance string,
	summary string,
	description string,
	expression string,
) error {
	start := time.Now()
	end := start.Add(3 * time.Minute)
	p := alert.NewPostAlertsParams()
	p.Alerts = append(
		p.Alerts,
		&models.PostableAlert{
			Alert: models.Alert{
				Labels: models.LabelSet{
					constant.AlertnameLabel: name,
					constant.SeverityLabel:  constant.CriticalSeverity,
					constant.InstanceLabel:  instance,
				},
				GeneratorURL: strfmt.URI(c.prometheus.GraphLink(expression)),
			},
			Annotations: map[string]string{
				constant.SummaryLabel:     summary,
				constant.DescriptionLabel: description,
			},
			StartsAt: strfmt.DateTime(start),
			EndsAt:   strfmt.DateTime(end),
		},
	)
	result, e := c.client.Alert.PostAlerts(p)

	if e != nil {
		return refusal.Classify(e)
	}

	if !result.IsSuccess() {
		return unexpected.Format("unexpected response: %+v", result)
	}

	return nil
}
