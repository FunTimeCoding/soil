package runner_tester

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/provision/runner"
	"github.com/funtimecoding/soil/pkg/system/run"
	"path/filepath"
	"testing"
)

func New(t *testing.T) *Tester {
	t.Helper()
	base := t.TempDir()
	remote := filepath.Join(base, "remote")
	clone := filepath.Join(base, "clone")
	c := run.New()
	c.Start("git", "init", "--bare", remote)
	c = run.New()
	c.Start("git", "clone", remote, clone)
	c = run.New()
	c.Directory = clone
	c.Start("git", "commit", "--allow-empty", "-m", "initial")
	c = run.New()
	c.Directory = clone
	c.Start("git", "push", "origin", "main")
	result := &Tester{t: t}
	result.Runner = runner.New(
		runner.Configuration{
			Repository: remote,
			ClonePath:  clone,
			ApplyFunction: func(
				parameters map[string]any,
				triggerSource string,
			) any {
				call := &ApplyCall{
					Parameters:    parameters,
					TriggerSource: triggerSource,
				}
				result.mutex.Lock()
				result.applied = append(result.applied, call)
				result.mutex.Unlock()

				return call
			},
		},
		logger.New(context.Background()),
		memory.New(),
	)
	result.Runner.Start()
	t.Cleanup(func() { result.Runner.Stop() })

	return result
}
