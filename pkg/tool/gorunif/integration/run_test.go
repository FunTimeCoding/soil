package integration

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/git/unit/repository_tester"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/gorunif/run_if"
	"github.com/funtimecoding/soil/pkg/tool/gorunif/run_if/option"
	"path/filepath"
	"testing"
)

func TestRunAgainstUpstream(t *testing.T) {
	r := repository_tester.New(t)
	r.Write("pkg/a.go", "package a\n")
	r.Commit("go change")
	o := option.New()
	o.Directory = r.Clone
	o.Pattern = "pkg/"
	o.Execute = "touch ran"
	assert.True(t, run_if.Run(o))
	assert.True(t, system.FileExists(filepath.Join(r.Clone, "ran")))
	o.Pattern = "doc/"
	o.Execute = "touch never"
	assert.False(t, run_if.Run(o))
	assert.False(t, system.FileExists(filepath.Join(r.Clone, "never")))
}

func TestRunExplicitRange(t *testing.T) {
	r := repository_tester.New(t)
	r.Write("doc/x.md", "x\n")
	r.Commit("doc")
	r.Write("pkg/a.go", "package a\n")
	r.Commit("code")
	o := option.New()
	o.Directory = r.Clone
	o.Base = "HEAD~1"
	o.Head = "HEAD"
	o.Pattern = "doc/"
	o.Execute = "touch never"
	assert.False(t, run_if.Run(o))
	o.Pattern = constant.GoExtension
	o.Suffix = true
	o.Execute = "touch ran"
	assert.True(t, run_if.Run(o))
	assert.True(t, system.FileExists(filepath.Join(r.Clone, "ran")))
}

func TestRunWithoutUpstreamRunsAlways(t *testing.T) {
	root := t.TempDir()
	repository_tester.Command(
		root,
		"init",
		"--initial-branch=main",
		constant.CurrentDirectory,
	)
	o := option.New()
	o.Directory = root
	o.Pattern = "nothing/"
	o.Execute = "touch ran"
	assert.True(t, run_if.Run(o))
	assert.True(t, system.FileExists(filepath.Join(root, "ran")))
}
