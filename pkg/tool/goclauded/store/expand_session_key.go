package store

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"gorm.io/gorm"
)

func expandSessionKey(
	d *gorm.DB,
	table string,
) (int64, int64) {
	if !columnExists(d, table, "session_identifier") {
		errors.PanicOnError(
			d.Exec(
				fmt.Sprintf(
					"ALTER TABLE %s ADD COLUMN session_identifier TEXT",
					table,
				),
			).Error,
		)
	}

	attributed := d.Exec(
		fmt.Sprintf(
			`UPDATE %s SET session_identifier = (
				SELECT s.identifier FROM session s
				WHERE s.name = %s.callsign
				AND s.started_at <= %s.created_at
				ORDER BY s.started_at DESC LIMIT 1
			)
			WHERE (session_identifier IS NULL OR session_identifier = '')
			AND callsign IS NOT NULL
			AND callsign != ''
			AND EXISTS (
				SELECT 1 FROM session s
				WHERE s.name = %s.callsign
				AND s.started_at <= %s.created_at
			)`,
			table,
			table,
			table,
			table,
			table,
		),
	)
	errors.PanicOnError(attributed.Error)
	unattributable := d.Exec(
		fmt.Sprintf(
			`DELETE FROM %s
			WHERE session_identifier IS NULL OR session_identifier = ''`,
			table,
		),
	)
	errors.PanicOnError(unattributable.Error)

	return attributed.RowsAffected, unattributable.RowsAffected
}
