package connection

import (
	"database/sql"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/relational/lite/constant"
	_ "github.com/glebarez/go-sqlite"
	"sync/atomic"
)

var memoryCounter atomic.Int64

func NewMemory() *sql.DB {
	// Named shared-cache database: every pooled connection sees the same
	// memory database (a plain :memory: is one database per connection),
	// and the counter keeps separate NewMemory calls isolated
	database, e := sql.Open(
		constant.DriverName,
		fmt.Sprintf(
			"file:memory%d?mode=memory&cache=shared&_pragma=foreign_keys(1)",
			memoryCounter.Add(1),
		),
	)
	errors.PanicOnError(e)

	return database
}
