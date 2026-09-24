package integration

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/gohook"
	"github.com/funtimecoding/soil/pkg/tool/gohook/configuration"
	"github.com/funtimecoding/soil/pkg/tool/gohook/constant"
	"github.com/funtimecoding/soil/pkg/tool/gohook/option"
	"path/filepath"
	"testing"
)

func TestRunPrePushFiltersJobs(t *testing.T) {
	clone := repository(t)
	write(
		clone,
		constant.ToolFile,
		`pre-push:
  - paths: ['**/*.go']
    run: touch ran-go
  - paths: [strata/manifest/]
    run: touch ran-manifest
  - run: touch ran-always
`,
	)
	write(clone, "pkg/a.go", "package a\n")
	commit(clone, "go change")
	o := option.New()
	o.Hook = "pre-push"
	gohook.Run(configuration.Load(clone), o)
	assert.True(t, system.FileExists(filepath.Join(clone, "ran-go")))
	assert.False(t, system.FileExists(filepath.Join(clone, "ran-manifest")))
	assert.True(t, system.FileExists(filepath.Join(clone, "ran-always")))
}

func TestRunExplicitRange(t *testing.T) {
	clone := repository(t)
	write(
		clone,
		constant.RootFile,
		"pre-push:\n  - paths: [doc/]\n    run: touch ran-doc\n",
	)
	write(clone, "doc/x.md", "x\n")
	commit(clone, "doc")
	o := option.New()
	o.Hook = "pre-push"
	o.Base = "HEAD~1"
	o.Head = "HEAD"
	gohook.Run(configuration.Load(clone), o)
	assert.True(t, system.FileExists(filepath.Join(clone, "ran-doc")))
}

func TestRunHookWithoutJobs(t *testing.T) {
	clone := repository(t)
	write(clone, constant.RootFile, "pre-push:\n  - run: touch never\n")
	o := option.New()
	o.Hook = "post-merge"
	gohook.Run(configuration.Load(clone), o)
	assert.False(t, system.FileExists(filepath.Join(clone, "never")))
}
