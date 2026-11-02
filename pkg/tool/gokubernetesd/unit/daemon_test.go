package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/cluster"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/port_forward"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/port_forward/registry"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/unit/service_tester"
	"go.uber.org/goleak"
	"testing"
)

func TestTheDaemonClosesTwice(t *testing.T) {
	d := service_tester.NewDaemon(
		t,
		service_tester.NewOption(t, map[string]*cluster.Cluster{}),
		memory.New(),
	)
	d.Close()
	d.Close()
}

func TestTheDaemonAnswersWhetherItHasClusters(t *testing.T) {
	a := service_tester.New(t)
	assert.True(
		t,
		service_tester.NewDaemon(
			t,
			service_tester.NewOption(
				t,
				map[string]*cluster.Cluster{"test": a.Cluster},
			),
			memory.New(),
		).HasClusters(),
	)
	assert.False(
		t,
		service_tester.NewDaemon(
			t,
			service_tester.NewOption(t, map[string]*cluster.Cluster{}),
			memory.New(),
		).HasClusters(),
	)
}

func TestTheDaemonHandsItsGuardsThenItsForwardsToTheLifecycle(t *testing.T) {
	a := service_tester.New(t)
	workers := service_tester.NewDaemon(
		t,
		service_tester.NewOption(
			t,
			map[string]*cluster.Cluster{
				"test":     a.Cluster,
				"november": a.Cluster,
			},
		),
		memory.New(),
	).Workers()
	assert.Count(t, 3, workers)
	_, first := workers[0].(*peer.Guard)
	assert.True(t, first)
	_, second := workers[1].(*peer.Guard)
	assert.True(t, second)
	_, forwards := workers[2].(*registry.Registry)
	assert.True(t, forwards)
}

func TestTheDaemonLeavesNothingRunning(t *testing.T) {
	a := service_tester.New(t)
	before := goleak.IgnoreCurrent()
	d := service_tester.NewDaemon(
		t,
		service_tester.NewOption(
			t,
			map[string]*cluster.Cluster{
				"test":     a.Cluster,
				"november": a.Cluster,
			},
		),
		memory.New(),
	)
	workers := d.Workers()

	for _, w := range workers {
		w.Start()
	}

	f := port_forward.New("pod-golf-8080", "pod/golf", "default", 8080, 80)
	workers[len(workers)-1].(*registry.Registry).Launch(
		f,
		func() { <-f.Stopping() },
	)

	for _, w := range workers {
		w.Close()
	}

	d.Close()
	goleak.VerifyNone(t, before)
}
