package runner_tester

import (
	"github.com/funtimecoding/soil/pkg/provision/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/run"
	"path/filepath"
)

func (o *Tester) PushCommit(
	name string,
	content string,
) {
	o.t.Helper()
	writer := o.t.TempDir()
	c := run.New()
	c.Start("git", "clone", o.remote, writer)
	system.WriteFile(filepath.Join(writer, name), []byte(content), 0o644)
	c = run.New()
	c.Directory = writer
	c.Start("git", "add", name)
	c = run.New()
	c.Directory = writer
	c.Start(
		"git",
		"-c",
		"user.name=runner-tester",
		"-c",
		"user.email=runner-tester@localhost",
		"commit",
		"-m",
		join.Space("update", name),
	)
	c = run.New()
	c.Directory = writer
	c.Start("git", "push", "origin", constant.RunnerBranch)
}
