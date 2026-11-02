package mock_client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/advanced_option"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/statistic"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert_filter"
	"github.com/funtimecoding/soil/pkg/prometheus/rule/rule_list"
)

func (c *Client) MustAlerts(
	p *advanced_option.Alert,
	f *alert_filter.Filter,
	r *rule_list.List,
) ([]*alert.Alert, *statistic.Statistic) {
	v, s, e := c.Alerts(p, f, r)
	errors.PanicOnError(e)

	return v, s
}
