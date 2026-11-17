package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/cluster"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/port_forward"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/port_forward/registry"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/unit/service_tester"
	"go.uber.org/goleak"
	"testing"
)

func TestDaemonClosesTwice(t *testing.T) {
	d := service_tester.New(
		t,
		service_tester.WithClusters(map[string]*cluster.Cluster{}),
	).Daemon()
	d.Close()
	d.Close()
}

func TestDaemonAnswersWhetherItHasClusters(t *testing.T) {
	assert.True(t, service_tester.New(t).Daemon().HasClusters())
	assert.False(
		t,
		service_tester.New(
			t,
			service_tester.WithClusters(map[string]*cluster.Cluster{}),
		).Daemon().HasClusters(),
	)
}

func TestDaemonHandsItsGuardsThenItsForwardsToLifecycle(t *testing.T) {
	a := service_tester.New(t)
	workers := service_tester.New(
		t,
		service_tester.WithClusters(
			map[string]*cluster.Cluster{
				"test":     a.Cluster,
				"november": a.Cluster,
			},
		),
	).Daemon().Workers()
	assert.Count(t, 3, workers)
	_, first := workers[0].(*peer.Guard)
	assert.True(t, first)
	_, second := workers[1].(*peer.Guard)
	assert.True(t, second)
	_, forwards := workers[2].(*registry.Registry)
	assert.True(t, forwards)
}

func TestDaemonLeavesNothingRunning(t *testing.T) {
	a := service_tester.New(t)
	s := service_tester.New(
		t,
		service_tester.WithClusters(
			map[string]*cluster.Cluster{
				"test":     a.Cluster,
				"november": a.Cluster,
			},
		),
	)
	before := goleak.IgnoreCurrent()
	d := s.Daemon()
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
