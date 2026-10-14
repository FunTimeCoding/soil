package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/unit/web_service_tester"
	"net/http"
	"testing"
)

func TestAlertsAfterPagesFiveByDefault(t *testing.T) {
	o := web_service_tester.New(t)
	createRecords(t, o.Store(), 10)
	r, e := o.Client.GetAlertsAfterWithResponse(
		context.Background(),
		&client.GetAlertsAfterParams{Identifier: 0},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusOK, r.StatusCode())
	assert.Count(t, 5, *r.JSON200)
	assert.Integer(t, 1, (*r.JSON200)[0].Identifier)
}

func TestAlertsAfterCapsTheLimit(t *testing.T) {
	o := web_service_tester.New(t)
	createRecords(t, o.Store(), 30)
	r, e := o.Client.GetAlertsAfterWithResponse(
		context.Background(),
		&client.GetAlertsAfterParams{Identifier: 0, Limit: new(100)},
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 25, *r.JSON200)
}

func TestAlertsByIdentifierReturnsTheOnesThatExist(t *testing.T) {
	o := web_service_tester.New(t)
	r, e := o.Client.GetAlertsByIdentifierWithResponse(
		context.Background(),
		&client.GetAlertsByIdentifierParams{Identifier: []int{2, 99}},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusOK, r.StatusCode())
	assert.Count(t, 1, *r.JSON200)
	assert.String(t, "fp2", (*r.JSON200)[0].Fingerprint)
}

func TestStatusNamesTheLatestIdentifier(t *testing.T) {
	o := web_service_tester.New(t)
	r, e := o.Client.GetStatusWithResponse(context.Background())
	assert.FatalOnError(t, e)
	assert.Integer(t, 2, r.JSON200.LatestIdentifier)
}
