package brave_tester

import (
	"github.com/funtimecoding/soil/pkg/brave/constant"
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"path/filepath"
	"testing"
)

func Settings(
	t *testing.T,
	preferences string,
	bookmarks string,
) {
	t.Helper()
	profile := filepath.Join(
		system_tester.Home(t),
		constant.BraveSettings,
		constant.DefaultProfile,
	)
	system_tester.Write(t, profile, constant.PreferencesFile, preferences)

	if bookmarks == "" {
		return
	}

	system_tester.Write(t, profile, constant.BookmarksFile, bookmarks)
}
