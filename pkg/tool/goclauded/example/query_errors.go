package example

import (
	"context"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"
	"os"
	"path/filepath"
)

func QueryErrors() {
	directory, e := os.MkdirTemp("", "goclauded-query-errors-*")
	errors.PanicOnError(e)
	defer func() { errors.PanicOnError(os.RemoveAll(directory)) }()
	path := filepath.Join(directory, constant.TestDatabase)
	d := lite.New(logger.New(context.Background()), path)
	errors.PanicOnError(d.AutoMigrate(session.Stub()))
	console.Line("=== Find on empty table (not found) ===")
	var i session.Session
	result := d.Where("alias = ?", "missing").Limit(1).Find(&i)
	console.Format("  Error: %v (type: %T)\n", result.Error, result.Error)
	console.Format("  RowsAffected: %d\n", result.RowsAffected)
	console.Line("\n=== Find with invalid column ===")
	var j session.Session
	result = d.Where("nonexistent_column = ?", "x").Limit(1).Find(&j)
	console.Format("  Error: %v (type: %T)\n", result.Error, result.Error)
	console.Line("\n=== Find after database file deleted ===")
	errors.PanicOnError(os.Remove(path))
	var k session.Session
	result = d.Where("alias = ?", "test").Limit(1).Find(&k)
	console.Format("  Error: %v (type: %T)\n", result.Error, result.Error)
	console.Line("\n=== Find after database closed ===")
	inner, e := d.DB()
	errors.PanicOnError(e)
	errors.PanicOnError(inner.Close())
	var l session.Session
	result = d.Where("alias = ?", "test").Limit(1).Find(&l)
	console.Format("  Error: %v (type: %T)\n", result.Error, result.Error)
}
