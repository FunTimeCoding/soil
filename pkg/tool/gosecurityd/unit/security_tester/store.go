package security_tester

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/migrate"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store"
	"path/filepath"
)

func (o *Tester) Store() *store.Store {
	o.t.Helper()

	if o.store != nil {
		return o.store
	}

	m := lite.New(
		logger.New(o.t.Context()),
		filepath.Join(o.t.TempDir(), "security.sqlite"),
	)
	o.t.Cleanup(func() { relational.CloseMapper(m) })
	migrate.AutoMigrate(m)
	o.store = store.New(m)

	return o.store
}
