package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/strings/join/key_value"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/environment"
	"os"
	"path/filepath"
	"testing"
)

func writeEnvrc(
	t *testing.T,
	path string,
	content string,
) {
	t.Helper()
	assert.FatalOnError(t, os.WriteFile(path, []byte(content), 0o644))
}

func buildContains(
	built []string,
	entry string,
) bool {
	for _, candidate := range built {
		if candidate == entry {
			return true
		}
	}

	return false
}

func TestLoadRemovesVanishedExport(t *testing.T) {
	path := filepath.Join(t.TempDir(), ".envrc")
	e := environment.New(
		[]string{
			key_value.Equals("PATH", os.Getenv("PATH")),
			key_value.Equals(
				constant.HomeEnvironment,
				os.Getenv(constant.HomeEnvironment),
			),
			"GHOST=stale",
		},
	)
	writeEnvrc(t, path, "export GHOST=stale\nexport FRESH=1\n")
	assert.FatalOnError(t, e.Load(path))

	if !buildContains(e.Build(), "GHOST=stale") {
		t.Fatal("expected GHOST before removal")
	}

	writeEnvrc(t, path, "export FRESH=1\n")
	assert.FatalOnError(t, e.Load(path))
	built := e.Build()

	if buildContains(built, "GHOST=stale") {
		t.Fatal("expected GHOST removed after reload")
	}

	if !buildContains(built, "FRESH=1") {
		t.Fatal("expected FRESH to survive")
	}
}

func TestLoadRestoresReturnedExport(t *testing.T) {
	path := filepath.Join(t.TempDir(), ".envrc")
	e := environment.New(
		[]string{
			key_value.Equals("PATH", os.Getenv("PATH")),
			key_value.Equals(
				constant.HomeEnvironment,
				os.Getenv(constant.HomeEnvironment),
			),
			"GHOST=stale",
		},
	)
	writeEnvrc(t, path, "export GHOST=stale\n")
	assert.FatalOnError(t, e.Load(path))
	writeEnvrc(t, path, "export OTHER=1\n")
	assert.FatalOnError(t, e.Load(path))

	if buildContains(e.Build(), "GHOST=stale") {
		t.Fatal("expected GHOST removed")
	}

	writeEnvrc(t, path, "export GHOST=returned\n")
	assert.FatalOnError(t, e.Load(path))

	if !buildContains(e.Build(), "GHOST=returned") {
		t.Fatal("expected GHOST restored with new value")
	}
}

func TestLoadKeepsUnrelatedBase(t *testing.T) {
	path := filepath.Join(t.TempDir(), ".envrc")
	e := environment.New(
		[]string{
			key_value.Equals("PATH", os.Getenv("PATH")),
			key_value.Equals(
				constant.HomeEnvironment,
				os.Getenv(constant.HomeEnvironment),
			),
			"UNRELATED=keep",
		},
	)
	writeEnvrc(t, path, "export SOMETHING=1\n")
	assert.FatalOnError(t, e.Load(path))
	writeEnvrc(t, path, "export SOMETHING=2\n")
	assert.FatalOnError(t, e.Load(path))

	if !buildContains(e.Build(), "UNRELATED=keep") {
		t.Fatal("expected unrelated base variable to survive reloads")
	}
}
