package service_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/service/backup"
	backupStore "github.com/funtimecoding/soil/pkg/tool/gopostgresd/store/backup"
)

func (o *Tester) Backup() *backup.Service {
	o.t.Helper()

	if o.backup == nil {
		o.backup = backup.New(
			o.Inventory(),
			o.peers(),
			backupStore.New(o.Directory),
			o.Dump,
			Now,
		)
	}

	return o.backup
}
