package unit

import (
	"context"
	"database/sql"
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	soil "github.com/funtimecoding/soil/pkg/constant"
	libraryErrors "github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/relational/constant"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/relational/lite/connection"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
	"path/filepath"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "POSTGRES_LOCATOR", constant.PostgresLocatorEnvironment)
	assert.String(t, "psql", constant.PostgresCommand)
	assert.String(t, "--username", constant.PostgresUserArgument)
	assert.String(t, "--command", constant.PostgresCommandArgument)
	assert.String(t, "--file", constant.PostgresFileArgument)
	assert.String(t, "--echo-all", constant.PostgresEchoAllFlag)
	assert.String(t, "pg_dump", constant.PostgresDumpCommand)
	assert.String(t, "postgres", constant.PostgresDialectName)
}

func TestNewCreatesParentAndAppliesParameters(t *testing.T) {
	p := filepath.Join(t.TempDir(), "nested", soil.TestDatabase)
	d := connection.New(logger.New(context.Background()), p)
	defer func() { libraryErrors.PanicOnError(d.Close()) }()
	var enabled int
	libraryErrors.PanicOnError(d.QueryRow("PRAGMA foreign_keys").Scan(&enabled))
	assert.Integer(t, 1, enabled)
	var journal string
	libraryErrors.PanicOnError(d.QueryRow("PRAGMA journal_mode").Scan(&journal))
	assert.String(t, "wal", journal)
	assert.True(t, system.FileExists(p))
}

func TestNewMemoryIsolatesCalls(t *testing.T) {
	first := connection.NewMemory()
	defer func() { libraryErrors.PanicOnError(first.Close()) }()
	second := connection.NewMemory()
	defer func() { libraryErrors.PanicOnError(second.Close()) }()
	_, e := first.Exec("CREATE TABLE probe (identifier INTEGER)")
	libraryErrors.PanicOnError(e)
	var count int
	libraryErrors.PanicOnError(
		second.QueryRow(
			"SELECT count(*) FROM sqlite_master WHERE name = 'probe'",
		).Scan(&count),
	)
	assert.Integer(t, 0, count)
	var enabled int
	libraryErrors.PanicOnError(
		first.QueryRow("PRAGMA foreign_keys").Scan(&enabled),
	)
	assert.Integer(t, 1, enabled)
}

func TestNewCreatesParentDirectory(t *testing.T) {
	path := filepath.Join(t.TempDir(), "nested", soil.TestDatabase)
	m := lite.New(logger.New(context.Background()), path)
	defer relational.CloseMapper(m)
	assert.True(t, system.FileExists(path))
}

func TestNewMemoryPinsPool(t *testing.T) {
	m := lite.NewMemory()
	defer relational.CloseMapper(m)
	assert.Integer(t, 1, m.ConnPool.(*sql.DB).Stats().MaxOpenConnections)
}

func TestNewMemoryEnforcesForeignKeys(t *testing.T) {
	m := lite.NewMemory()
	defer relational.CloseMapper(m)
	var enabled int
	libraryErrors.PanicOnError(
		m.Raw("PRAGMA foreign_keys").Scan(&enabled).Error,
	)
	assert.Integer(t, 1, enabled)
}

func TestIsErrorCode(t *testing.T) {
	assert.True(t, relational.IsErrorCode(&pgconn.PgError{Code: "a"}, "a"))
	assert.False(t, relational.IsErrorCode(&pgconn.PgError{Code: "b"}, "a"))
}

func TestNotFound(t *testing.T) {
	assert.False(t, relational.NotFound(nil))
	assert.False(t, relational.NotFound(errors.New("test")))
	assert.True(t, relational.NotFound(gorm.ErrRecordNotFound))
}
