package integration

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/gohook"
	"github.com/funtimecoding/soil/pkg/tool/gohook/configuration"
	"github.com/funtimecoding/soil/pkg/tool/gohook/constant"
	"path/filepath"
	"testing"
)

func TestInstallAndUninstall(t *testing.T) {
	clone := repository(t)
	write(
		clone,
		constant.ToolFile,
		"pre-push:\n  - run: a\npre-commit:\n  - run: b\n",
	)
	c := configuration.Load(clone)
	foreign := filepath.Join(c.HooksDirectory(), "post-merge")
	write(c.HooksDirectory(), "post-merge", "#!/bin/sh\necho theirs\n")
	gohook.Install(c)
	stub := filepath.Join(c.HooksDirectory(), "pre-push")
	assert.True(t, system.FileExists(stub))
	assert.True(
		t,
		system.FileExists(filepath.Join(c.HooksDirectory(), "pre-commit")),
	)
	assert.StringContains(
		t,
		"exec gohook run pre-push",
		system.ReadFileUnsafe(stub),
	)
	assert.True(t, system.IsExecutable(stub))
	gohook.Uninstall(c)
	assert.False(t, system.FileExists(stub))
	assert.True(t, system.FileExists(foreign))
}
