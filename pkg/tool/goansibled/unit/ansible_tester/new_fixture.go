package ansible_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	gitConstant "github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/network"
	provisionConstant "github.com/funtimecoding/soil/pkg/provision/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/run"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/constant"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/option"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func NewFixture(
	t *testing.T,
	inventory http.Handler,
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
		"user.name=ansible-tester",
		"-c",
		"user.email=ansible-tester@localhost",
		"commit",
		"--allow-empty",
		gitConstant.MessageArgument,
		"initial",
	)
	c = run.New()
	c.Directory = clone
	c.Start("git", "push", "origin", provisionConstant.RunnerBranch)
	binary := filepath.Join(base, "bin")
	system.EnsurePathExists(binary)
	system.WriteFile(
		filepath.Join(binary, constant.Command),
		[]byte(constant.FixtureScript),
		0755,
	)
	t.Setenv(
		"PATH",
		join.Empty(binary, string(os.PathListSeparator), os.Getenv("PATH")),
	)
	capture := filepath.Join(base, "capture")
	system.EnsurePathExists(capture)
	t.Setenv(constant.FixtureCaptureEnvironment, capture)
	o := option.New()
	o.LitePath = filepath.Join(base, "runs.db")
	o.Repository = remote
	o.CloneDirectory = clone
	o.AnsibleDirectory = "."
	o.Playbook = []string{constant.FixturePlaybook}

	if inventory != nil {
		s := httptest.NewServer(inventory)
		t.Cleanup(s.Close)
		host, port := network.SplitHostPort(s.Listener.Addr().String())
		l := locator.New(host).Port(port).Insecure()
		o.Inventory = func() face.Client {
			return client.New(l, constant.FixtureToken)
		}
	}

	return &Tester{
		t:        t,
		Option:   o,
		Reporter: memory.New(),
		capture:  capture,
		litePath: o.LitePath,
	}
}
