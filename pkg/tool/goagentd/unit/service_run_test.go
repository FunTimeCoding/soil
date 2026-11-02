package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/system/run/stream/mock_stream"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/unit/service_tester"
	"io"
	"testing"
	"time"
)

func TestARunReturnsTheAgentOutput(t *testing.T) {
	var arguments []string
	o := service_tester.New(
		t,
		func(
			_ string,
			_ io.Writer,
			s ...string,
		) face.Stream {
			arguments = s

			return mock_stream.New([]string{"alfa", "bravo"}, false, nil)
		},
	)
	assert.FatalOnError(t, o.Service.Run("charlie"))
	assert.String(
		t,
		"alfa\nbravo",
		service_tester.AwaitDone(t, o.Service).Result,
	)
	assert.Strings(t, []string{"claude", "--print", "charlie"}, arguments)
}

func TestAnEmptyIntentIsRefused(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.Fixed(mock_stream.New(nil, false, nil)),
	)
	assert.True(t, validation.Is(o.Service.Run(" \n")))
	assert.String(t, "idle", string(o.Service.Status().State))
}

func TestASecondRunWhileRunningConflicts(t *testing.T) {
	s := mock_stream.New(nil, true, nil)
	defer s.Close()
	o := service_tester.New(t, service_tester.Fixed(s))
	assert.FatalOnError(t, o.Service.Run("delta"))
	assert.True(t, conflict.Is(o.Service.Run("echo")))
	o.Runner.Interrupt()
	service_tester.AwaitDone(t, o.Service)
}

func TestAFailedRunCarriesTheAgentError(t *testing.T) {
	o := service_tester.New(
		t,
		func(
			_ string,
			stderr io.Writer,
			_ ...string,
		) face.Stream {
			_, e := io.WriteString(stderr, "token expired")
			assert.FatalOnError(t, e)

			return mock_stream.New(
				[]string{"foxtrot"},
				false,
				errors.New("exit status 1"),
			)
		},
	)
	assert.FatalOnError(t, o.Service.Run("golf"))
	assert.String(
		t,
		"foxtrot\n\nerror: exit status 1\ntoken expired",
		service_tester.AwaitDone(t, o.Service).Result,
	)
}

func TestShutdownWaitsForTheRunThenInterrupts(t *testing.T) {
	s := mock_stream.New(nil, true, nil)
	defer s.Close()
	o := service_tester.New(t, service_tester.Fixed(s))
	d := o.Runner
	assert.FatalOnError(t, o.Service.Run("hotel"))
	o.Runner.Close()
	assert.True(t, unavailable.Is(o.Service.Run("india")))

	select {
	case <-d.Finished():
		t.Fatalf("finished while the agent still runs")
	case <-time.After(50 * time.Millisecond):
	}

	d.Interrupt()

	select {
	case <-d.Finished():
	case <-time.After(5 * time.Second):
		t.Fatalf("interrupt did not finish the run")
	}

	assert.Integer(t, 1, s.Killed())
	assert.String(t, "done", string(o.Service.Status().State))
}

func TestAnIdleRunnerFinishesOnClose(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.Fixed(mock_stream.New(nil, false, nil)),
	)
	o.Runner.Close()

	select {
	case <-o.Runner.Finished():
	default:
		t.Fatalf("idle runner did not finish on close")
	}
}

func TestALaunchThatPanicsEndsTheRun(t *testing.T) {
	r := memory.New()
	o := service_tester.NewWith(
		t,
		func(
			string,
			io.Writer,
			...string,
		) face.Stream {
			panic("juliett")
		},
		r,
	)
	assert.FatalOnError(t, o.Service.Run("kilo"))
	assert.String(
		t,
		"error: agent stopped before it finished",
		service_tester.AwaitDone(t, o.Service).Result,
	)
	deadline := time.Now().Add(5 * time.Second)

	for len(r.Events()) == 0 && time.Now().Before(deadline) {
		time.Sleep(10 * time.Millisecond)
	}

	assert.Count(t, 1, r.Events())
}
