package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/record"
	"gorm.io/gorm"
)

func migrateLabelChange(d *gorm.DB) {
	var rows []*record.LabelChange
	d.Raw(
		"SELECT event_identifier, value FROM event_metadata WHERE key = ?",
		constant.LegacyChange,
	).Scan(&rows)

	if len(rows) == 0 {
		return
	}

	for _, r := range rows {
		key, past, now := parseLabelChange(r.Value)

		if key == "" {
			continue
		}

		insertEventMetadata(d, r.EventIdentifier, constant.Key, key)
		insertEventMetadata(d, r.EventIdentifier, constant.Past, past)
		insertEventMetadata(d, r.EventIdentifier, constant.Now, now)
	}

	d.Exec("DELETE FROM event_metadata WHERE key = ?", constant.LegacyChange)
}
