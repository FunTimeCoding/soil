package integration

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/git"
	"github.com/funtimecoding/soil/pkg/tool/gohook"
	"github.com/funtimecoding/soil/pkg/tool/gohook/configuration"
	"github.com/funtimecoding/soil/pkg/tool/gohook/constant"
	"testing"
)

func TestInstallRefusesRedirectedHooksPath(t *testing.T) {
	clone := repository(t)
	write(clone, constant.RootFile, "pre-push:\n  - run: a\n")
	command(clone, "config", "core.hooksPath", "/dev/null")
	assert.String(t, "/dev/null", git.HooksPath(clone))
	defer func() { assert.NotNil(t, recover()) }()
	gohook.Install(configuration.Load(clone))
}
