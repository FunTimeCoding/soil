package unit

import (
	"context"
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/worker"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"
	"net"
	"os"
	"strings"
	"syscall"
	"testing"
)

func TestPollCountsFailuresPerSource(t *testing.T) {
	v := service_tester.New(t)
	v.SeedInventory()
	v.Kubernetes.Failure = errors.New("pod list refused")
	y := prometheus.NewRegistry()
	r := memory.New()
	w := worker.New(
		v.Service,
		constant.FixtureInterval,
		logger.New(context.Background()),
		y,
		r,
	)
	defer w.Close()
	w.Poll()
	assert.FatalOnError(
		t,
		testutil.GatherAndCompare(
			y,
			strings.NewReader(
				`
# HELP atlas_collect_failures_total Cycles a source failed to collect or save
# TYPE atlas_collect_failures_total counter
atlas_collect_failures_total{source="gopnsensed"} 0
atlas_collect_failures_total{source="goprocessd"} 0
atlas_collect_failures_total{source="kubernetes"} 1
atlas_collect_failures_total{source="netbox"} 0
`,
			),
			"atlas_collect_failures_total",
		),
	)
	assert.Count(t, 1, r.Events())
}

func TestPollReportsNothingForAnAwayPeer(t *testing.T) {
	v := service_tester.New(t)
	v.SeedInventory()
	v.Kubernetes.Failure = &net.OpError{
		Op:  "dial",
		Net: "tcp",
		Err: os.NewSyscallError("connect", syscall.ECONNREFUSED),
	}
	r := memory.New()
	w := worker.New(
		v.Service,
		constant.FixtureInterval,
		logger.New(context.Background()),
		prometheus.NewRegistry(),
		r,
	)
	defer w.Close()
	w.Poll()
	assert.Count(t, 0, r.Events())
}
