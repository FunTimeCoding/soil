package worker_tester

import (
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/migrate"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/store"
)

func (o *Tester) Store() *store.Store {
	o.t.Helper()
	m := lite.NewMemory()
	o.t.Cleanup(func() { relational.CloseMapper(m) })
	migrate.AutoMigrate(m)

	return store.New(m)
}
