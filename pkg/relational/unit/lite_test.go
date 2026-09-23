package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/system"
	"path/filepath"
	"testing"
)

func TestNewCreatesParentDirectory(t *testing.T) {
	path := filepath.Join(t.TempDir(), "nested", constant.TestDatabase)
	m := lite.New(logger.New(context.Background()), path)
	defer closeMapper(m)
	assert.True(t, system.FileExists(path))
}

func TestNewMemoryPinsPool(t *testing.T) {
	m := lite.NewMemory()
	defer closeMapper(m)
	inner, e := m.DB()
	errors.PanicOnError(e)
	assert.Integer(t, 1, inner.Stats().MaxOpenConnections)
}

func TestNewMemoryEnforcesForeignKeys(t *testing.T) {
	m := lite.NewMemory()
	defer closeMapper(m)
	var enabled int
	errors.PanicOnError(m.Raw("PRAGMA foreign_keys").Scan(&enabled).Error)
	assert.Integer(t, 1, enabled)
}
