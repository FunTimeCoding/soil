package connection

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/system"
	"path/filepath"
	"testing"
)

func TestNewCreatesParentAndAppliesParameters(t *testing.T) {
	p := filepath.Join(
		t.TempDir(),
		"nested",
		constant.TestDatabase,
	)
	d := New(logger.New(context.Background()), p)
	defer func() { errors.PanicOnError(d.Close()) }()
	var enabled int
	errors.PanicOnError(
		d.QueryRow("PRAGMA foreign_keys").Scan(&enabled),
	)
	assert.Integer(t, 1, enabled)
	var journal string
	errors.PanicOnError(
		d.QueryRow("PRAGMA journal_mode").Scan(&journal),
	)
	assert.String(t, "wal", journal)
	assert.True(t, system.FileExists(p))
}

func TestNewMemoryIsolatesCalls(t *testing.T) {
	first := NewMemory()
	defer func() { errors.PanicOnError(first.Close()) }()
	second := NewMemory()
	defer func() { errors.PanicOnError(second.Close()) }()
	_, e := first.Exec("CREATE TABLE probe (identifier INTEGER)")
	errors.PanicOnError(e)
	var count int
	errors.PanicOnError(
		second.QueryRow(
			"SELECT count(*) FROM sqlite_master WHERE name = 'probe'",
		).Scan(&count),
	)
	assert.Integer(t, 0, count)
	var enabled int
	errors.PanicOnError(
		first.QueryRow("PRAGMA foreign_keys").Scan(&enabled),
	)
	assert.Integer(t, 1, enabled)
}
