package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert_filter"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/mock_client"
	prometheusMock "github.com/funtimecoding/soil/pkg/prometheus/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/unit/service_tester"
	"testing"
	"time"
)

func TestAlertsAnswerFromTheNamedInstance(t *testing.T) {
	s := refusingBravo(t)
	_, _, e := s.Alerts(t.Context(), "alfa", alert_filter.New())
	assert.FatalOnError(t, e)
	_, _, f := s.Alerts(t.Context(), "bravo", alert_filter.New())
	assert.True(t, validation.Is(f))
}

func TestSilencesAnswerFromTheNamedInstance(t *testing.T) {
	s := refusingBravo(t)
	_, e := s.Silences(t.Context(), "alfa", false)
	assert.FatalOnError(t, e)
	_, f := s.Silences(t.Context(), "bravo", false)
	assert.True(t, validation.Is(f))
}

func TestCreateSilenceAnswersFromTheNamedInstance(t *testing.T) {
	s := refusingBravo(t)
	_, e := s.CreateSilence(t.Context(), "alfa", "Watchdog", "", time.Hour)
	assert.FatalOnError(t, e)
	_, f := s.CreateSilence(t.Context(), "bravo", "Watchdog", "", time.Hour)
	assert.True(t, validation.Is(f))
}

func TestDeleteSilenceAnswersFromTheNamedInstance(t *testing.T) {
	s := refusingBravo(t)
	assert.FatalOnError(t, s.DeleteSilence(t.Context(), "alfa", "silence-1"))
	assert.True(
		t,
		validation.Is(s.DeleteSilence(t.Context(), "bravo", "silence-1")),
	)
}

func TestStatusAnswersFromTheNamedInstance(t *testing.T) {
	s := refusingBravo(t)
	_, e := s.Status(t.Context(), "alfa")
	assert.FatalOnError(t, e)
	_, f := s.Status(t.Context(), "bravo")
	assert.True(t, validation.Is(f))
}

func TestAlertsReadTheirRulesFromThePrometheusPeer(t *testing.T) {
	p := prometheusMock.New()
	p.Failure = validation.New("bad rules request")
	s := service_tester.New(
		t,
		map[string]*mock_client.Client{"alfa": mock_client.New()},
		map[string]*prometheusMock.Client{"alfa": p},
	)
	_, _, e := s.Alerts(t.Context(), "alfa", alert_filter.New())
	assert.True(t, validation.Is(e))
}

func TestCreateSilenceReadsItsRulesFromThePrometheusPeer(t *testing.T) {
	p := prometheusMock.New()
	p.Failure = validation.New("bad rules request")
	s := service_tester.New(
		t,
		map[string]*mock_client.Client{"alfa": mock_client.New()},
		map[string]*prometheusMock.Client{"alfa": p},
	)
	_, e := s.CreateSilence(t.Context(), "alfa", "Watchdog", "", time.Hour)
	assert.True(t, validation.Is(e))
}

func TestAPrometheusOutageLeavesAlertmanagerAnswering(t *testing.T) {
	p := prometheusMock.New()
	p.Failure = refusedConnection()
	s := service_tester.New(
		t,
		map[string]*mock_client.Client{"alfa": mock_client.New()},
		map[string]*prometheusMock.Client{"alfa": p},
	)
	_, _, e := s.Alerts(t.Context(), "alfa", alert_filter.New())
	assert.True(t, unavailable.Is(e))
	assert.StringContains(t, "prometheus:alfa unreachable since", e.Error())
	_, f := s.Silences(t.Context(), "alfa", false)
	assert.FatalOnError(t, f)
}

func TestAnUnreachableAlertmanagerFailsFastWithoutBeingAskedAgain(t *testing.T) {
	a := mock_client.New()
	a.Fail(refusedConnection())
	s := service_tester.New(
		t,
		map[string]*mock_client.Client{"alfa": a},
		map[string]*prometheusMock.Client{"alfa": prometheusMock.New()},
	)
	_, e := s.Silences(t.Context(), "alfa", false)
	assert.True(t, unavailable.Is(e))
	a.Fail(nil)
	_, f := s.Silences(t.Context(), "alfa", false)
	assert.True(t, unavailable.Is(f))
	assert.StringContains(t, "alertmanager:alfa unreachable since", f.Error())
}

func TestTheServiceHandsTwoGuardsPerInstanceToTheLifecycle(t *testing.T) {
	assert.Count(t, 4, healthy(t, "alfa", "bravo").Workers())
}

func TestClosingTheServiceClosesBothClientsOfEveryInstance(t *testing.T) {
	a := mock_client.New()
	p := prometheusMock.New()
	s := service_tester.New(
		t,
		map[string]*mock_client.Client{"alfa": a},
		map[string]*prometheusMock.Client{"alfa": p},
	)
	s.Close()
	assert.Integer(t, 1, a.Closed)
	assert.Integer(t, 1, p.Closed)
}
