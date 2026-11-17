package unit

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert"
	prometheusConstant "github.com/funtimecoding/soil/pkg/prometheus/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/unit/client_tester"
	"testing"
)

func TestMockAnswersWhatServerAnswers(t *testing.T) {
	s := base.New(t)

	for i := 3; i <= 30; i++ {
		s.Alertmanager.Add(
			alert.NewBasic(
				fmt.Sprintf("fp%d", i),
				fmt.Sprintf("Alert%d", i),
				prometheusConstant.WarningSeverity,
				"",
			),
		)
	}

	s.Clock.Poll()
	s.Alertmanager.Remove("fp1")
	s.Clock.Poll()
	c := s.AlertClient()
	m := client_tester.Mirror(t, c)
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
		assert.String(
			t,
			client_tester.AlertLines(want),
			client_tester.AlertLines(got),
		)
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
		assert.String(
			t,
			client_tester.AlertLines(want),
			client_tester.AlertLines(got),
		)
	}
}
