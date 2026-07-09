package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors"
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
	d := lite.New(path)
	errors.PanicOnError(d.AutoMigrate(session.Stub()))
	fmt.Println("=== Find on empty table (not found) ===")
	var i session.Session
	result := d.Where("alias = ?", "missing").Limit(1).Find(&i)
	fmt.Printf("  Error: %v (type: %T)\n", result.Error, result.Error)
	fmt.Printf("  RowsAffected: %d\n", result.RowsAffected)
	fmt.Println("\n=== Find with invalid column ===")
	var j session.Session
	result = d.Where("nonexistent_column = ?", "x").Limit(1).Find(&j)
	fmt.Printf("  Error: %v (type: %T)\n", result.Error, result.Error)
	fmt.Println("\n=== Find after database file deleted ===")
	errors.PanicOnError(os.Remove(path))
	var k session.Session
	result = d.Where("alias = ?", "test").Limit(1).Find(&k)
	fmt.Printf("  Error: %v (type: %T)\n", result.Error, result.Error)
	fmt.Println("\n=== Find after database closed ===")
	inner, e := d.DB()
	errors.PanicOnError(e)
	errors.PanicOnError(inner.Close())
	var l session.Session
	result = d.Where("alias = ?", "test").Limit(1).Find(&l)
	fmt.Printf("  Error: %v (type: %T)\n", result.Error, result.Error)
}
