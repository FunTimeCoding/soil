package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_configured"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	provision "github.com/funtimecoding/soil/pkg/provision/client"
	"github.com/funtimecoding/soil/pkg/provision/constant"
	"github.com/funtimecoding/soil/pkg/provision/generated/client"
	"github.com/funtimecoding/soil/pkg/provision/generated/server"
	"github.com/funtimecoding/soil/pkg/provision/mock_runner"
	"github.com/funtimecoding/soil/pkg/provision/unit/runner_tester"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
)

func TestClientHostCommandReturnsRun(t *testing.T) {
	n := mock_runner.New()
	s := runner_tester.Mount(t, n, memory.New())
	c := provision.New(
		locator.New(webConstant.Localhost).Port(s.Port).Insecure(),
		generativeConstant.ModelContextTestToken,
	)
	defer c.Close()
	result, e := c.PostHostCommand(t.Context(), "alfa", "bravo", "charlie")
	assert.FatalOnError(t, e)
	assert.Any(
		t,
		client.HostCommandRun{Identifier: 1, Status: "success"},
		result,
	)
	assert.Any(
		t,
		[]server.HostCommandRequest{
			{Command: "alfa", Host: "bravo", Source: "charlie"},
		},
		n.HostCommands(),
	)
}

func TestClientHostCommandUnknownNotFound(t *testing.T) {
	r := memory.New()
	s := runner_tester.Mount(
		t,
		mock_runner.New().WithFail(not_found.New("host command", "alfa")),
		r,
	)
	c := provision.New(
		locator.New(webConstant.Localhost).Port(s.Port).Insecure(),
		generativeConstant.ModelContextTestToken,
	)
	defer c.Close()
	_, e := c.PostHostCommand(t.Context(), "alfa", "bravo", "charlie")
	assert.True(t, not_found.Is(e))
	assert.Count(t, 0, r.Events())
}

func TestClientHostCommandOutsideGroupInvalid(t *testing.T) {
	r := memory.New()
	s := runner_tester.Mount(
		t,
		mock_runner.New().WithFail(validation.New("host %s outside", "bravo")),
		r,
	)
	c := provision.New(
		locator.New(webConstant.Localhost).Port(s.Port).Insecure(),
		generativeConstant.ModelContextTestToken,
	)
	defer c.Close()
	_, e := c.PostHostCommand(t.Context(), "alfa", "bravo", "charlie")
	assert.True(t, validation.Is(e))
	assert.Count(t, 0, r.Events())
}

func TestClientHostCommandStoppedRunnerUnavailable(t *testing.T) {
	r := memory.New()
	s := runner_tester.Mount(
		t,
		mock_runner.New().WithFail(unavailable.Format(constant.RunnerStopped)),
		r,
	)
	c := provision.New(
		locator.New(webConstant.Localhost).Port(s.Port).Insecure(),
		generativeConstant.ModelContextTestToken,
	)
	defer c.Close()
	_, e := c.PostHostCommand(t.Context(), "alfa", "bravo", "charlie")
	assert.True(t, unavailable.Is(e))
	assert.Count(t, 0, r.Events())
}

func TestClientHostCommandFaultCaptured(t *testing.T) {
	r := memory.New()
	s := runner_tester.Mount(
		t,
		mock_runner.New().WithFail(errors.New("delta")),
		r,
	)
	c := provision.New(
		locator.New(webConstant.Localhost).Port(s.Port).Insecure(),
		generativeConstant.ModelContextTestToken,
	)
	defer c.Close()
	_, e := c.PostHostCommand(t.Context(), "alfa", "bravo", "charlie")
	assert.Error(t, e)
	assert.Count(t, 1, r.Events())
}

func TestRunnerWithoutHostCommandNotConfigured(t *testing.T) {
	_, e := runner_tester.New(t).Runner.HostCommand("alfa", "bravo", "charlie")
	assert.True(t, not_configured.Is(e))
}
