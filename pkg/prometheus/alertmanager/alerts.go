package alertmanager

import (
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/advanced_option"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/statistic"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert_filter"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/refusal"
	"github.com/funtimecoding/soil/pkg/prometheus/rule/rule_list"
	raw "github.com/prometheus/alertmanager/api/v2/client/alert"
)

func (c *Client) Alerts(
	p *advanced_option.Alert,
	f *alert_filter.Filter,
	r *rule_list.List,
) ([]*alert.Alert, *statistic.Statistic, error) {
	a := raw.NewGetAlertsParams()
	a.Filter = f.Filter
	a.Active = f.Active
	a.Silenced = f.Silenced
	a.Inhibited = f.Inhibited
	a.Unprocessed = f.Unprocessed

	if f.Receiver != "" {
		a.Receiver = &f.Receiver
	}

	result, e := c.client.Alert.GetAlerts(a)

	if e != nil {
		return nil, nil, refusal.Classify(e)
	}

	v, s := c.processor(p, r).Process(
		alert.NewSlice(result.GetPayload(), c.link),
	)

	return v, s, nil
}
