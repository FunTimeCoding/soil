package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"gorm.io/gorm"
)

func migrateForeignKeys(d *gorm.DB) {
	recreateWithFK(
		d,
		"event",
		`CREATE TABLE event_new (
			identifier INTEGER PRIMARY KEY AUTOINCREMENT,
			session_identifier TEXT NOT NULL
				REFERENCES session(identifier) ON DELETE CASCADE,
			kind TEXT,
			actor TEXT,
			created_at DATETIME
		)`,
		"identifier, session_identifier, kind, actor, created_at",
	)
	recreateWithFK(
		d,
		"completion",
		`CREATE TABLE completion_new (
			identifier INTEGER PRIMARY KEY AUTOINCREMENT,
			session_identifier TEXT NOT NULL
				REFERENCES session(identifier) ON DELETE CASCADE,
			name TEXT,
			kind TEXT,
			topic TEXT,
			summary TEXT,
			created_at DATETIME,
			sequence INTEGER
		)`,
		"identifier, session_identifier, name, kind, topic, summary, created_at, sequence",
	)
	recreateWithFK(
		d,
		constant.SummaryTable,
		`CREATE TABLE summary_new (
			identifier INTEGER PRIMARY KEY AUTOINCREMENT,
			session_identifier TEXT NOT NULL
				REFERENCES session(identifier) ON DELETE CASCADE,
			name TEXT,
			body TEXT,
			created_at DATETIME
		)`,
		"identifier, session_identifier, name, body, created_at",
	)
	recreateWithFK(
		d,
		constant.Label,
		`CREATE TABLE label_new (
			identifier INTEGER PRIMARY KEY AUTOINCREMENT,
			session_identifier TEXT NOT NULL
				REFERENCES session(identifier) ON DELETE CASCADE,
			key TEXT,
			value TEXT
		)`,
		"identifier, session_identifier, key, value",
	)
	recreateWithFK(
		d,
		constant.Pulse,
		`CREATE TABLE pulse_new (
			identifier INTEGER PRIMARY KEY AUTOINCREMENT,
			session_identifier TEXT NOT NULL
				REFERENCES session(identifier) ON DELETE CASCADE,
			from_name TEXT,
			body TEXT,
			consumed NUMERIC,
			created_at DATETIME
		)`,
		"identifier, session_identifier, from_name, body, consumed, created_at",
	)
}
