package connection

import (
	"database/sql"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/relational/constant"
	_ "github.com/glebarez/go-sqlite"
	"sync/atomic"
)

var memoryCounter atomic.Int64

func NewMemory() *sql.DB {
	database, e := sql.Open(
		constant.LiteDriverName,
		fmt.Sprintf(
			"file:memory%d?mode=memory&cache=shared&_pragma=foreign_keys(1)",
			memoryCounter.Add(1),
		),
	)
	errors.PanicOnError(e)

	return database
}
