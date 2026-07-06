package alertmanager

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/advanced_option"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/statistic"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert_filter"
)

func (c *Client) MustAlerts(
	p *advanced_option.Alert,
	f *alert_filter.Filter,
) ([]*alert.Alert, *statistic.Statistic) {
	v, s, e := c.Alerts(p, f)
	errors.PanicOnError(e)

	return v, s
}
