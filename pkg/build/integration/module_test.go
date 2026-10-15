//go:build local

package integration

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/build"
	"github.com/funtimecoding/soil/pkg/build/option"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/run"
	"path/filepath"
	"runtime"
	"testing"
)

func TestModuleBuildsEverySelectedArchitecture(t *testing.T) {
	t.Chdir(t.TempDir())
	o := option.New()
	o.Name = "gosed"
	o.Module = "github.com/funtimecoding/soil"
	o.Version = "v0.11.190"
	o.Targets = []string{"linux-amd64", "darwin-arm64"}
	assert.FatalOnError(t, build.Module(o))
	assert.True(
		t,
		system.FileExists(filepath.Join("tmp", "gosed", "linux-amd64", "gosed")),
	)
	assert.True(
		t,
		system.FileExists(
			filepath.Join("tmp", "gosed", "darwin-arm64", "gosed"),
		),
	)
	assert.False(
		t,
		system.FileExists(
			filepath.Join("tmp", "gosed", "darwin-amd64", "gosed"),
		),
	)
	native := filepath.Join(
		"tmp",
		"gosed",
		build.SystemArchitecture(runtime.GOOS, runtime.GOARCH),
		"gosed",
	)

	if !system.FileExists(native) {
		t.Skipf("no build for the host architecture %s", native)
	}

	r := run.New()
	r.Panic = false
	assert.StringContains(t, "Version: v0.11.190", r.Start(native, "--version"))
}

func TestModuleRefusesAnUnknownVersion(t *testing.T) {
	t.Chdir(t.TempDir())
	o := option.New()
	o.Name = "gosed"
	o.Module = "github.com/funtimecoding/soil"
	o.Version = "v0.0.0-not-a-version"
	o.Targets = []string{"linux-amd64"}
	assert.NotNil(t, build.Module(o))
}

func TestModuleRefusesAMissingCommand(t *testing.T) {
	t.Chdir(t.TempDir())
	o := option.New()
	o.Name = "gonotatool"
	o.Module = "github.com/funtimecoding/soil"
	o.Version = "v0.11.190"
	o.Targets = []string{"linux-amd64"}
	assert.NotNil(t, build.Module(o))
}
