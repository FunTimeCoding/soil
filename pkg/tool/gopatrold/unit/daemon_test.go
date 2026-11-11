package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/metric"
	"github.com/funtimecoding/soil/pkg/patrol/face"
	"github.com/funtimecoding/soil/pkg/patrol/mock_client"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gopatrold/constant"
	"github.com/funtimecoding/soil/pkg/tool/gopatrold/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gopatrold/option"
	"github.com/funtimecoding/soil/pkg/tool/gopatrold/unit/daemon_tester"
	"github.com/funtimecoding/soil/pkg/tool/gopatrold/worker"
	"github.com/prometheus/client_golang/prometheus/testutil"
	"go.uber.org/goleak"
	"strings"
	"testing"
)

func TestRoundGaugesHealthyHost(t *testing.T) {
	m := metric.New()
	d := daemon_tester.New(
		t,
		map[string]*mock_client.Client{
			"alfa": fixtureClient(
				"mdstat_healthy.txt",
				"smartctl_nvme_healthy.json",
			),
		},
		m,
		memory.New(),
	)
	poll(t, d)
	assert.FatalOnError(
		t,
		testutil.GatherAndCompare(
			m.Registry(),
			strings.NewReader(
				`
# HELP patrol_host_reached 1 when the last round read the host
# TYPE patrol_host_reached gauge
patrol_host_reached{host="alfa"} 1
# HELP patrol_array_members_active Members the array is running on
# TYPE patrol_array_members_active gauge
patrol_array_members_active{array="md1",host="alfa"} 2
# HELP patrol_array_members_failed Members the kernel marked failed
# TYPE patrol_array_members_failed gauge
patrol_array_members_failed{array="md1",host="alfa"} 0
# HELP patrol_drive_health_passed 1 when the drive's own health assessment passes
# TYPE patrol_drive_health_passed gauge
patrol_drive_health_passed{drive="/dev/nvme1",host="alfa",model="Example NVMe 2TB"} 1
# HELP patrol_drive_media_errors NVMe unrecovered data integrity errors
# TYPE patrol_drive_media_errors gauge
patrol_drive_media_errors{drive="/dev/nvme1",host="alfa",model="Example NVMe 2TB"} 0
`,
			),
			constant.ReachedMetric,
			constant.ActiveMetric,
			constant.FailedMetric,
			constant.PassedMetric,
			constant.MediaMetric,
		),
	)
}

func TestRoundGaugesDegradedArrayAndFailingDrive(t *testing.T) {
	m := metric.New()
	d := daemon_tester.New(
		t,
		map[string]*mock_client.Client{
			"alfa": fixtureClient(
				"mdstat_degraded_derived.txt",
				"smartctl_nvme_failing_derived.json",
			),
		},
		m,
		memory.New(),
	)
	poll(t, d)
	assert.FatalOnError(
		t,
		testutil.GatherAndCompare(
			m.Registry(),
			strings.NewReader(
				`
# HELP patrol_array_members_active Members the array is running on
# TYPE patrol_array_members_active gauge
patrol_array_members_active{array="md1",host="alfa"} 1
# HELP patrol_array_rebuilding 1 while the array recovers, resyncs or reshapes
# TYPE patrol_array_rebuilding gauge
patrol_array_rebuilding{array="md1",host="alfa"} 1
# HELP patrol_drive_health_passed 1 when the drive's own health assessment passes
# TYPE patrol_drive_health_passed gauge
patrol_drive_health_passed{drive="/dev/nvme1",host="alfa",model="Example NVMe 2TB"} 0
# HELP patrol_drive_critical_warning NVMe critical warning flags, 0 is none
# TYPE patrol_drive_critical_warning gauge
patrol_drive_critical_warning{drive="/dev/nvme1",host="alfa",model="Example NVMe 2TB"} 4
`,
			),
			constant.ActiveMetric,
			constant.RebuildingMetric,
			constant.PassedMetric,
			constant.WarningMetric,
		),
	)
}

func TestUnreachedHostKeepsOtherHostsGauged(t *testing.T) {
	down := mock_client.New()
	down.Failure = errors.New("sudo: a password is required")
	m := metric.New()
	d := daemon_tester.New(
		t,
		map[string]*mock_client.Client{
			"alfa":  down,
			"bravo": fixtureClient("mdstat_healthy.txt"),
		},
		m,
		memory.New(),
	)
	poll(t, d)
	assert.FatalOnError(
		t,
		testutil.GatherAndCompare(
			m.Registry(),
			strings.NewReader(
				`
# HELP patrol_host_reached 1 when the last round read the host
# TYPE patrol_host_reached gauge
patrol_host_reached{host="alfa"} 0
patrol_host_reached{host="bravo"} 1
# HELP patrol_array_members_active Members the array is running on
# TYPE patrol_array_members_active gauge
patrol_array_members_active{array="md1",host="bravo"} 2
`,
			),
			constant.ReachedMetric,
			constant.ActiveMetric,
		),
	)
}

func TestRoundDropsDriveNoLongerScanned(t *testing.T) {
	c := fixtureClient(
		"mdstat_healthy.txt",
		"smartctl_nvme_healthy.json",
		"smartctl_nvme_system.json",
	)
	m := metric.New()
	d := daemon_tester.New(
		t,
		map[string]*mock_client.Client{"alfa": c},
		m,
		memory.New(),
	)
	poll(t, d)
	c.RemoveDrive("/dev/nvme1")
	poll(t, d)
	assert.FatalOnError(
		t,
		testutil.GatherAndCompare(
			m.Registry(),
			strings.NewReader(
				`
# HELP patrol_drive_health_passed 1 when the drive's own health assessment passes
# TYPE patrol_drive_health_passed gauge
patrol_drive_health_passed{drive="/dev/nvme0",host="alfa",model="Example NVMe 128GB"} 1
`,
			),
			constant.PassedMetric,
		),
	)
}

func TestDaemonHandsOutGuardsBeforeClock(t *testing.T) {
	w := daemon_tester.New(
		t,
		map[string]*mock_client.Client{
			"alfa":  mock_client.New(),
			"bravo": mock_client.New(),
		},
		metric.New(),
		memory.New(),
	).Workers()
	assert.Count(t, 3, w)

	for _, g := range w[:2] {
		_, okay := g.(*peer.Guard)
		assert.True(t, okay)
	}

	_, okay := w[2].(*worker.Worker)
	assert.True(t, okay)
}

func TestDaemonLeavesNothingRunning(t *testing.T) {
	before := goleak.IgnoreCurrent()
	c := mock_client.New()
	o := option.New()
	o.Hosts = []string{"alfa"}
	o.Source = func(string) face.PatrolClient { return c }
	d := daemon.New(o, logger.New(t.Context()), memory.New(), metric.New())
	workers := d.Workers()

	for _, w := range workers {
		w.Start()
	}

	for _, w := range workers {
		w.Close()
	}

	d.Close()
	goleak.VerifyNone(t, before)
}
