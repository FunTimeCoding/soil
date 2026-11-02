package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/advanced_option"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert_filter"
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
	"github.com/funtimecoding/soil/pkg/prometheus/rule/rule_list"
	"testing"
)

func TestAnEmptyFilterSendsNoQuery(t *testing.T) {
	c, queries := newServingAlertmanager(
		t,
		constant.FixtureSuppressedAndUnratedAlerts,
	)
	_, _, e := c.Alerts(
		advanced_option.New(),
		alert_filter.New(),
		rule_list.New(),
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, *queries)
	assert.String(t, "", (*queries)[0])
}

func TestAnOptionForAllKeepsEveryAlert(t *testing.T) {
	c, _ := newServingAlertmanager(
		t,
		constant.FixtureSuppressedAndUnratedAlerts,
	)
	o := advanced_option.New()
	o.All = true
	v, _, e := c.Alerts(o, alert_filter.New(), rule_list.New())
	assert.FatalOnError(t, e)
	assert.Count(t, 2, v)
}

func TestAnEmptyOptionDropsSuppressedAndUnratedAlerts(t *testing.T) {
	c, _ := newServingAlertmanager(
		t,
		constant.FixtureSuppressedAndUnratedAlerts,
	)
	v, _, e := c.Alerts(
		advanced_option.New(),
		alert_filter.New(),
		rule_list.New(),
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 0, v)
}
