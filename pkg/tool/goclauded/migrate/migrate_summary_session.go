package migrate

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"gorm.io/gorm"
	"log"
)

func migrateSummarySession(d *gorm.DB) {
	superseded := d.Exec(
		`DELETE FROM summary WHERE identifier NOT IN (
			SELECT MAX(identifier) FROM summary GROUP BY session_identifier
		)`,
	)
	errors.PanicOnError(superseded.Error)

	if superseded.RowsAffected > 0 {
		log.Printf(
			"summary session migration: %d superseded rows dropped",
			superseded.RowsAffected,
		)
	}

	errors.PanicOnError(
		d.Exec(
			"CREATE UNIQUE INDEX IF NOT EXISTS summary_session ON summary (session_identifier)",
		).Error,
	)
}
