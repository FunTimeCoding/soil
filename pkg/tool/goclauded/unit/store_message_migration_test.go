package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/migrate"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/legacy_store"
	"testing"
	"time"
)

func TestMigrateAddsMessageIdentifierToQueue(t *testing.T) {
	path := legacy_store.Message(t)
	d := lite.New(logger.New(context.Background()), path)
	migrate.AutoMigrate(d)
	s := store.New(d, func() time.Time { return time.Now().UTC() })
	defer relational.CloseMapper(d)
	assert.True(t, lite.ColumnExists(d, "queue", "message_identifier"))
	drained, e := s.DrainQueue("session-1", "Cedar")
	assert.FatalOnError(t, e)
	assert.Count(t, 1, drained)
	assert.String(t, "Ash: an answer", drained[0].Body)
	assert.True(t, drained[0].MessageIdentifier == nil)
}

func TestMigrateMessageIsIdempotent(t *testing.T) {
	path := legacy_store.Message(t)
	clock := func() time.Time { return time.Now().UTC() }
	first := lite.New(logger.New(context.Background()), path)
	migrate.AutoMigrate(first)
	relational.CloseMapper(first)
	d := lite.New(logger.New(context.Background()), path)
	migrate.AutoMigrate(d)
	s := store.New(d, clock)
	defer relational.CloseMapper(d)
	found, e := s.IdentifiedMessages([]uint{1})
	assert.FatalOnError(t, e)
	assert.Count(t, 1, found)
	assert.String(t, "Ash", found[0].FromName)
	assert.String(t, "an answer", found[0].Body)
}
