package service_tester

import (
	"github.com/funtimecoding/soil/pkg/event/notifier"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/migrate"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/service"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/store"
	"path/filepath"
	"testing"
)

func New(t *testing.T) *Tester {
	t.Helper()
	d := relational.Open(
		logger.New(t.Context()),
		"",
		filepath.Join(t.TempDir(), "maintenance.db"),
	)
	migrate.AutoMigrate(d)
	n := notifier.New()
	result := &Tester{
		Service:  service.New(store.New(d), n),
		Notifier: n,
		database: d,
	}
	t.Cleanup(result.Close)

	return result
}
