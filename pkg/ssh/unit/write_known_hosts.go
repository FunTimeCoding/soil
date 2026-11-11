package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"os"
	"path/filepath"
	"testing"
)

func writeKnownHosts(
	t *testing.T,
	content string,
) {
	t.Helper()
	home := t.TempDir()
	t.Setenv(soil.HomeEnvironment, home)
	assert.FatalOnError(t, os.Mkdir(filepath.Join(home, ".ssh"), 0o700))
	assert.FatalOnError(
		t,
		os.WriteFile(
			filepath.Join(home, ".ssh", "known_hosts"),
			[]byte(content),
			0o600,
		),
	)
}
