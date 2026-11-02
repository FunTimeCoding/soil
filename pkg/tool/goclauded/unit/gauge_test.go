package unit

import (
	"context"
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/finding"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/gauge"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/service_tester"
	"github.com/prometheus/client_golang/prometheus"
	"testing"
)

func TestGaugePublishesEveryKind(t *testing.T) {
	s := service_tester.New(t)
	r := prometheus.NewRegistry()
	g := gauge.New(
		s.Roster.Findings,
		constant.FindingInterval,
		logger.New(context.Background()),
		r,
		memory.New(),
	)
	defer g.Close()
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
	g := gauge.New(
		s.Roster.Findings,
		constant.FindingInterval,
		logger.New(context.Background()),
		r,
		memory.New(),
	)
	defer g.Close()
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
	g := gauge.New(
		s.Roster.Findings,
		constant.FindingInterval,
		logger.New(context.Background()),
		r,
		memory.New(),
	)
	defer g.Close()
	g.Poll()
	assert.Float(t, 1, value(t, r, constant.UnownedQueue))
	assert.FatalOnError(
		t,
		s.Store.Store.DeletePendingQueue("", "Nobody", constant.QueueTimeout),
	)
	g.Poll()
	assert.Float(t, 0, value(t, r, constant.UnownedQueue))
}

func TestGaugeReportsAFailedPoll(t *testing.T) {
	m := memory.New()
	g := gauge.New(
		func() ([]*finding.Finding, error) {
			return nil, errors.New("findings refused")
		},
		constant.FindingInterval,
		logger.New(context.Background()),
		prometheus.NewRegistry(),
		m,
	)
	defer g.Close()
	g.Poll()
	assert.Count(t, 1, m.Events())
}
