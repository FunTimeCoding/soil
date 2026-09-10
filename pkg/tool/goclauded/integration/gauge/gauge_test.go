package gauge

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/gauge"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/service_tester"
	"github.com/prometheus/client_golang/prometheus"
	"testing"
)

func value(
	t *testing.T,
	r *prometheus.Registry,
	kind string,
) float64 {
	t.Helper()
	families, e := r.Gather()
	assert.FatalOnError(t, e)

	for _, f := range families {
		if f.GetName() != "goclauded_findings" {
			continue
		}

		for _, m := range f.GetMetric() {
			for _, l := range m.GetLabel() {
				if l.GetName() == "kind" && l.GetValue() == kind {
					return m.GetGauge().GetValue()
				}
			}
		}
	}

	return -1
}

func TestGaugePublishesEveryKind(t *testing.T) {
	s := service_tester.New(t)
	r := prometheus.NewRegistry()
	g := gauge.New(s.Service, logger.New(context.Background()), r)
	g.Poll()

	for _, kind := range constant.FindingKinds {
		assert.Float(t, 0, value(t, r, kind))
	}
}

func TestGaugeCarriesTheCountNotTheFindingTally(t *testing.T) {
	s := service_tester.New(t)

	for i := 0; i < 3; i++ {
		assert.FatalOnError(
			t,
			s.Store.Store.PushQueue(
				"",
				"Nobody",
				constant.QueueTimeout,
				"orphaned",
			),
		)
	}

	r := prometheus.NewRegistry()
	g := gauge.New(s.Service, logger.New(context.Background()), r)
	g.Poll()
	assert.Float(t, 3, value(t, r, constant.UnownedQueue))
}

func TestGaugeResetsWhenFindingResolves(t *testing.T) {
	s := service_tester.New(t)
	assert.FatalOnError(
		t,
		s.Store.Store.PushQueue(
			"",
			"Nobody",
			constant.QueueTimeout,
			"orphaned",
		),
	)
	r := prometheus.NewRegistry()
	g := gauge.New(s.Service, logger.New(context.Background()), r)
	g.Poll()
	assert.Float(t, 1, value(t, r, constant.UnownedQueue))
	assert.FatalOnError(
		t,
		s.Store.Store.DeletePendingQueue("", "Nobody", constant.QueueTimeout),
	)
	g.Poll()
	assert.Float(t, 0, value(t, r, constant.UnownedQueue))
}
