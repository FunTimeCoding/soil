package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"path/filepath"
	"testing"
)

func seedLegacyMessage(t *testing.T) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "goclauded.sqlite")
	d := lite.New(logger.New(context.Background()), path)

	for _, statement := range []string{
		"CREATE TABLE `message` (\"identifier\" integer PRIMARY KEY AUTOINCREMENT,`from_name` text,`to_name` text,`body` text,`read` numeric,`created_at` datetime)",
		"CREATE TABLE `queue` (`identifier` integer PRIMARY KEY AUTOINCREMENT,`callsign` text,`kind` text,`body` text,`consumed` numeric,`created_at` datetime, `session_identifier` text, `consumed_at` datetime, `immediate` numeric)",
		"INSERT INTO message (from_name, to_name, body, read, created_at) VALUES ('Ash', 'Cedar', 'an answer', 1, '2026-01-01 12:00:00+00:00')",
		"INSERT INTO queue (callsign, kind, body, consumed, created_at, session_identifier, immediate) VALUES ('Cedar', 'message', 'Ash: an answer', 0, '2026-01-01 12:00:00+00:00', 'session-1', 0)",
	} {
		assert.FatalOnError(t, d.Exec(statement).Error)
	}

	b, e := d.DB()
	assert.FatalOnError(t, e)
	assert.FatalOnError(t, b.Close())

	return path
}
