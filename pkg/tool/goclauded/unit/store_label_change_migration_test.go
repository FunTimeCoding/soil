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

func TestMigrateLabelChangeFromUnset(t *testing.T) {
	path := seedLegacyChange(t, "colour (unset)→green")
	d := lite.New(logger.New(context.Background()), path)
	s := store.New(d, func() time.Time { return time.Now().UTC() })
	defer s.Close()
	assert.String(t, "colour", metadataValue(t, d, constant.Key))
	assert.String(t, "", metadataValue(t, d, constant.Past))
	assert.String(t, "green", metadataValue(t, d, constant.Now))
	assert.String(t, "", metadataValue(t, d, constant.LegacyChange))
}

func TestMigrateLabelChangeReplace(t *testing.T) {
	path := seedLegacyChange(t, "topic first draft→second draft")
	d := lite.New(logger.New(context.Background()), path)
	s := store.New(d, func() time.Time { return time.Now().UTC() })
	defer s.Close()
	assert.String(t, "topic", metadataValue(t, d, constant.Key))
	assert.String(t, "first draft", metadataValue(t, d, constant.Past))
	assert.String(t, "second draft", metadataValue(t, d, constant.Now))
}

func TestMigrateLabelChangeRemove(t *testing.T) {
	path := seedLegacyChange(t, "colour blue→ (unset)")
	d := lite.New(logger.New(context.Background()), path)
	s := store.New(d, func() time.Time { return time.Now().UTC() })
	defer s.Close()
	assert.String(t, "colour", metadataValue(t, d, constant.Key))
	assert.String(t, "blue", metadataValue(t, d, constant.Past))
	assert.String(t, "", metadataValue(t, d, constant.Now))
}

func TestMigrateLabelChangeIsIdempotent(t *testing.T) {
	path := seedLegacyChange(t, "colour blue→green")
	clock := func() time.Time { return time.Now().UTC() }
	first := store.New(lite.New(logger.New(context.Background()), path), clock)
	first.Close()
	d := lite.New(logger.New(context.Background()), path)
	s := store.New(d, clock)
	defer s.Close()
	assert.String(t, "blue", metadataValue(t, d, constant.Past))
	assert.String(t, "green", metadataValue(t, d, constant.Now))
	assert.String(t, "", metadataValue(t, d, constant.LegacyChange))
}
