package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/kubernetes/argocd/mock_client"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/prometheus/query_result"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/service"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/worker"
	"github.com/prometheus/common/model"
	"testing"
	"time"
)

func TestTheServiceClosesWhatItBuiltOnce(t *testing.T) {
	p := newPrometheus(t)
	u := newUsage(t)
	v := service.New(
		service_tester.NewOption(t, service_tester.FixtureBoard(), p, nil, u),
		logger.New(t.Context()),
		memory.New(),
	)
	v.Close()
	v.Close()
	assert.Integer(t, 1, p.Closed)
	assert.Integer(t, 1, u.Closed)
	_, e := v.Summaries()
	assert.Error(t, e)
}

func TestAClickOnABoardLabelIsCounted(t *testing.T) {
	v := service_tester.New(
		t,
		service_tester.NewOption(
			t,
			service_tester.FixtureBoard(),
			newPrometheus(t),
			nil,
			nil,
		),
		memory.New(),
	)
	assert.FatalOnError(t, v.Click("Grafana"))
	assert.True(t, validation.Is(v.Click("Invented")))
	summaries, e := v.Summaries()
	assert.FatalOnError(t, e)
	assert.Count(t, 1, summaries)
	assert.String(t, "Grafana", summaries[0].Label)
}

func TestARefreshFillsEachPanel(t *testing.T) {
	p := newPrometheus(t)
	p.SeedResult(
		"count(up == 1)",
		query_result.New(model.Vector{&model.Sample{Value: 7}}, nil, nil),
	)
	u := newUsage(t)
	u.Set(12, 3)
	v := service_tester.New(
		t,
		service_tester.NewOption(t, service_tester.FixtureBoard(), p, nil, u),
		memory.New(),
	)
	assert.Count(t, 0, v.Refresh())
	values := v.Values()
	assert.String(t, "7", values["Prometheus"][0])
	assert.String(t, "12", values["Nextcloud"][0])
	assert.String(t, "3", values["Nextcloud"][1])
}

func TestEveryRowLabelGetsAValue(t *testing.T) {
	v := service_tester.New(
		t,
		service_tester.NewOption(
			t,
			service_tester.FixtureBoard(),
			newPrometheus(t),
			mock_client.New(),
			newUsage(t),
		),
		memory.New(),
	)
	assert.Count(t, 0, v.Refresh())
	values := v.Values()

	for _, e := range service_tester.FixtureBoard().Entries() {
		assert.Integer(t, len(e.RowLabels()), len(values[e.Label]))
	}

	assert.Count(t, 4, values["Argo CD"])
}

func TestAFailingPanelIsReportedAndShownPending(t *testing.T) {
	p := newPrometheus(t)
	u := newUsage(t)
	u.Failure = errors.New("token expired")
	r := memory.New()
	l := logger.New(t.Context())
	v := service_tester.New(
		t,
		service_tester.NewOption(t, service_tester.FixtureBoard(), p, nil, u),
		r,
	)
	worker.New(v, time.Hour, l, r).Poll()
	assert.Count(t, 1, r.Events())
	assert.StringContains(t, "token expired", r.Events()[0].Error.Error())
	assert.String(t, "–", v.Values()["Nextcloud"][0])
}
