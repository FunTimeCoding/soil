package envrc

import (
	"github.com/funtimecoding/soil/pkg/assert"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/strings/join/key_value"
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/environment"
	"os"
	"path/filepath"
	"slices"
	"testing"
)

func TestLoadRemovesVanishedExport(t *testing.T) {
	directory := t.TempDir()
	path := filepath.Join(directory, ".envrc")
	e := environment.New(
		[]string{
			key_value.Equals("PATH", os.Getenv("PATH")),
			key_value.Equals(
				soil.HomeEnvironment,
				os.Getenv(soil.HomeEnvironment),
			),
			"GHOST=stale",
		},
	)
	system_tester.New(t).Write(
		directory,
		".envrc",
		"export GHOST=stale\nexport FRESH=1\n",
	)
	assert.FatalOnError(t, e.Load(path))

	if !slices.Contains(e.Build(), "GHOST=stale") {
		t.Fatal("expected GHOST before removal")
	}

	system_tester.New(t).Write(directory, ".envrc", "export FRESH=1\n")
	assert.FatalOnError(t, e.Load(path))
	built := e.Build()

	if slices.Contains(built, "GHOST=stale") {
		t.Fatal("expected GHOST removed after reload")
	}

	if !slices.Contains(built, "FRESH=1") {
		t.Fatal("expected FRESH to survive")
	}
}

func TestLoadRestoresReturnedExport(t *testing.T) {
	directory := t.TempDir()
	path := filepath.Join(directory, ".envrc")
	e := environment.New(
		[]string{
			key_value.Equals("PATH", os.Getenv("PATH")),
			key_value.Equals(
				soil.HomeEnvironment,
				os.Getenv(soil.HomeEnvironment),
			),
			"GHOST=stale",
		},
	)
	system_tester.New(t).Write(directory, ".envrc", "export GHOST=stale\n")
	assert.FatalOnError(t, e.Load(path))
	system_tester.New(t).Write(directory, ".envrc", "export OTHER=1\n")
	assert.FatalOnError(t, e.Load(path))

	if slices.Contains(e.Build(), "GHOST=stale") {
		t.Fatal("expected GHOST removed")
	}

	system_tester.New(t).Write(directory, ".envrc", "export GHOST=returned\n")
	assert.FatalOnError(t, e.Load(path))

	if !slices.Contains(e.Build(), "GHOST=returned") {
		t.Fatal("expected GHOST restored with new value")
	}
}

func TestLoadKeepsUnrelatedBase(t *testing.T) {
	directory := t.TempDir()
	path := filepath.Join(directory, ".envrc")
	e := environment.New(
		[]string{
			key_value.Equals("PATH", os.Getenv("PATH")),
			key_value.Equals(
				soil.HomeEnvironment,
				os.Getenv(soil.HomeEnvironment),
			),
			"UNRELATED=keep",
		},
	)
	system_tester.New(t).Write(directory, ".envrc", "export SOMETHING=1\n")
	assert.FatalOnError(t, e.Load(path))
	system_tester.New(t).Write(directory, ".envrc", "export SOMETHING=2\n")
	assert.FatalOnError(t, e.Load(path))

	if !slices.Contains(e.Build(), "UNRELATED=keep") {
		t.Fatal("expected unrelated base variable to survive reloads")
	}
}
