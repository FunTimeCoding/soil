package unit

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/unit/base"
	"testing"
)

func TestTheMockAnswersWhatTheServerAnswers(t *testing.T) {
	s := base.New(t)

	for i := 3; i <= 30; i++ {
		s.MockClient.Add(
			alert.NewBasic(
				fmt.Sprintf("fp%d", i),
				fmt.Sprintf("Alert%d", i),
				constant.WarningSeverity,
				"",
			),
		)
	}

	s.Worker.Poll()
	s.MockClient.Remove("fp1")
	s.Worker.Poll()
	c := s.AlertClient(t)
	m := mirrorAlerts(t, c)
	defer m.Close()
	served, e := c.GetStatus(t.Context())
	assert.FatalOnError(t, e)
	mocked, f := m.GetStatus(t.Context())
	assert.FatalOnError(t, f)
	assert.Integer(t, 30, served.LatestIdentifier)
	assert.Integer(t, served.LatestIdentifier, mocked.LatestIdentifier)
	assert.Integer(t, served.TotalRecords, mocked.TotalRecords)

	for _, p := range []*client.GetAlertsAfterParams{
		{Identifier: 0},
		{Identifier: 0, Limit: new(1)},
		{Identifier: 0, Limit: new(100)},
		{Identifier: 0, Limit: new(0)},
		{Identifier: 0, Limit: new(-3)},
		{Identifier: 10, Limit: new(25)},
		{Identifier: -5},
		{Identifier: 30},
	} {
		want, g := c.GetAlertsAfter(t.Context(), p)
		assert.FatalOnError(t, g)
		got, h := m.GetAlertsAfter(t.Context(), p)
		assert.FatalOnError(t, h)
		assert.String(t, alertLines(want), alertLines(got))
	}

	var all []int

	for i := 1; i <= 30; i++ {
		all = append(all, i)
	}

	for _, p := range []*client.GetAlertsByIdentifierParams{
		{Identifier: []int{1, 2}},
		{Identifier: []int{0, -1, 2}},
		{Identifier: all},
		{Identifier: []int{999}},
	} {
		want, g := c.GetAlertsByIdentifier(t.Context(), p)
		assert.FatalOnError(t, g)
		got, h := m.GetAlertsByIdentifier(t.Context(), p)
		assert.FatalOnError(t, h)
		assert.String(t, alertLines(want), alertLines(got))
	}
}
