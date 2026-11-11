package migrate

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"gorm.io/gorm"
)

func AutoMigrate(d *gorm.DB) {
	migrateCruiseSetting(d)
	errors.PanicOnError(d.AutoMigrate(Models()...))
	migrateCompletedAt(d)

	if foreignKeysPrecondition(d) {
		migrateEventMetadataConstraint(d)
		migrateForeignKeys(d)
		errors.PanicOnError(d.AutoMigrate(Models()...))
	}

	migrateLabelSet(d)
	migrateSummarySession(d)
}
