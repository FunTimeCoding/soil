package migrate

import "gorm.io/gorm"

func migrateLabelSet(d *gorm.DB) {
	d.Exec("DROP INDEX IF EXISTS session_key")
}
