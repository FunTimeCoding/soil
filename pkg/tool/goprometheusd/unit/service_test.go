package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/prometheus/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goprometheusd/unit/service_tester"
	"testing"
	"time"
)

func TestQueryAnswersFromNamedInstance(t *testing.T) {
	s := refusingBravo(t)
	_, e := s.Query(t.Context(), "alfa", "up")
	assert.FatalOnError(t, e)
	_, f := s.Query(t.Context(), "bravo", "up")
	assert.True(t, validation.Is(f))
}

func TestQueryRangeAnswersFromNamedInstance(t *testing.T) {
	s := refusingBravo(t)
	_, e := s.QueryRange(t.Context(), "alfa", "up", "1h", "", time.Minute)
	assert.FatalOnError(t, e)
	_, f := s.QueryRange(t.Context(), "bravo", "up", "1h", "", time.Minute)
	assert.True(t, validation.Is(f))
}

func TestLabelNamesAnswersFromNamedInstance(t *testing.T) {
	s := refusingBravo(t)
	_, e := s.LabelNames(t.Context(), "alfa", nil, time.Hour)
	assert.FatalOnError(t, e)
	_, f := s.LabelNames(t.Context(), "bravo", nil, time.Hour)
	assert.True(t, validation.Is(f))
}

func TestLabelValuesAnswersFromNamedInstance(t *testing.T) {
	s := refusingBravo(t)
	_, e := s.LabelValues(t.Context(), "alfa", "job", nil, time.Hour)
	assert.FatalOnError(t, e)
	_, f := s.LabelValues(t.Context(), "bravo", "job", nil, time.Hour)
	assert.True(t, validation.Is(f))
}

func TestRulesAnswersFromNamedInstance(t *testing.T) {
	s := refusingBravo(t)
	_, e := s.Rules(t.Context(), "alfa")
	assert.FatalOnError(t, e)
	_, f := s.Rules(t.Context(), "bravo")
	assert.True(t, validation.Is(f))
}

func TestSeriesAnswersFromNamedInstance(t *testing.T) {
	s := refusingBravo(t)
	_, e := s.Series(t.Context(), "alfa")
	assert.FatalOnError(t, e)
	_, f := s.Series(t.Context(), "bravo")
	assert.True(t, validation.Is(f))
}

func TestUnreachableInstanceFailsFastWithoutBeingAskedAgain(t *testing.T) {
	alfa := mock_client.New()
	alfa.Failure = refusedConnection()
	s := service_tester.New(t, map[string]*mock_client.Client{"alfa": alfa})
	_, e := s.Query(t.Context(), "alfa", "up")
	assert.True(t, unavailable.Is(e))
	alfa.Failure = nil
	_, f := s.Query(t.Context(), "alfa", "up")
	assert.True(t, unavailable.Is(f))
	assert.StringContains(t, "prometheus:alfa unreachable since", f.Error())
}

func TestOneUnreachableInstanceLeavesOthersAnswering(t *testing.T) {
	alfa := mock_client.New()
	alfa.Failure = refusedConnection()
	s := service_tester.New(
		t,
		map[string]*mock_client.Client{"alfa": alfa, "bravo": mock_client.New()},
	)
	_, e := s.Query(t.Context(), "alfa", "up")
	assert.True(t, unavailable.Is(e))
	_, f := s.Query(t.Context(), "bravo", "up")
	assert.FatalOnError(t, f)
}
