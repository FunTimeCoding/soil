package service_tester

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	gitConstant "github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/kubernetes/client/mock_client"
	"github.com/funtimecoding/soil/pkg/kubernetes/face"
	provisionConstant "github.com/funtimecoding/soil/pkg/provision/constant"
	provisionFace "github.com/funtimecoding/soil/pkg/provision/face"
	"github.com/funtimecoding/soil/pkg/system/run"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/option"
	"os"
	"path/filepath"
	"testing"
)

func NewFixture(
	t *testing.T,
	c *mock_client.Client,
	o ...Option,
) *Tester {
	t.Helper()
	base := t.TempDir()
	result := &Tester{
		t:        t,
		Cluster:  c,
		Reporter: memory.New(),
		hostname: constant.FixtureHostname,
		plan:     constant.FixtureEmptyPlan,
		log:      filepath.Join(base, "calls.log"),
		litePath: filepath.Join(base, "runs.db"),
	}

	for _, f := range o {
		f(result)
	}

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
	executable := filepath.Join(base, "recorder")
	planFile := filepath.Join(base, "plan.json")

	if e := os.WriteFile(planFile, []byte(result.plan), 0o644); e != nil {
		t.Fatal(e)
	}

	if e := os.WriteFile(
		executable,
		[]byte(
			fmt.Sprintf(
				constant.FixtureRecorder,
				result.log,
				result.unlockExit,
				planFile,
			),
		),
		0o755,
	); e != nil {
		t.Fatal(e)
	}

	p := option.New()
	p.LitePath = result.litePath
	p.Cluster = func() face.KubernetesClient { return c }
	p.Repository = remote
	p.CloneDirectory = clone
	p.TerraformDirectory = "."
	p.Executable = executable
	p.Hostname = result.hostname
	p.Namespace = constant.StateNamespace

	if len(result.downstream) > 0 {
		var clients []provisionFace.Client

		for _, d := range result.downstream {
			clients = append(clients, d)
		}

		p.DownstreamClients = func() []provisionFace.Client { return clients }
	}

	result.Option = p

	return result
}
