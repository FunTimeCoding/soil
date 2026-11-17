package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert_filter"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"testing"
	"time"
)

func TestAlertsAnswerFromNamedInstance(t *testing.T) {
	s := service_tester.NewRefusingBravo(t)
	_, _, e := s.Alerts(t.Context(), "alfa", alert_filter.New())
	assert.FatalOnError(t, e)
	_, _, f := s.Alerts(t.Context(), "bravo", alert_filter.New())
	assert.True(t, validation.Is(f))
}

func TestSilencesAnswerFromNamedInstance(t *testing.T) {
	s := service_tester.NewRefusingBravo(t)
	_, e := s.Silences(t.Context(), "alfa", false)
	assert.FatalOnError(t, e)
	_, f := s.Silences(t.Context(), "bravo", false)
	assert.True(t, validation.Is(f))
}

func TestCreateSilenceAnswersFromNamedInstance(t *testing.T) {
	s := service_tester.NewRefusingBravo(t)
	_, e := s.CreateSilence(t.Context(), "alfa", "Watchdog", "", time.Hour)
	assert.FatalOnError(t, e)
	_, f := s.CreateSilence(t.Context(), "bravo", "Watchdog", "", time.Hour)
	assert.True(t, validation.Is(f))
}

func TestDeleteSilenceAnswersFromNamedInstance(t *testing.T) {
	s := service_tester.NewRefusingBravo(t)
	assert.FatalOnError(t, s.DeleteSilence(t.Context(), "alfa", "silence-1"))
	assert.True(
		t,
		validation.Is(s.DeleteSilence(t.Context(), "bravo", "silence-1")),
	)
}

func TestStatusAnswersFromNamedInstance(t *testing.T) {
	s := service_tester.NewRefusingBravo(t)
	_, e := s.Status(t.Context(), "alfa")
	assert.FatalOnError(t, e)
	_, f := s.Status(t.Context(), "bravo")
	assert.True(t, validation.Is(f))
}

func TestAlertsReadTheirRulesFromPrometheusPeer(t *testing.T) {
	o := service_tester.New(t)
	o.Alfa().Prometheus.Failure = validation.New("bad rules request")
	s := o.Service()
	_, _, e := s.Alerts(t.Context(), "alfa", alert_filter.New())
	assert.True(t, validation.Is(e))
}

func TestCreateSilenceReadsItsRulesFromPrometheusPeer(t *testing.T) {
	o := service_tester.New(t)
	o.Alfa().Prometheus.Failure = validation.New("bad rules request")
	s := o.Service()
	_, e := s.CreateSilence(t.Context(), "alfa", "Watchdog", "", time.Hour)
	assert.True(t, validation.Is(e))
}

func TestPrometheusOutageLeavesAlertmanagerAnswering(t *testing.T) {
	o := service_tester.New(t)
	o.Alfa().Prometheus.Failure = outage_tester.RefusedDial()
	s := o.Service()
	_, _, e := s.Alerts(t.Context(), "alfa", alert_filter.New())
	assert.True(t, unavailable.Is(e))
	assert.StringContains(t, "prometheus:alfa unreachable since", e.Error())
	_, f := s.Silences(t.Context(), "alfa", false)
	assert.FatalOnError(t, f)
}

func TestUnreachableAlertmanagerFailsFastWithoutBeingAskedAgain(t *testing.T) {
	o := service_tester.New(t)
	o.Alfa().Alertmanager.Fail(outage_tester.RefusedDial())
	s := o.Service()
	_, e := s.Silences(t.Context(), "alfa", false)
	assert.True(t, unavailable.Is(e))
	o.Alfa().Alertmanager.Fail(nil)
	_, f := s.Silences(t.Context(), "alfa", false)
	assert.True(t, unavailable.Is(f))
	assert.StringContains(t, "alertmanager:alfa unreachable since", f.Error())
}
