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
		system_tester.New(t).Home(),
		constant.BraveSettings,
		constant.DefaultProfile,
	)
	system_tester.New(t).Write(profile, constant.PreferencesFile, preferences)

	if bookmarks == "" {
		return
	}

	system_tester.New(t).Write(profile, constant.BookmarksFile, bookmarks)
}
