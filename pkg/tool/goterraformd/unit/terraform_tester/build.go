package terraform_tester

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	gitConstant "github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/log/logger"
	provisionConstant "github.com/funtimecoding/soil/pkg/provision/constant"
	"github.com/funtimecoding/soil/pkg/provision/store"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/system/reaper"
	"github.com/funtimecoding/soil/pkg/system/run"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/option"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/runner"
	"github.com/prometheus/client_golang/prometheus"
	"os"
	"path/filepath"
	"testing"
)

func build(
	t *testing.T,
	c *mock_client.Client,
	hostname string,
	unlockExit int,
) *Tester {
	t.Helper()
	base := t.TempDir()
	remote := filepath.Join(base, "remote.git")
	clone := filepath.Join(base, "clone")
	g := run.New()
	g.Start(
		gitConstant.Command,
		gitConstant.Init,
		"--bare",
		"--initial-branch=main",
		remote,
	)
	g = run.New()
	g.Start("git", "clone", remote, clone)
	g = run.New()
	g.Directory = clone
	g.Start(
		"git",
		"-c",
		"user.name=terraform-tester",
		"-c",
		"user.email=terraform-tester@localhost",
		"commit",
		"--allow-empty",
		gitConstant.MessageArgument,
		"initial",
	)
	g = run.New()
	g.Directory = clone
	g.Start("git", "push", "origin", provisionConstant.RunnerBranch)
	result := &Tester{
		t:        t,
		Cluster:  c,
		Reporter: memory.New(),
		store:    store.New(lite.NewMemory(), "terraform_runs"),
		log:      filepath.Join(base, "calls.log"),
	}
	executable := filepath.Join(base, "recorder")

	if e := os.WriteFile(
		executable,
		[]byte(fmt.Sprintf(constant.FixtureRecorder, result.log, unlockExit)),
		0o755,
	); e != nil {
		t.Fatal(e)
	}

	o := option.New()
	o.Repository = remote
	o.ClonePath = clone
	o.TerraformPath = "."
	o.Executable = executable
	o.Hostname = hostname
	o.Namespace = constant.StateNamespace
	result.Runner = runner.New(
		o,
		result.store,
		logger.New(context.Background()),
		result.Reporter,
		reaper.New(memory.New()),
		prometheus.NewRegistry(),
		c,
	)
	result.Runner.Start()
	t.Cleanup(func() { result.Runner.Close() })
	result.waitForApply()

	return result
}
