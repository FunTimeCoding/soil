package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/advanced_option"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert_filter"
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
	"github.com/funtimecoding/soil/pkg/prometheus/rule/rule_list"
	"github.com/funtimecoding/soil/pkg/prometheus/unit/alertmanager_tester"
	"github.com/funtimecoding/soil/pkg/web/upstream_server"
	"testing"
)

func TestEmptyFilterSendsNoQuery(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithAnswer(constant.FixtureSuppressedAndUnratedAlerts),
		upstream_server.WithObject(),
	)
	_, _, e := alertmanager_tester.New(t, s.URL).Alerts(
		advanced_option.New(),
		alert_filter.New(),
		rule_list.New(),
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 1, s.Count())
	assert.String(t, "", s.Last().URL.RawQuery)
}

func TestOptionForAllKeepsEveryAlert(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithAnswer(constant.FixtureSuppressedAndUnratedAlerts),
		upstream_server.WithObject(),
	)
	o := advanced_option.New()
	o.All = true
	v, _, e := alertmanager_tester.New(t, s.URL).Alerts(
		o,
		alert_filter.New(),
		rule_list.New(),
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 2, v)
}

func TestEmptyOptionDropsSuppressedAndUnratedAlerts(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithAnswer(constant.FixtureSuppressedAndUnratedAlerts),
		upstream_server.WithObject(),
	)
	v, _, e := alertmanager_tester.New(t, s.URL).Alerts(
		advanced_option.New(),
		alert_filter.New(),
		rule_list.New(),
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 0, v)
}
