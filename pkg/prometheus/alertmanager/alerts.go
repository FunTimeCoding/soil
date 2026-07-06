package alertmanager

import (
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/advanced_option"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/statistic"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert_filter"
	raw "github.com/prometheus/alertmanager/api/v2/client/alert"
)

func (c *Client) Alerts(
	p *advanced_option.Alert,
	f *alert_filter.Filter,
) ([]*alert.Alert, *statistic.Statistic, error) {
	var a *raw.GetAlertsParams

	if f != nil {
		a = raw.NewGetAlertsParams()
		a.Filter = f.Filter
		a.Active = f.Active
		a.Silenced = f.Silenced
		a.Inhibited = f.Inhibited
		a.Unprocessed = f.Unprocessed

		if f.Receiver != "" {
			a.Receiver = &f.Receiver
		}
	}

	result, e := c.client.Alert.GetAlerts(a)

	if e != nil {
		return nil, nil, e
	}

	proc, g := c.processor(p)

	if g != nil {
		return nil, nil, g
	}

	v, s := proc.Process(alert.NewSlice(result.GetPayload(), c.host))

	return v, s, nil
}
