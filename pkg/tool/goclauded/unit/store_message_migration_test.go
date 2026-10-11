package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store"
	"testing"
	"time"
)

func TestMigrateDropsMessageReadAndKeepsTheRow(t *testing.T) {
	path := seedLegacyMessage(t)
	d := lite.New(logger.New(context.Background()), path)
	s := store.New(d, func() time.Time { return time.Now().UTC() })
	defer s.Close()
	assert.Integer(t, 0, columnCount(t, d, constant.MessageTable, "read"))
	found, e := s.MessagesByIdentifiers([]uint{1})
	assert.FatalOnError(t, e)
	assert.Count(t, 1, found)
	assert.String(t, "Ash", found[0].FromName)
	assert.String(t, "an answer", found[0].Body)
}

func TestMigrateAddsMessageIdentifierToQueue(t *testing.T) {
	path := seedLegacyMessage(t)
	d := lite.New(logger.New(context.Background()), path)
	s := store.New(d, func() time.Time { return time.Now().UTC() })
	defer s.Close()
	assert.Integer(t, 1, columnCount(t, d, "queue", "message_identifier"))
	drained, e := s.DrainQueue("session-1", "Cedar")
	assert.FatalOnError(t, e)
	assert.Count(t, 1, drained)
	assert.String(t, "Ash: an answer", drained[0].Body)
	assert.True(t, drained[0].MessageIdentifier == nil)
}

func TestMigrateMessageIsIdempotent(t *testing.T) {
	path := seedLegacyMessage(t)
	clock := func() time.Time { return time.Now().UTC() }
	first := store.New(lite.New(logger.New(context.Background()), path), clock)
	first.Close()
	d := lite.New(logger.New(context.Background()), path)
	s := store.New(d, clock)
	defer s.Close()
	assert.Integer(t, 0, columnCount(t, d, constant.MessageTable, "read"))
	found, e := s.MessagesByIdentifiers([]uint{1})
	assert.FatalOnError(t, e)
	assert.Count(t, 1, found)
}
