package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/advanced_option"
	"net/http"
	"testing"
)

func TestABadAlertmanagerFilterIsTheCallersMistake(t *testing.T) {
	c := newRefusingAlertmanager(
		t,
		http.StatusBadRequest,
		`"bad matcher format: severity~"`,
	)
	_, _, e := c.Alerts(advanced_option.New(), nil)
	assert.True(t, validation.Is(e))
	assert.StringContains(t, "bad matcher format", e.Error())
}

func TestAnUnknownSilenceIsNotFound(t *testing.T) {
	c := newRefusingAlertmanager(t, http.StatusNotFound, `"silence not found"`)
	e := c.DeleteSilence("8d1f0c4a-0000-4000-8000-000000000000")
	assert.True(t, not_found.Is(e))
}

func TestAnAlertmanagerBehindAnUnavailableProxyIsUnavailable(t *testing.T) {
	c := newRefusingAlertmanager(
		t,
		http.StatusServiceUnavailable,
		"<html>503 Service Temporarily Unavailable</html>",
	)
	_, _, e := c.Alerts(advanced_option.New(), nil)
	assert.True(t, unavailable.Is(e))
}

func TestAnAlertmanagerFaultKeepsItsText(t *testing.T) {
	c := newRefusingAlertmanager(
		t,
		http.StatusInternalServerError,
		`"failed to iterate alerts"`,
	)
	_, _, e := c.Alerts(advanced_option.New(), nil)
	assert.Error(t, e)
	assert.False(t, validation.Is(e))
	assert.False(t, unavailable.Is(e))
	assert.StringContains(t, "failed to iterate alerts", e.Error())
}
