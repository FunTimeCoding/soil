package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/unit/base"
	"net/http"
	"testing"
)

func TestAlertsAfterPagesFiveByDefault(t *testing.T) {
	s := base.New(t)
	s.Seed(10)
	r, e := s.GeneratedClient().GetAlertsAfterWithResponse(
		context.Background(),
		&client.GetAlertsAfterParams{Identifier: 0},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusOK, r.StatusCode())
	assert.Count(t, 5, *r.JSON200)
	assert.Integer(t, 1, (*r.JSON200)[0].Identifier)
}

func TestAlertsAfterCapsLimit(t *testing.T) {
	s := base.New(t)
	s.Seed(30)
	r, e := s.GeneratedClient().GetAlertsAfterWithResponse(
		context.Background(),
		&client.GetAlertsAfterParams{Identifier: 0, Limit: new(100)},
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 25, *r.JSON200)
}

func TestAlertsByIdentifierReturnsOnesThatExist(t *testing.T) {
	r, e := base.New(t).GeneratedClient().GetAlertsByIdentifierWithResponse(
		context.Background(),
		&client.GetAlertsByIdentifierParams{Identifier: []int{2, 99}},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusOK, r.StatusCode())
	assert.Count(t, 1, *r.JSON200)
	assert.String(t, "fp2", (*r.JSON200)[0].Fingerprint)
}

func TestStatusNamesLatestIdentifier(t *testing.T) {
	r, e := base.New(t).GeneratedClient().GetStatusWithResponse(
		context.Background(),
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 2, r.JSON200.LatestIdentifier)
}
