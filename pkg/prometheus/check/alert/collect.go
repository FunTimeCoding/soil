package alert

import (
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/advanced_option"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/statistic"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert_filter"
	"github.com/funtimecoding/soil/pkg/prometheus/check/alert/option"
	"github.com/funtimecoding/soil/pkg/prometheus/rule/rule_list"
)

func collect(
	c *alertmanager.Client,
	o *option.Alert,
	r *rule_list.List,
) ([]*alert.Alert, *statistic.Statistic) {
	d := advanced_option.New()
	d.All = o.All
	d.CriticalOnly = o.Critical
	d.WarningOnly = o.Warning
	d.InformationOnly = o.Information
	d.Suppressed = o.Suppressed

	return c.MustAlerts(d, alert_filter.New(), r)
}
