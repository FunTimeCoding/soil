package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/completion"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/event"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/event_metadata"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/label"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/message"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/notification"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/pool_name"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/pulse"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/summary"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/usage_snapshot"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"time"
)

func New(
	path string,
	clock func() time.Time,
) *Store {
	d, e := gorm.Open(sqlite.Open(path), &gorm.Config{})
	errors.PanicOnError(e)
	errors.PanicOnError(d.Exec("PRAGMA journal_mode = WAL").Error)
	errors.PanicOnError(d.Exec("PRAGMA foreign_keys = ON").Error)
	migrateColumns(d)
	errors.PanicOnError(
		d.AutoMigrate(
			session.Stub(),
			message.Stub(),
			event.Stub(),
			event_metadata.Stub(),
			completion.Stub(),
			summary.Stub(),
			pool_name.Stub(),
			usage_snapshot.Stub(),
			label.Stub(),
			pulse.Stub(),
			notification.Stub(),
			queue.Stub(),
		),
	)
	migrateEventMetadata(d)

	if foreignKeysPrecondition(d) {
		migrateEventMetadataConstraint(d)
		migrateForeignKeys(d)
	}

	return &Store{database: d, clock: clock}
}
