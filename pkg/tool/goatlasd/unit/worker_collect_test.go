package unit

import (
	"context"
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/face"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/mock_collector"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/unit/collector_tester"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/unit/gazetteer_tester"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/unit/store_tester"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/unit/worker_tester"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"
	"strings"
	"testing"
)

func TestCollectKeepsSourcesAfterAFailingOne(t *testing.T) {
	s := store_tester.NewStore(t)
	broken := mock_collector.New(constant.SourceKubernetes)
	broken.Fail(errors.New("device list status: 500"))
	healthy := mock_collector.New(constant.SourceNetbox)
	healthy.Add(
		collector_tester.NewPlacement(
			constant.SourceNetbox,
			constant.FixtureService,
		),
	)
	worker_tester.NewWorker(
		t,
		s,
		[]face.Collector{broken, healthy},
		prometheus.NewRegistry(),
	).Collect(context.Background(), gazetteer_tester.NewGazetteer())
	v, e := s.Placements()
	assert.FatalOnError(t, e)
	assert.Count(t, 1, v)
	assert.String(t, "netbox", v[0].Source)
	assert.String(t, "relay", v[0].Name)
}

func TestCollectKeepsTheFailingSourcePlacements(t *testing.T) {
	s := store_tester.NewStore(t)
	healthy := mock_collector.New(constant.SourceProcess)
	healthy.Add(
		collector_tester.NewPlacement(
			constant.SourceProcess,
			constant.FixtureService,
		),
	)
	w := worker_tester.NewWorker(
		t,
		s,
		[]face.Collector{healthy},
		prometheus.NewRegistry(),
	)
	w.Collect(context.Background(), gazetteer_tester.NewGazetteer())
	healthy.Fail(errors.New("process list status: 500"))
	w.Collect(context.Background(), gazetteer_tester.NewGazetteer())
	v, e := s.Placements()
	assert.FatalOnError(t, e)
	assert.Count(t, 1, v)
	assert.String(t, "goprocessd", v[0].Source)
}

func TestCollectCountsFailuresPerSource(t *testing.T) {
	broken := mock_collector.New(constant.SourceKubernetes)
	broken.Fail(errors.New("device list status: 500"))
	healthy := mock_collector.New(constant.SourceNetbox)
	y := prometheus.NewRegistry()
	worker_tester.NewWorker(
		t,
		store_tester.NewStore(t),
		[]face.Collector{broken, healthy},
		y,
	).Collect(context.Background(), gazetteer_tester.NewGazetteer())
	assert.FatalOnError(
		t,
		testutil.GatherAndCompare(
			y,
			strings.NewReader(
				`
# HELP atlas_collect_failures_total Cycles a source failed to collect or save
# TYPE atlas_collect_failures_total counter
atlas_collect_failures_total{source="kubernetes"} 1
atlas_collect_failures_total{source="netbox"} 0
`,
			),
			"atlas_collect_failures_total",
		),
	)
}
