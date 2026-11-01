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
	v := service_tester.New(
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
	assert.FatalOnError(t, v.Run("charlie"))
	assert.String(t, "alfa\nbravo", service_tester.AwaitDone(t, v).Result)
	assert.Strings(t, []string{"claude", "--print", "charlie"}, arguments)
}

func TestAnEmptyIntentIsRefused(t *testing.T) {
	v := service_tester.New(
		t,
		service_tester.Fixed(mock_stream.New(nil, false, nil)),
	)
	assert.True(t, validation.Is(v.Run(" \n")))
	assert.String(t, "idle", string(v.Status().State))
}

func TestASecondRunWhileRunningConflicts(t *testing.T) {
	s := mock_stream.New(nil, true, nil)
	defer s.Close()
	v := service_tester.New(t, service_tester.Fixed(s))
	assert.FatalOnError(t, v.Run("delta"))
	assert.True(t, conflict.Is(v.Run("echo")))
	v.Workers()[0].(face.Drainer).Interrupt()
	service_tester.AwaitDone(t, v)
}

func TestAFailedRunCarriesTheAgentError(t *testing.T) {
	v := service_tester.New(
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
	assert.FatalOnError(t, v.Run("golf"))
	assert.String(
		t,
		"foxtrot\n\nerror: exit status 1\ntoken expired",
		service_tester.AwaitDone(t, v).Result,
	)
}

func TestShutdownWaitsForTheRunThenInterrupts(t *testing.T) {
	s := mock_stream.New(nil, true, nil)
	defer s.Close()
	v := service_tester.New(t, service_tester.Fixed(s))
	d := v.Workers()[0].(face.Drainer)
	assert.FatalOnError(t, v.Run("hotel"))
	v.Close()
	assert.True(t, unavailable.Is(v.Run("india")))

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
	assert.String(t, "done", string(v.Status().State))
}

func TestAnIdleRunnerFinishesOnClose(t *testing.T) {
	v := service_tester.New(
		t,
		service_tester.Fixed(mock_stream.New(nil, false, nil)),
	)
	v.Close()

	select {
	case <-v.Workers()[0].(face.Drainer).Finished():
	default:
		t.Fatalf("idle runner did not finish on close")
	}
}

func TestALaunchThatPanicsEndsTheRun(t *testing.T) {
	r := memory.New()
	v := service_tester.NewWith(
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
	assert.FatalOnError(t, v.Run("kilo"))
	assert.String(
		t,
		"error: agent stopped before it finished",
		service_tester.AwaitDone(t, v).Result,
	)
	deadline := time.Now().Add(5 * time.Second)

	for len(r.Events()) == 0 && time.Now().Before(deadline) {
		time.Sleep(10 * time.Millisecond)
	}

	assert.Count(t, 1, r.Events())
}
