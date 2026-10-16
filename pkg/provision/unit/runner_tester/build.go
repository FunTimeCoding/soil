package runner_tester

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	library "github.com/funtimecoding/soil/pkg/face"
	gitConstant "github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/provision/constant"
	"github.com/funtimecoding/soil/pkg/provision/face"
	"github.com/funtimecoding/soil/pkg/provision/runner"
	"github.com/funtimecoding/soil/pkg/provision/types/apply_call"
	"github.com/funtimecoding/soil/pkg/provision/types/change"
	"github.com/funtimecoding/soil/pkg/provision/types/host_event"
	"github.com/funtimecoding/soil/pkg/provision/types/runner_option"
	"github.com/funtimecoding/soil/pkg/system/run"
	"path/filepath"
	"testing"
)

func build(
	t *testing.T,
	hold func(),
	registry library.ProcessRegistry,
	downstream face.Downstream,
) *Tester {
	t.Helper()
	base := t.TempDir()
	remote := filepath.Join(base, "remote.git")
	clone := filepath.Join(base, "clone")
	c := run.New()
	c.Start("git", "init", "--bare", "--initial-branch=main", remote)
	c = run.New()
	c.Start("git", "clone", remote, clone)
	c = run.New()
	c.Directory = clone
	c.Start(
		"git",
		"-c",
		"user.name=runner-tester",
		"-c",
		"user.email=runner-tester@localhost",
		"commit",
		"--allow-empty",
		gitConstant.MessageArgument,
		"initial",
	)
	c = run.New()
	c.Directory = clone
	c.Start("git", "push", "origin", constant.RunnerBranch)
	result := &Tester{
		t:         t,
		ClonePath: clone,
		remote:    remote,
		Reporter:  memory.New(),
	}
	o := runner_option.Option{
		Repository: remote,
		ClonePath:  clone,
		ToolPath:   ".",
		ApplyFunction: func(
			parameters map[string]any,
			triggerSource string,
		) any {
			call := apply_call.New(parameters, triggerSource)
			result.mutex.Lock()
			result.applied = append(result.applied, call)
			result.mutex.Unlock()
			hold()

			return call
		},
		Registry: registry,
	}

	if downstream != nil {
		o.Downstream = []face.Downstream{downstream}
		o.ChangeFunction = func(any) *change.Change {
			return change.New(
				[]string{"alfa"},
				[]*host_event.Event{
					host_event.New("bravo", constant.HostUpdate),
				},
			)
		}
	}

	result.Runner = runner.New(
		o,
		logger.New(context.Background()),
		result.Reporter,
	)
	result.Runner.Start()
	t.Cleanup(func() { result.Runner.Close() })

	return result
}
