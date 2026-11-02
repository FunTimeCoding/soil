package worker_tester

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/migrate"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/store"
	"testing"
)

func NewStoreOnFile(
	t *testing.T,
	path string,
) *store.Store {
	t.Helper()
	m := lite.New(logger.New(t.Context()), path)
	t.Cleanup(func() { relational.CloseMapper(m) })
	migrate.AutoMigrate(m)

	return store.New(m)
}
