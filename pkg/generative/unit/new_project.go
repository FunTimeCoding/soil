package unit

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"os"
	"path/filepath"
	"testing"
)

func newProject(t *testing.T) (string, string) {
	t.Helper()
	root := t.TempDir()
	scratchpad := t.TempDir()

	for _, d := range []string{
		filepath.Join(constant.FixtureScratchDirectory, "work"),
		constant.FixtureDocumentDirectory,
		constant.SettingsDirectory,
	} {
		assert.FatalOnError(t, os.MkdirAll(filepath.Join(root, d), 0o755))
	}

	b, e := json.Marshal(
		map[string]any{
			"permissions": map[string]any{"allow": constant.FixtureAllowlist},
		},
	)
	assert.FatalOnError(t, e)
	assert.FatalOnError(t, os.WriteFile(settingPath(root), b, 0o644))

	return root, scratchpad
}
