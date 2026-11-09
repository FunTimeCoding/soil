package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/worker"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"
	"strings"
	"testing"
)

func TestWorkerExportsEachObjectsStatus(t *testing.T) {
	c := service_tester.NewSource(t)
	c.SeedDevice("alfa").Status = "active"
	c.SeedDevice("bravo").Status = "offline"
	c.SeedVirtualMachine("charlie").Status = "offline"
	y := prometheus.NewRegistry()
	w := worker.New(
		service_tester.Service(t, c).Intents,
		constant.IntentInterval,
		logger.New(t.Context()),
		y,
		memory.New(),
	)
	defer w.Close()
	w.Poll()
	assert.FatalOnError(
		t,
		testutil.GatherAndCompare(
			y,
			strings.NewReader(
				`
# HELP netbox_status A device or virtual machine's netbox status, 1 for the status it holds
# TYPE netbox_status gauge
netbox_status{kind="device",name="alfa",status="active"} 1
netbox_status{kind="device",name="bravo",status="offline"} 1
netbox_status{kind="virtual_machine",name="charlie",status="offline"} 1
`,
			),
			constant.StatusMetric,
		),
	)
}

func TestChangedStatusLeavesNoStaleSeries(t *testing.T) {
	c := service_tester.NewSource(t)
	d := c.SeedDevice("alfa")
	d.Status = "offline"
	y := prometheus.NewRegistry()
	w := worker.New(
		service_tester.Service(t, c).Intents,
		constant.IntentInterval,
		logger.New(t.Context()),
		y,
		memory.New(),
	)
	defer w.Close()
	w.Poll()
	d.Status = "active"
	w.Poll()
	assert.FatalOnError(
		t,
		testutil.GatherAndCompare(
			y,
			strings.NewReader(
				`
# HELP netbox_status A device or virtual machine's netbox status, 1 for the status it holds
# TYPE netbox_status gauge
netbox_status{kind="device",name="alfa",status="active"} 1
`,
			),
			constant.StatusMetric,
		),
	)
}

func TestUnavailableNetboxIsOnlyLogged(t *testing.T) {
	c := service_tester.NewSource(t)
	c.Failure = errors.New("netbox answered nonsense")
	r := memory.New()
	w := worker.New(
		service_tester.Service(t, c).Intents,
		constant.IntentInterval,
		logger.New(t.Context()),
		prometheus.NewRegistry(),
		r,
	)
	defer w.Close()
	w.Poll()
	assert.Count(t, 1, r.Events())
	c.Failure = unavailable.Format("netbox restarting")
	w.Poll()
	assert.Count(t, 1, r.Events())
}

func TestQuietFailureIsOnlyLogged(t *testing.T) {
	c := service_tester.NewSource(t)
	c.Failure = not_found.Format("no devices endpoint")
	r := memory.New()
	w := worker.New(
		service_tester.Service(t, c).Intents,
		constant.IntentInterval,
		logger.New(t.Context()),
		prometheus.NewRegistry(),
		r,
	)
	defer w.Close()
	w.Poll()
	assert.Count(t, 0, r.Events())
}
