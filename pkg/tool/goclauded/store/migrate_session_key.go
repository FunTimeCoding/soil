package store

import (
	"gorm.io/gorm"
	"log"
)

func migrateSessionKey(d *gorm.DB) {
	var pending []string

	for _, t := range sessionKeyTables() {
		if sessionKeyPending(d, t) {
			pending = append(pending, t)
		}
	}

	if len(pending) == 0 {
		return
	}

	if !timestampsNormalized(d) {
		log.Print(
			"session key migration declined: timestamps are not universal",
		)

		return
	}

	var attributed int64
	var dropped int64

	for _, t := range pending {
		a, u := expandSessionKey(d, t)
		attributed += a
		dropped += u
	}

	if attributed == 0 && dropped == 0 {
		return
	}

	log.Printf(
		"session key migration: %d rows keyed on their session, %d dropped as unattributable",
		attributed,
		dropped,
	)
}
