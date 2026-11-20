package worker_tester

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/migrate"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/store"
)

func (o *Tester) StoreOnFile(path string) *store.Store {
	o.t.Helper()
	m := lite.New(logger.New(o.t.Context()), path)
	o.t.Cleanup(func() { relational.CloseMapper(m) })
	migrate.AutoMigrate(m)

	return store.New(m)
}
