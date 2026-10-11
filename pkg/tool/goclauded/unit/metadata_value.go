package unit

import (
	"gorm.io/gorm"
	"testing"
)

func metadataValue(
	t *testing.T,
	d *gorm.DB,
	key string,
) string {
	t.Helper()
	var result []string
	d.Raw("SELECT value FROM event_metadata WHERE key = ?", key).Scan(&result)

	if len(result) == 0 {
		return ""
	}

	return result[0]
}
