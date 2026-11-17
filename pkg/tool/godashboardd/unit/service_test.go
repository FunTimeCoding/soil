package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/prometheus/query_result"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/worker"
	"github.com/prometheus/common/model"
	"testing"
	"time"
)

func TestClickOnBoardLabelIsCounted(t *testing.T) {
	v := service_tester.New(t).Service()
	assert.FatalOnError(t, v.Click("Grafana"))
	assert.True(t, validation.Is(v.Click("Invented")))
	summaries, e := v.Summaries()
	assert.FatalOnError(t, e)
	assert.Count(t, 1, summaries)
	assert.String(t, "Grafana", summaries[0].Label)
}

func TestRefreshFillsEachPanel(t *testing.T) {
	s := service_tester.New(t, service_tester.WithUsage())
	s.Prometheus.SeedResult(
		"count(up == 1)",
		query_result.New(model.Vector{&model.Sample{Value: 7}}, nil, nil),
	)
	s.Usage.Set(12, 3)
	v := s.Service()
	assert.Count(t, 0, v.Refresh())
	values := v.Values()
	assert.String(t, "7", values["Prometheus"][0])
	assert.String(t, "12", values["Nextcloud"][0])
	assert.String(t, "3", values["Nextcloud"][1])
}

func TestEveryRowLabelGetsValue(t *testing.T) {
	v := service_tester.New(
		t,
		service_tester.WithArgocd(),
		service_tester.WithUsage(),
	).Service()
	assert.Count(t, 0, v.Refresh())
	values := v.Values()

	for _, e := range service_tester.FixtureBoard().Entries() {
		assert.Integer(t, len(e.RowLabels()), len(values[e.Label]))
	}

	assert.Count(t, 4, values["Argo CD"])
}

func TestFailingPanelIsReportedAndShownPending(t *testing.T) {
	s := service_tester.New(t, service_tester.WithUsage())
	s.Usage.Failure = errors.New("token expired")
	v := s.Service()
	owned := worker.NewEvery(
		v.Refresh,
		logger.New(t.Context()),
		s.Reporter,
		time.Hour,
	)
	defer owned.Close()
	owned.Poll()
	assert.Count(t, 1, s.Reporter.Events())
	assert.StringContains(
		t,
		"token expired",
		s.Reporter.Events()[0].Error.Error(),
	)
	assert.String(t, "–", v.Values()["Nextcloud"][0])
}
