package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/unit/base"
	"net/http"
	"testing"
)

func TestWebService(t *testing.T) {
	c := base.New(t).GeneratedClient()
	x := context.Background()
	status, e := c.GetStatusWithResponse(x)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusOK, status.StatusCode())
	assert.NotNil(t, status.JSON200)
	assert.Integer(t, 2, status.JSON200.TotalRecords)
	assert.True(t, status.JSON200.LastPoll != nil)
	alerts, e := c.GetAlertsWithResponse(
		x,
		&client.GetAlertsParams{Name: "HighMemory"},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusOK, alerts.StatusCode())
	assert.NotNil(t, alerts.JSON200)
	assert.Count(t, 1, *alerts.JSON200)
	assert.String(t, "fp1", (*alerts.JSON200)[0].Fingerprint)
	assert.String(t, "firing", string((*alerts.JSON200)[0].Status))
	recent, e := c.GetRecentAlertsWithResponse(
		x,
		&client.GetRecentAlertsParams{},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusOK, recent.StatusCode())
	assert.Count(t, 2, *recent.JSON200)
	top, e := c.GetTopAlertsWithResponse(x, &client.GetTopAlertsParams{})
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusOK, top.StatusCode())
	assert.Count(t, 2, *top.JSON200)
	assert.Integer(t, 1, (*top.JSON200)[0].Count)
}

func TestWebServiceResolve(t *testing.T) {
	s := base.New(t)
	c := s.GeneratedClient()
	x := context.Background()
	s.Alertmanager.Remove("fp1")
	s.Clock.Poll()
	alerts, e := c.GetAlertsWithResponse(
		x,
		&client.GetAlertsParams{Name: "HighMemory"},
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, *alerts.JSON200)
	assert.String(t, "resolved", string((*alerts.JSON200)[0].Status))
	assert.True(t, (*alerts.JSON200)[0].End != nil)
	status, e := c.GetStatusWithResponse(x)
	assert.FatalOnError(t, e)
	assert.Integer(t, 2, status.JSON200.TotalRecords)
}
