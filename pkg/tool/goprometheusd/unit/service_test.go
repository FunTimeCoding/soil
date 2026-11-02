package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/prometheus/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goprometheusd/unit/service_tester"
	"github.com/prometheus/client_golang/api/prometheus/v1"
	"testing"
	"time"
)

func TestQueryAnswersFromTheNamedInstance(t *testing.T) {
	s := refusingBravo(t)
	_, e := s.Query(t.Context(), "alfa", "up", time.Now())
	assert.FatalOnError(t, e)
	_, f := s.Query(t.Context(), "bravo", "up", time.Now())
	assert.True(t, validation.Is(f))
}

func TestQueryRangeAnswersFromTheNamedInstance(t *testing.T) {
	s := refusingBravo(t)
	r := v1.Range{Start: time.Now().Add(-time.Hour), End: time.Now()}
	_, e := s.QueryRange(t.Context(), "alfa", "up", r)
	assert.FatalOnError(t, e)
	_, f := s.QueryRange(t.Context(), "bravo", "up", r)
	assert.True(t, validation.Is(f))
}

func TestLabelNamesAnswersFromTheNamedInstance(t *testing.T) {
	s := refusingBravo(t)
	_, e := s.LabelNames(t.Context(), "alfa", nil, time.Now())
	assert.FatalOnError(t, e)
	_, f := s.LabelNames(t.Context(), "bravo", nil, time.Now())
	assert.True(t, validation.Is(f))
}

func TestLabelValuesAnswersFromTheNamedInstance(t *testing.T) {
	s := refusingBravo(t)
	_, e := s.LabelValues(t.Context(), "alfa", "job", nil, time.Now())
	assert.FatalOnError(t, e)
	_, f := s.LabelValues(t.Context(), "bravo", "job", nil, time.Now())
	assert.True(t, validation.Is(f))
}

func TestRulesAnswersFromTheNamedInstance(t *testing.T) {
	s := refusingBravo(t)
	_, e := s.Rules(t.Context(), "alfa")
	assert.FatalOnError(t, e)
	_, f := s.Rules(t.Context(), "bravo")
	assert.True(t, validation.Is(f))
}

func TestSeriesAnswersFromTheNamedInstance(t *testing.T) {
	s := refusingBravo(t)
	_, e := s.Series(t.Context(), "alfa")
	assert.FatalOnError(t, e)
	_, f := s.Series(t.Context(), "bravo")
	assert.True(t, validation.Is(f))
}

func TestAnUnreachableInstanceFailsFastWithoutBeingAskedAgain(t *testing.T) {
	alfa := mock_client.New()
	alfa.Failure = refusedConnection()
	s := service_tester.New(t, map[string]*mock_client.Client{"alfa": alfa})
	_, e := s.Query(t.Context(), "alfa", "up", time.Now())
	assert.True(t, unavailable.Is(e))
	alfa.Failure = nil
	_, f := s.Query(t.Context(), "alfa", "up", time.Now())
	assert.True(t, unavailable.Is(f))
	assert.StringContains(t, "prometheus:alfa unreachable since", f.Error())
}

func TestOneUnreachableInstanceLeavesTheOthersAnswering(t *testing.T) {
	alfa := mock_client.New()
	alfa.Failure = refusedConnection()
	s := service_tester.New(
		t,
		map[string]*mock_client.Client{"alfa": alfa, "bravo": mock_client.New()},
	)
	_, e := s.Query(t.Context(), "alfa", "up", time.Now())
	assert.True(t, unavailable.Is(e))
	_, f := s.Query(t.Context(), "bravo", "up", time.Now())
	assert.FatalOnError(t, f)
}

func TestTheServiceHandsOneGuardPerInstanceToTheLifecycle(t *testing.T) {
	assert.Count(t, 2, refusingBravo(t).Workers())
}

func TestClosingTheServiceClosesEveryInstanceClient(t *testing.T) {
	alfa := mock_client.New()
	bravo := mock_client.New()
	s := service_tester.New(
		t,
		map[string]*mock_client.Client{"alfa": alfa, "bravo": bravo},
	)
	s.Close()
	assert.Integer(t, 1, alfa.Closed)
	assert.Integer(t, 1, bravo.Closed)
}
