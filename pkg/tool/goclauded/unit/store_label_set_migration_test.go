package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/migrate"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/store_tester"
	"path/filepath"
	"testing"
	"time"
)

func TestSingleValueIndexIsDroppedOnStart(t *testing.T) {
	s := store_tester.NewOnFile(t, filepath.Join(t.TempDir(), "label.sqlite"))
	s.EnsureSession("session-1")
	assert.FatalOnError(
		t,
		s.Mapper().Exec(
			"CREATE UNIQUE INDEX session_key ON label (session_identifier, key)",
		).Error,
	)
	migrate.AutoMigrate(s.Mapper())
	reopened := store.New(s.Mapper(), time.Now)
	_, e := reopened.SetLabel(
		"session-1",
		constant.FixtureLabelKey,
		[]string{constant.FixtureCompatible, "on_call"},
	)
	assert.FatalOnError(t, e)
	values, f := reopened.GetLabel("session-1", constant.FixtureLabelKey)
	assert.FatalOnError(t, f)
	assert.Any(t, []string{"lead", "on_call"}, values)
}
