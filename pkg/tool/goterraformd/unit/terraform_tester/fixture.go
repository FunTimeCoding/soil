package terraform_tester

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	gitConstant "github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/kubernetes/client/mock_client"
	"github.com/funtimecoding/soil/pkg/kubernetes/face"
	provisionMock "github.com/funtimecoding/soil/pkg/provision/client/mock_client"
	provisionConstant "github.com/funtimecoding/soil/pkg/provision/constant"
	provisionFace "github.com/funtimecoding/soil/pkg/provision/face"
	"github.com/funtimecoding/soil/pkg/system/run"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/option"
	"os"
	"path/filepath"
	"testing"
)

func fixture(
	t *testing.T,
	c *mock_client.Client,
	hostname string,
	unlockExit int,
	plan string,
	downstream *provisionMock.Client,
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
		log:      filepath.Join(base, "calls.log"),
		litePath: filepath.Join(base, "runs.db"),
	}
	executable := filepath.Join(base, "recorder")
	planFile := filepath.Join(base, "plan.json")

	if e := os.WriteFile(planFile, []byte(plan), 0o644); e != nil {
		t.Fatal(e)
	}

	if e := os.WriteFile(
		executable,
		[]byte(
			fmt.Sprintf(
				constant.FixtureRecorder,
				result.log,
				unlockExit,
				planFile,
			),
		),
		0o755,
	); e != nil {
		t.Fatal(e)
	}

	o := option.New()
	o.LitePath = result.litePath
	o.Cluster = func() face.KubernetesClient { return c }
	o.Repository = remote
	o.ClonePath = clone
	o.TerraformPath = "."
	o.Executable = executable
	o.Hostname = hostname
	o.Namespace = constant.StateNamespace

	if downstream != nil {
		o.DownstreamClients = func() []provisionFace.Client {
			return []provisionFace.Client{downstream}
		}
	}

	result.Option = o

	return result
}
