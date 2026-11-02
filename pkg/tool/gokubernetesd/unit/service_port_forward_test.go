package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/port_forward"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/port_forward/registry"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/types/request"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/unit/service_tester"
	"testing"
	"time"
)

func TestAPortForwardRefusesAnIncompleteRequest(t *testing.T) {
	v := service_tester.New(t)

	for _, q := range []request.PortForward{
		{Name: "alfa", LocalPort: 8080, TargetPort: 80},
		{ResourceType: "pod", LocalPort: 8080, TargetPort: 80},
		{ResourceType: "pod", Name: "alfa", TargetPort: 80},
		{ResourceType: "pod", Name: "alfa", LocalPort: 8080},
		{ResourceType: "deployment", Name: "alfa", LocalPort: 8080, TargetPort: 80},
	} {
		_, e := v.Service.PortForward(t.Context(), "test", q)
		assert.True(t, validation.Is(e))
	}

	assert.Count(t, 0, v.Service.PortForwards())
}

func TestStoppingAnUnknownPortForwardIsNotFound(t *testing.T) {
	assert.True(
		t,
		not_found.Is(service_tester.New(t).Service.StopPortForward("bravo")),
	)
}

func TestStoppingAPortForwardEndsItAndForgetsIt(t *testing.T) {
	r := registry.New()
	t.Cleanup(r.Close)
	f := port_forward.New(
		"pod-charlie-8080",
		"pod/charlie",
		"default",
		8080,
		80,
	)
	r.Launch(f, func() { <-f.Stopping() })
	assert.Count(t, 1, r.All())
	assert.FatalOnError(t, r.StopOne(f.Identifier))
	assert.True(t, stopped(f))
	assert.Count(t, 0, r.All())
}

func TestClosingTheRegistryStopsEveryPortForward(t *testing.T) {
	r := registry.New()
	f := port_forward.New("pod-delta-8080", "pod/delta", "default", 8080, 80)
	g := port_forward.New("pod-echo-8081", "pod/echo", "default", 8081, 80)
	r.Launch(f, func() { <-f.Stopping() })
	r.Launch(g, func() { <-g.Stopping() })
	r.Close()
	assert.True(t, stopped(f))
	assert.True(t, stopped(g))
	assert.Count(t, 0, r.All())
}

func TestAPortForwardThatEndsOnItsOwnLeavesTheRegistry(t *testing.T) {
	r := registry.New()
	f := port_forward.New("pod-hotel-8080", "pod/hotel", "default", 8080, 80)
	r.Launch(f, func() {})
	r.Close()
	assert.Count(t, 0, r.All())
}

func TestAPortForwardStopsSafelyTwice(t *testing.T) {
	f := port_forward.New(
		"pod-foxtrot-8080",
		"pod/foxtrot",
		"default",
		8080,
		80,
	)
	f.Stop()
	f.Stop()
	assert.True(t, stopped(f))
}

func TestAPortForwardThatFailsBeforeReadyAnswersAtOnce(t *testing.T) {
	failed := make(chan error, 1)
	failed <- errors.New("unable to listen on any of the requested ports")
	done := make(chan error)
	go func() { done <- port_forward.AwaitReady(make(chan struct{}), failed) }()

	select {
	case e := <-done:
		assert.StringContains(t, "unable to listen", e.Error())
	case <-time.After(time.Second):
		t.Fatal("waited on a forward that never became ready")
	}
}

func TestAReadyPortForwardAnswersWithoutError(t *testing.T) {
	ready := make(chan struct{})
	close(ready)
	assert.Nil(t, port_forward.AwaitReady(ready, make(chan error)))
}
