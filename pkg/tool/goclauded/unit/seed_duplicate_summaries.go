package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store"
	"path/filepath"
	"testing"
	"time"
)

func seedDuplicateSummaries(t *testing.T) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "goclauded.sqlite")
	d := lite.New(logger.New(context.Background()), path)
	defer relational.CloseMapper(d)
	s := store.New(d, func() time.Time { return time.Now().UTC() })
	_, e := s.EnsureSession("session-1")
	assert.FatalOnError(t, e)
	assert.FatalOnError(t, d.Exec("DROP INDEX summary_session").Error)

	for _, body := range []string{"first draft", "rewrite"} {
		assert.FatalOnError(
			t,
			d.Exec(
				"INSERT INTO summary (session_identifier, name, body) VALUES (?, ?, ?)",
				"session-1",
				"Esme",
				body,
			).Error,
		)
	}

	return path
}
