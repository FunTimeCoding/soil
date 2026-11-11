package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/store_tester"
	"path/filepath"
	"testing"
	"time"
)

func TestLegacyCruiseColumnsAreDropped(t *testing.T) {
	path := filepath.Join(t.TempDir(), "goclauded.db")
	s := store_tester.NewOnFile(t, path)
	assert.FatalOnError(
		t,
		s.Store.Database().Exec("DROP TABLE cruise_setting").Error,
	)
	assert.FatalOnError(
		t,
		s.Store.Database().Exec(
			`CREATE TABLE cruise_setting (
				identifier INTEGER PRIMARY KEY AUTOINCREMENT,
				session_identifier TEXT UNIQUE,
				mode TEXT NOT NULL,
				pace INTEGER,
				promoted_at DATETIME,
				woken INTEGER NOT NULL DEFAULT 0,
				updated_at DATETIME
			)`,
		).Error,
	)
	reopened := store_tester.NewOnFile(t, path)
	assert.FatalOnError(
		t,
		reopened.Store.SetCruise("session-1", new(10*time.Minute)),
	)
	setting, e := reopened.Store.CruiseSetting("session-1")
	assert.FatalOnError(t, e)
	assert.Duration(t, 10*time.Minute, *setting.Grace)
}
