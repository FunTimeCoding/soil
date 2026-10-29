package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/port_forward"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/request"
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
	v := service_tester.New(t)
	f := port_forward.New(
		"pod-charlie-8080",
		"pod/charlie",
		"default",
		8080,
		80,
	)
	v.Service.StorePortForward(f)
	assert.FatalOnError(t, v.Service.StopPortForward(f.Identifier))
	assert.True(t, stopped(f))
	assert.Count(t, 0, v.Service.PortForwards())
}

func TestClosingTheServiceStopsEveryPortForward(t *testing.T) {
	v := service_tester.New(t)
	f := port_forward.New("pod-delta-8080", "pod/delta", "default", 8080, 80)
	g := port_forward.New("pod-echo-8081", "pod/echo", "default", 8081, 80)
	v.Service.StorePortForward(f)
	v.Service.StorePortForward(g)
	v.Service.Close()
	assert.True(t, stopped(f))
	assert.True(t, stopped(g))
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
