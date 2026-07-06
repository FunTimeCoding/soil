//go:build local

package web_service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/integration_test/web_service_tester"
	"net/http"
	"testing"
)

func TestWebService(t *testing.T) {
	o := web_service_tester.New(t)
	defer o.Close()
	c := o.Client
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
	assert.String(
		t,
		string(client.Firing),
		string((*alerts.JSON200)[0].Status),
	)
	recent, e := c.GetRecentAlertsWithResponse(
		x,
		&client.GetRecentAlertsParams{},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusOK, recent.StatusCode())
	assert.Count(t, 2, *recent.JSON200)
	top, e := c.GetTopAlertsWithResponse(
		x,
		&client.GetTopAlertsParams{},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusOK, top.StatusCode())
	assert.Count(t, 2, *top.JSON200)
	assert.Integer(t, 1, (*top.JSON200)[0].Count)
}

func TestWebServiceResolve(t *testing.T) {
	o := web_service_tester.New(t)
	defer o.Close()
	c := o.Client
	x := context.Background()
	o.MockClient.Remove("fp1")
	o.Worker.Poll()
	alerts, e := c.GetAlertsWithResponse(
		x,
		&client.GetAlertsParams{Name: "HighMemory"},
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, *alerts.JSON200)
	assert.String(
		t,
		string(client.Resolved),
		string((*alerts.JSON200)[0].Status),
	)
	assert.True(t, (*alerts.JSON200)[0].End != nil)
	status, e := c.GetStatusWithResponse(x)
	assert.FatalOnError(t, e)
	assert.Integer(t, 2, status.JSON200.TotalRecords)
}
