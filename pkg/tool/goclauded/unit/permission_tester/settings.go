package permission_tester

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/assert"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"os"
	"path/filepath"
)

func (o *Tester) Settings(
	root string,
	file string,
	allow ...string,
) {
	o.t.Helper()
	directory := filepath.Join(root, generativeConstant.SettingsDirectory)
	assert.FatalOnError(o.t, os.MkdirAll(directory, 0o755))
	b, e := json.Marshal(
		map[string]any{"permissions": map[string]any{"allow": allow}},
	)
	assert.FatalOnError(o.t, e)
	assert.FatalOnError(
		o.t,
		os.WriteFile(filepath.Join(directory, file), b, 0o644),
	)
}
