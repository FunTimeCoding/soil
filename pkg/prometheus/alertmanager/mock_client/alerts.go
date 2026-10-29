package mock_client

import (
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/advanced_option"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/statistic"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert_filter"
)

func (c *Client) Alerts(
	_ *advanced_option.Alert,
	_ *alert_filter.Filter,
) ([]*alert.Alert, *statistic.Statistic, error) {
	if c.failure != nil {
		return nil, nil, c.failure
	}

	return c.alerts, statistic.New(), nil
}
