package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/unit/base"
	"testing"
)

func TestClientWalksTheCursorFromTheHead(t *testing.T) {
	s := base.New(t)
	c := s.AlertClient(t)
	status, e := c.GetStatus(t.Context())
	assert.FatalOnError(t, e)
	assert.Integer(t, 2, status.LatestIdentifier)
	limit := 5
	all, f := c.GetAlertsAfter(
		t.Context(),
		&client.GetAlertsAfterParams{Identifier: 0, Limit: &limit},
	)
	assert.FatalOnError(t, f)
	assert.Count(t, 2, all)
	assert.String(t, "HighMemory", all[0].Name)
	assert.True(t, all[0].Status == client.Firing)
	after, g := c.GetAlertsAfter(
		t.Context(),
		&client.GetAlertsAfterParams{
			Identifier: status.LatestIdentifier,
			Limit:      &limit,
		},
	)
	assert.FatalOnError(t, g)
	assert.Count(t, 0, after)
}

func TestClientSeesAResolutionByIdentifier(t *testing.T) {
	s := base.New(t)
	c := s.AlertClient(t)
	s.MockClient.Remove("fp1")
	s.Worker.Poll()
	result, e := c.GetAlertsByIdentifier(
		t.Context(),
		&client.GetAlertsByIdentifierParams{Identifier: []int{1, 2}},
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 2, result)
	assert.True(t, result[0].Status == client.Resolved)
	assert.True(t, result[1].Status == client.Firing)
}
