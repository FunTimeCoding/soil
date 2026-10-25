package connection

import (
	"database/sql"
	"fmt"
	"github.com/funtimecoding/soil/pkg/relational/lite/zone/connector"
	"sync/atomic"
)

var memoryCounter atomic.Int64

func NewMemory() *sql.DB {
	return sql.OpenDB(
		connector.New(
			Driver(),
			fmt.Sprintf(
				"file:memory%d?mode=memory&cache=shared&_pragma=foreign_keys(1)",
				memoryCounter.Add(1),
			),
		),
	)
}
