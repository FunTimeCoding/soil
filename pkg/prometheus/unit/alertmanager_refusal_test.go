package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/advanced_option"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert_filter"
	"github.com/funtimecoding/soil/pkg/prometheus/rule/rule_list"
	"github.com/funtimecoding/soil/pkg/prometheus/unit/alertmanager_tester"
	"github.com/funtimecoding/soil/pkg/web/upstream_server"
	"net/http"
	"testing"
)

func TestBadAlertmanagerFilterIsCallersMistake(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusBadRequest),
		upstream_server.WithAnswer(`"bad matcher format: severity~"`),
	)
	_, _, e := alertmanager_tester.New(t, s.URL).Alerts(
		advanced_option.New(),
		alert_filter.New(),
		rule_list.New(),
	)
	assert.True(t, validation.Is(e))
	assert.StringContains(t, "bad matcher format", e.Error())
}

func TestUnknownSilenceIsNotFound(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusNotFound),
		upstream_server.WithAnswer(`"silence not found"`),
	)
	e := alertmanager_tester.New(t, s.URL).DeleteSilence(
		"8d1f0c4a-0000-4000-8000-000000000000",
	)
	assert.True(t, not_found.Is(e))
}

func TestAlertmanagerBehindUnavailableProxyIsUnavailable(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusServiceUnavailable),
		upstream_server.WithAnswer(
			"<html>503 Service Temporarily Unavailable</html>",
		),
	)
	_, _, e := alertmanager_tester.New(t, s.URL).Alerts(
		advanced_option.New(),
		alert_filter.New(),
		rule_list.New(),
	)
	assert.True(t, unavailable.Is(e))
}

func TestAlertmanagerFaultKeepsItsText(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusInternalServerError),
		upstream_server.WithAnswer(`"failed to iterate alerts"`),
	)
	_, _, e := alertmanager_tester.New(t, s.URL).Alerts(
		advanced_option.New(),
		alert_filter.New(),
		rule_list.New(),
	)
	assert.Error(t, e)
	assert.False(t, validation.Is(e))
	assert.False(t, unavailable.Is(e))
	assert.StringContains(t, "failed to iterate alerts", e.Error())
}
