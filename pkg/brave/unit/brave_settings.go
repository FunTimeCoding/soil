package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/brave/constant"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"os"
	"path/filepath"
	"testing"
)

func braveSettings(
	t *testing.T,
	preferences string,
	bookmarks string,
) {
	t.Helper()
	home := t.TempDir()
	t.Setenv(soil.HomeEnvironment, home)
	p := filepath.Join(home, constant.BraveSettings, constant.DefaultProfile)
	assert.FatalOnError(t, os.MkdirAll(p, 0o755))
	assert.FatalOnError(
		t,
		os.WriteFile(
			filepath.Join(p, constant.PreferencesFile),
			[]byte(preferences),
			0o644,
		),
	)

	if bookmarks == "" {
		return
	}

	assert.FatalOnError(
		t,
		os.WriteFile(
			filepath.Join(p, constant.BookmarksFile),
			[]byte(bookmarks),
			0o644,
		),
	)
}
