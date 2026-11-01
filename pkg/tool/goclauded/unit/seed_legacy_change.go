package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store"
	"path/filepath"
	"testing"
	"time"
)

func seedLegacyChange(
	t *testing.T,
	value string,
) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "goclauded.sqlite")
	d := lite.New(logger.New(context.Background()), path)
	defer relational.CloseMapper(d)
	s := store.New(d, func() time.Time { return time.Now().UTC() })
	_, e := s.EnsureSession("session-1")
	assert.FatalOnError(t, e)
	assert.FatalOnError(
		t,
		s.LogEvent(
			"session-1",
			constant.Label,
			"Ash",
			map[string]string{constant.LegacyChange: value},
		),
	)

	return path
}
