package unit

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/assert"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"os"
	"path/filepath"
	"testing"
)

func permissionSettings(
	t *testing.T,
	root string,
	file string,
	allow ...string,
) {
	t.Helper()
	directory := filepath.Join(root, generativeConstant.SettingsDirectory)
	assert.FatalOnError(t, os.MkdirAll(directory, 0o755))
	b, e := json.Marshal(
		map[string]any{"permissions": map[string]any{"allow": allow}},
	)
	assert.FatalOnError(t, e)
	assert.FatalOnError(
		t,
		os.WriteFile(filepath.Join(directory, file), b, 0o644),
	)
}
